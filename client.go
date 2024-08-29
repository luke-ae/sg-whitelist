package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/luke-ae/sg-whitelist/gql"
)

type Cache interface {
	SetWithTTL(key, value interface{}, cost int64, ttl time.Duration) bool
	Get(key interface{}) (interface{}, bool)
}

type Client struct {
	cli     *http.Client
	baseURL string
	cache   Cache
}

type GQLClient struct {
	cli       *http.Client
	baseURL   string
	cache     Cache
	client    *Client
	gqlClient gql.ConstellationsGraphqlClient
}

func NewGQLClient(url string, client *Client, cache Cache) *GQLClient {
	cli := &http.Client{
		Timeout: time.Second * 30,
	}
	return &GQLClient{
		cli:       cli,
		baseURL:   url,
		cache:     cache,
		client:    client,
		gqlClient: gql.NewClient(cli, url, &clientv2.Options{}),
	}
}

func NewHTTPClient(url string, cache Cache) *Client {
	return &Client{
		cli: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL: url,
		cache:   cache,
	}
}

func (c *Client) SmartQuery(ctx context.Context, contract string, query any, v any) error {
	bz, err := json.Marshal(query)
	if err != nil {
		return err
	}
	base64Query := base64.StdEncoding.EncodeToString(bz)
	queryUrl := fmt.Sprintf("%s/cosmwasm/wasm/v1/contract/%s/smart/%s", c.baseURL, contract, base64Query)

	slog.Info("base64 encoded", "url", queryUrl)

	req, err := http.NewRequest(http.MethodGet, queryUrl, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	resp, err := c.cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bz, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("invalid request url: %s, response: %s", req.URL, string(bz))
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("error unmarshalling data into value %v: %v", v, err)
	}
	return nil
}

type MinterAddrResponse struct {
	Data struct {
		Minter string `json:"minter"`
	} `json:"data,omitempty"`
}

func (c *Client) FetchMinterAddr(ctx context.Context, contract string) (*MinterAddr, error) {
	key := fmt.Sprintf("minter_config/%v", contract)
	value, found := c.cache.Get(key)

	if found {
		slog.Info("cache hit")
		v, ok := value.(*MinterAddr)
		if ok {
			return v, nil
		}
	}

	query := &Queries{
		Minter: &MinterQuery{},
	}

	resp := &MinterAddrResponse{}
	err := c.SmartQuery(ctx, contract, query, resp)
	if err != nil {
		return nil, err
	}

	c.cache.SetWithTTL(key, resp.Data, 1, time.Minute*5)

	return (*MinterAddr)(&resp.Data), nil
}

type MinterConfigResponse struct {
	Data MinterConfig `json:"data"`
}

type MinterConfig struct {
	Admin   string `json:"admin"`
	NftData struct {
		NftDataType string `json:"nft_data_type"`
		Extension   any    `json:"extension"`
		TokenURI    string `json:"token_uri"`
	} `json:"nft_data"`
	PaymentAddress  string `json:"payment_address"`
	PerAddressLimit int    `json:"per_address_limit"`
	NumTokens       int    `json:"num_tokens"`
	EndTime         any    `json:"end_time"`
	Sg721Address    string `json:"sg721_address"`
	Sg721CodeID     int    `json:"sg721_code_id"`
	StartTime       string `json:"start_time"`
	MintPrice       struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"mint_price"`
	Factory   string `json:"factory"`
	Whitelist string `json:"whitelist"`
}

type WhitelistAddr struct {
	Whitelist string `json:"whitelist"`
}

func (c *Client) FetchMinterConfig(ctx context.Context, contract string) (*WhitelistAddr, error) {
	key := fmt.Sprintf("minter_config/%v", contract)
	value, found := c.cache.Get(key)

	if found {
		slog.Info("cache hit")
		v, ok := value.(*WhitelistAddr)
		if ok {
			return v, nil
		}
	}

	query := &Queries{
		Config: &ConfigQuery{},
	}

	resp := &MinterConfigResponse{}
	err := c.SmartQuery(ctx, contract, query, resp)
	if err != nil {
		return nil, err
	}

	c.cache.SetWithTTL(key, resp.Data, 1, time.Minute*5)

	return &WhitelistAddr{
		Whitelist: resp.Data.Whitelist,
	}, nil
}

type IsMember struct {
	IsMember bool `json:"has_member"`
}

func (c *Client) FetchIsWhitelistMember(ctx context.Context, contract, address string) (*IsMember, error) {
	key := fmt.Sprintf("is_member/%v/%v", contract, address)
	value, found := c.cache.Get(key)

	if found {
		slog.Info("cache hit")
		isMember, ok := value.(*IsMember)
		if ok {
			return isMember, nil
		}
	}

	query := &WhitelistQueries{
		HasMember: &WhitelistMemberQuery{
			Member: address,
		},
	}
	resp := &WhitelistHasMemberQueryResponse{}

	err := c.SmartQuery(ctx, contract, query, resp)
	if err != nil {
		return nil, err
	}

	isMember := &IsMember{IsMember: resp.Data.HasMember}

	c.cache.SetWithTTL(key, isMember, 1, time.Minute*5)

	return isMember, nil
}
