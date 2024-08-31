// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"fmt"
	"io"
	"strconv"
)

type AttributeFilter struct {
	Operator *EventAttributeFilterOperator `json:"operator,omitempty"`
	Filters  []*DataFilter                 `json:"filters"`
}

type AttributeMessageFilter struct {
	Operator *EventAttributeFilterOperator `json:"operator,omitempty"`
	Filters  []*MessageDataFilter          `json:"filters"`
}

type Block struct {
	ID          string `json:"id"`
	BlockHeight int    `json:"blockHeight"`
	CreatedAt   string `json:"createdAt"`
}

type BlockConnection struct {
	Edges    []*BlockEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

type BlockEdge struct {
	Node   *Block `json:"node"`
	Cursor string `json:"cursor"`
}

type ContractDataFilter struct {
	Type     string                      `json:"type"`
	Name     string                      `json:"name"`
	Value    string                      `json:"value"`
	Operator *ContractDataFilterOperator `json:"operator,omitempty"`
}

type ContractFilter struct {
	ContractType string                 `json:"contractType"`
	Events       []*ContractFilterEvent `json:"events"`
}

type ContractFilterEvent struct {
	Name   string  `json:"name"`
	Action *string `json:"action,omitempty"`
}

type DataFilter struct {
	Name     string             `json:"name"`
	Value    string             `json:"value"`
	Operator *EventDataOperator `json:"operator,omitempty"`
}

type DateRange struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
}

type Event struct {
	ID              string       `json:"id"`
	EventName       string       `json:"eventName"`
	Action          *string      `json:"action,omitempty"`
	TxHash          *string      `json:"txHash,omitempty"`
	ContractAddr    *string      `json:"contractAddr,omitempty"`
	ContractCodeID  *string      `json:"contractCodeId,omitempty"`
	CreatedAt       string       `json:"createdAt"`
	ContractType    *string      `json:"contractType,omitempty"`
	ContractInfo    *string      `json:"contractInfo,omitempty"`
	IsValid         bool         `json:"isValid"`
	BlockHeight     int          `json:"blockHeight"`
	BlockTxIdx      int          `json:"blockTxIdx"`
	BlockMessageIdx int          `json:"blockMessageIdx"`
	BlockEventIdx   int          `json:"blockEventIdx"`
	Location        int          `json:"location"`
	Expired         bool         `json:"expired"`
	Data            string       `json:"data"`
	Metadata        string       `json:"metadata"`
	Transaction     *Transaction `json:"transaction,omitempty"`
	Message         *Message     `json:"message,omitempty"`
	Contract        *Contract    `json:"contract,omitempty"`
	Owner           *Address     `json:"owner,omitempty"`
	Sender          *Address     `json:"sender,omitempty"`
	Recipient       *Address     `json:"recipient,omitempty"`
	Bidder          *Address     `json:"bidder,omitempty"`
	Buyer           *Address     `json:"buyer,omitempty"`
	Seller          *Address     `json:"seller,omitempty"`
	Operator        *Address     `json:"operator,omitempty"`
	Price           *Price       `json:"price,omitempty"`
	BidPrice        *Price       `json:"bidPrice,omitempty"`
}

type EventConnection struct {
	Edges    []*EventEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

type EventEdge struct {
	Node   *Event `json:"node"`
	Cursor string `json:"cursor"`
}

type EventsFilter struct {
	Operator *EventAttributeFilterOperator `json:"operator,omitempty"`
	Events   []*ContractFilterEvent        `json:"events"`
}

type Message struct {
	ID             string       `json:"id"`
	TypeURL        string       `json:"typeUrl"`
	CosmwasmMethod *string      `json:"cosmwasmMethod,omitempty"`
	Data           string       `json:"data"`
	BlockHeight    int          `json:"blockHeight"`
	BlockTxIdx     int          `json:"blockTxIdx"`
	Idx            int          `json:"idx"`
	CreatedAt      string       `json:"createdAt"`
	Contract       *Contract    `json:"contract,omitempty"`
	Transaction    *Transaction `json:"transaction,omitempty"`
	Events         []*Event     `json:"events,omitempty"`
}

type MessageConnection struct {
	Edges    []*MessageEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type MessageDataFilter struct {
	Path     string             `json:"path"`
	Value    string             `json:"value"`
	Operator *EventDataOperator `json:"operator,omitempty"`
}

type MessageEdge struct {
	Node   *Message `json:"node"`
	Cursor string   `json:"cursor"`
}

type Mutation struct {
}

type PageInfo struct {
	EndCursor       *string `json:"endCursor,omitempty"`
	HasNextPage     bool    `json:"hasNextPage"`
	StartCursor     *string `json:"startCursor,omitempty"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
}

type Price struct {
	Denom         *string        `json:"denom,omitempty"`
	BaseDenom     *string        `json:"baseDenom,omitempty"`
	PrefixedDenom *string        `json:"prefixedDenom,omitempty"`
	Amount        string         `json:"amount"`
	Usd           *PriceWithRate `json:"usd,omitempty"`
	CurrentUsd    *PriceWithRate `json:"currentUsd"`
}

type PriceFilter struct {
	Min *string `json:"min,omitempty"`
	Max *string `json:"max,omitempty"`
}

type PriceWithRate struct {
	Denom       string  `json:"denom"`
	Amount      string  `json:"amount"`
	Rate        string  `json:"rate"`
	CoingeckoID *string `json:"coingeckoId,omitempty"`
	FromDenom   string  `json:"fromDenom"`
	Tracing     *string `json:"tracing,omitempty"`
}

type Query struct {
}

type RecordFilter struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Verified *bool  `json:"verified,omitempty"`
}

type TokenInput struct {
	CollectionAddr string `json:"collectionAddr"`
	TokenID        string `json:"tokenId"`
}

type TraitFilter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Transaction struct {
	ID           string     `json:"id"`
	Hash         string     `json:"hash"`
	CreatedAt    string     `json:"createdAt"`
	Memo         *string    `json:"memo,omitempty"`
	Code         int        `json:"code"`
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	GasWanted    int        `json:"gasWanted"`
	GasUsed      int        `json:"gasUsed"`
	BlockHeight  int        `json:"blockHeight"`
	Idx          int        `json:"idx"`
	Events       []*Event   `json:"events,omitempty"`
	Messages     []*Message `json:"messages,omitempty"`
}

type TransactionConnection struct {
	Edges    []*TransactionEdge `json:"edges"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type TransactionEdge struct {
	Node   *Transaction `json:"node"`
	Cursor string       `json:"cursor"`
}

type Address struct {
	Addr string `json:"addr"`
	Name *Name  `json:"name,omitempty"`
}

type Collection struct {
	ID               string            `json:"id"`
	CollectionAddr   string            `json:"collectionAddr"`
	CreatedAt        string            `json:"createdAt"`
	MintedAt         *string           `json:"mintedAt,omitempty"`
	Image            string            `json:"image"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	TokensCount      *int              `json:"tokensCount,omitempty"`
	OwnerTokensCount *int              `json:"ownerTokensCount,omitempty"`
	Website          *string           `json:"website,omitempty"`
	CreatedByAddr    string            `json:"createdByAddr"`
	Blocked          bool              `json:"blocked"`
	FloorPrice       *int              `json:"floorPrice,omitempty"`
	CreatedBy        *Address          `json:"createdBy"`
	ImageObject      *Image            `json:"imageObject,omitempty"`
	OwnersCount      *int              `json:"ownersCount,omitempty"`
	Owners           *CollectionOwners `json:"owners"`
}

type CollectionNode struct {
	Offset      int           `json:"offset"`
	Limit       int           `json:"limit"`
	Total       int           `json:"total"`
	Collections []*Collection `json:"collections"`
}

type CollectionOwner struct {
	Count int      `json:"count"`
	Owner *Address `json:"owner"`
}

type CollectionOwners struct {
	TotalCount int                `json:"totalCount"`
	Owners     []*CollectionOwner `json:"owners"`
}

type CollectionTokenCount struct {
	CollectionAddr string `json:"collectionAddr"`
	TokenCount     int    `json:"tokenCount"`
}

type CollectionTrait struct {
	Name   string                  `json:"name"`
	Values []*CollectionTraitValue `json:"values"`
}

type CollectionTraitValue struct {
	Value            string   `json:"value"`
	RarityPercent    *float64 `json:"rarityPercent,omitempty"`
	NumTokens        int      `json:"numTokens"`
	NumTokensForSale int      `json:"numTokensForSale"`
	TraitFloorPrice  *int     `json:"traitFloorPrice,omitempty"`
}

type Contract struct {
	ContractAddr     string  `json:"contractAddr"`
	ContractType     string  `json:"contractType"`
	ContractCodeID   *string `json:"contractCodeId,omitempty"`
	ContractVersion  *string `json:"contractVersion,omitempty"`
	ContractLabel    *string `json:"contractLabel,omitempty"`
	BlockHeight      int     `json:"blockHeight"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
	LastErrorAt      *string `json:"lastErrorAt,omitempty"`
	Blocked          bool    `json:"blocked"`
	CreatedMessageID *string `json:"createdMessageId,omitempty"`
	ContractInfo     string  `json:"contractInfo"`
}

type ContractNode struct {
	Offset    int         `json:"offset"`
	Limit     int         `json:"limit"`
	Total     int         `json:"total"`
	Contracts []*Contract `json:"contracts"`
}

type IbcClient struct {
	ClientID            string `json:"clientId"`
	ClientType          string `json:"clientType"`
	CounterpartyChainID string `json:"counterpartyChainId"`
	Signer              string `json:"signer"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
	BlockHeight         int    `json:"blockHeight"`
}

type Image struct {
	URL           string  `json:"url"`
	ContentType   *string `json:"contentType,omitempty"`
	ContentLength *int    `json:"contentLength,omitempty"`
	Width         *int    `json:"width,omitempty"`
	Height        *int    `json:"height,omitempty"`
}

type Name struct {
	ID                string        `json:"id"`
	ContractAddr      string        `json:"contractAddr"`
	Name              string        `json:"name"`
	OwnerAddr         string        `json:"ownerAddr"`
	ImageURL          *string       `json:"imageUrl,omitempty"`
	AssociatedAddr    *string       `json:"associatedAddr,omitempty"`
	AskPrice          *Price        `json:"askPrice,omitempty"`
	ForSale           bool          `json:"forSale"`
	MintedAt          *string       `json:"mintedAt,omitempty"`
	HighestOffer      *string       `json:"highestOffer,omitempty"`
	Records           []*NameRecord `json:"records,omitempty"`
	HighestOfferEvent *Event        `json:"highestOfferEvent,omitempty"`
}

type NameNode struct {
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Total  int     `json:"total"`
	Names  []*Name `json:"names"`
}

type NameRecord struct {
	ID           string `json:"id"`
	ContractAddr string `json:"contractAddr"`
	Name         string `json:"name"`
	RecordName   string `json:"recordName"`
	RecordValue  string `json:"recordValue"`
	Verified     bool   `json:"verified"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type OwnedCollectionsNode struct {
	Offset                     int                     `json:"offset"`
	Limit                      int                     `json:"limit"`
	Total                      int                     `json:"total"`
	Collections                []*Collection           `json:"collections"`
	OwnedCollectionsTokenCount []*CollectionTokenCount `json:"ownedCollectionsTokenCount"`
}

type Sale struct {
	SaleType       string  `json:"saleType"`
	TxHash         *string `json:"txHash,omitempty"`
	BlockHeight    int     `json:"blockHeight"`
	CollectionAddr string  `json:"collectionAddr"`
	TokenID        string  `json:"tokenId"`
	SalePriceStars int     `json:"salePriceStars"`
	SalePriceUsd   string  `json:"salePriceUsd"`
	Seller         *string `json:"seller,omitempty"`
	Buyer          *string `json:"buyer,omitempty"`
	RarityOrder    *int    `json:"rarityOrder,omitempty"`
	CreatedAt      string  `json:"createdAt"`
}

type SalesNode struct {
	Offset int     `json:"offset"`
	Limit  int     `json:"limit"`
	Total  int     `json:"total"`
	Sales  []*Sale `json:"sales"`
}

type SalesStats struct {
	TimePeriod    *string `json:"timePeriod,omitempty"`
	VolumeStars   string  `json:"volumeStars"`
	VolumeUsd     string  `json:"volumeUsd"`
	SalesCount    int     `json:"salesCount"`
	AvgPriceStars string  `json:"avgPriceStars"`
	AvgPriceUsd   string  `json:"avgPriceUsd"`
}

type SalesStatsNode struct {
	Offset      int               `json:"offset"`
	Limit       int               `json:"limit"`
	Total       int               `json:"total"`
	SalesStats  []*SalesStats     `json:"salesStats"`
	StatsTotals *SalesStatsTotals `json:"statsTotals"`
}

type SalesStatsTotals struct {
	VolumeStars   string `json:"volumeStars"`
	VolumeUsd     string `json:"volumeUsd"`
	SalesCount    int    `json:"salesCount"`
	AvgPriceStars string `json:"avgPriceStars"`
	AvgPriceUsd   string `json:"avgPriceUsd"`
}

type Token struct {
	ID                         string        `json:"id"`
	CollectionAddr             string        `json:"collectionAddr"`
	TokenID                    string        `json:"tokenId"`
	CreatedAt                  string        `json:"createdAt"`
	OwnerAddr                  *string       `json:"ownerAddr,omitempty"`
	Name                       *string       `json:"name,omitempty"`
	Description                *string       `json:"description,omitempty"`
	ForSale                    bool          `json:"forSale"`
	ImageURL                   *string       `json:"imageUrl,omitempty"`
	AnimationURL               *string       `json:"animationUrl,omitempty"`
	RarityOrder                *int          `json:"rarityOrder,omitempty"`
	RarityScore                *string       `json:"rarityScore,omitempty"`
	Price                      *Price        `json:"price,omitempty"`
	PriceExpiresAt             *string       `json:"priceExpiresAt,omitempty"`
	SaleType                   *string       `json:"saleType,omitempty"`
	MintedAt                   *string       `json:"mintedAt,omitempty"`
	LiveAuctionStartTime       *string       `json:"liveAuctionStartTime,omitempty"`
	LiveAuctionEndTime         *string       `json:"liveAuctionEndTime,omitempty"`
	LiveAuctionHighestBid      *Price        `json:"liveAuctionHighestBid,omitempty"`
	ListedAt                   *string       `json:"listedAt,omitempty"`
	SellerAddr                 *string       `json:"sellerAddr,omitempty"`
	IsEscrowed                 bool          `json:"isEscrowed"`
	EscrowContractAddr         *string       `json:"escrowContractAddr,omitempty"`
	EscrowContractType         *string       `json:"escrowContractType,omitempty"`
	Traits                     []*TokenTrait `json:"traits,omitempty"`
	LiveAuctionEvent           *Event        `json:"liveAuctionEvent,omitempty"`
	HighestLiveAuctionBidEvent *Event        `json:"highestLiveAuctionBidEvent,omitempty"`
	Image                      *Image        `json:"image,omitempty"`
	Animation                  *Image        `json:"animation,omitempty"`
	Owner                      *Address      `json:"owner,omitempty"`
	LastSale                   *Event        `json:"lastSale,omitempty"`
	HighestBidEvent            *Event        `json:"highestBidEvent,omitempty"`
	HighestCollectionBidEvent  *Event        `json:"highestCollectionBidEvent,omitempty"`
	HighestBid                 *string       `json:"highestBid,omitempty"`
	HighestCollectionBid       *string       `json:"highestCollectionBid,omitempty"`
}

type TokenNode struct {
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
	Total  int      `json:"total"`
	Tokens []*Token `json:"tokens"`
}

type TokenTrait struct {
	Name          string   `json:"name"`
	Value         string   `json:"value"`
	RarityPercent *float64 `json:"rarityPercent,omitempty"`
	RarityScore   *float64 `json:"rarityScore,omitempty"`
	Rarity        *float64 `json:"rarity,omitempty"`
}

type AuctionEndPreset string

const (
	AuctionEndPresetNext1Hour   AuctionEndPreset = "NEXT_1_HOUR"
	AuctionEndPresetNext6Hours  AuctionEndPreset = "NEXT_6_HOURS"
	AuctionEndPresetNext24Hours AuctionEndPreset = "NEXT_24_HOURS"
	AuctionEndPresetNext48Hours AuctionEndPreset = "NEXT_48_HOURS"
	AuctionEndPresetNext7Days   AuctionEndPreset = "NEXT_7_DAYS"
)

var AllAuctionEndPreset = []AuctionEndPreset{
	AuctionEndPresetNext1Hour,
	AuctionEndPresetNext6Hours,
	AuctionEndPresetNext24Hours,
	AuctionEndPresetNext48Hours,
	AuctionEndPresetNext7Days,
}

func (e AuctionEndPreset) IsValid() bool {
	switch e {
	case AuctionEndPresetNext1Hour, AuctionEndPresetNext6Hours, AuctionEndPresetNext24Hours, AuctionEndPresetNext48Hours, AuctionEndPresetNext7Days:
		return true
	}
	return false
}

func (e AuctionEndPreset) String() string {
	return string(e)
}

func (e *AuctionEndPreset) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuctionEndPreset(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AuctionEndPreset", str)
	}
	return nil
}

func (e AuctionEndPreset) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type BlockSortBy string

const (
	BlockSortByBlockHeightAsc  BlockSortBy = "BLOCK_HEIGHT_ASC"
	BlockSortByBlockHeightDesc BlockSortBy = "BLOCK_HEIGHT_DESC"
)

var AllBlockSortBy = []BlockSortBy{
	BlockSortByBlockHeightAsc,
	BlockSortByBlockHeightDesc,
}

func (e BlockSortBy) IsValid() bool {
	switch e {
	case BlockSortByBlockHeightAsc, BlockSortByBlockHeightDesc:
		return true
	}
	return false
}

func (e BlockSortBy) String() string {
	return string(e)
}

func (e *BlockSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BlockSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BlockSortBy", str)
	}
	return nil
}

func (e BlockSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CollectionSortBy string

const (
	CollectionSortByTokensCountAsc  CollectionSortBy = "TOKENS_COUNT_ASC"
	CollectionSortByTokensCountDesc CollectionSortBy = "TOKENS_COUNT_DESC"
	CollectionSortByMintedAtAsc     CollectionSortBy = "MINTED_AT_ASC"
	CollectionSortByMintedAtDesc    CollectionSortBy = "MINTED_AT_DESC"
	CollectionSortByVolume24hDesc   CollectionSortBy = "VOLUME_24H_DESC"
	CollectionSortByVolume7dDesc    CollectionSortBy = "VOLUME_7D_DESC"
)

var AllCollectionSortBy = []CollectionSortBy{
	CollectionSortByTokensCountAsc,
	CollectionSortByTokensCountDesc,
	CollectionSortByMintedAtAsc,
	CollectionSortByMintedAtDesc,
	CollectionSortByVolume24hDesc,
	CollectionSortByVolume7dDesc,
}

func (e CollectionSortBy) IsValid() bool {
	switch e {
	case CollectionSortByTokensCountAsc, CollectionSortByTokensCountDesc, CollectionSortByMintedAtAsc, CollectionSortByMintedAtDesc, CollectionSortByVolume24hDesc, CollectionSortByVolume7dDesc:
		return true
	}
	return false
}

func (e CollectionSortBy) String() string {
	return string(e)
}

func (e *CollectionSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CollectionSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CollectionSortBy", str)
	}
	return nil
}

func (e CollectionSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContractDataFilterOperator string

const (
	ContractDataFilterOperatorEqual          ContractDataFilterOperator = "EQUAL"
	ContractDataFilterOperatorGreater        ContractDataFilterOperator = "GREATER"
	ContractDataFilterOperatorGreaterOrEqual ContractDataFilterOperator = "GREATER_OR_EQUAL"
	ContractDataFilterOperatorLower          ContractDataFilterOperator = "LOWER"
	ContractDataFilterOperatorLowerOrEqual   ContractDataFilterOperator = "LOWER_OR_EQUAL"
)

var AllContractDataFilterOperator = []ContractDataFilterOperator{
	ContractDataFilterOperatorEqual,
	ContractDataFilterOperatorGreater,
	ContractDataFilterOperatorGreaterOrEqual,
	ContractDataFilterOperatorLower,
	ContractDataFilterOperatorLowerOrEqual,
}

func (e ContractDataFilterOperator) IsValid() bool {
	switch e {
	case ContractDataFilterOperatorEqual, ContractDataFilterOperatorGreater, ContractDataFilterOperatorGreaterOrEqual, ContractDataFilterOperatorLower, ContractDataFilterOperatorLowerOrEqual:
		return true
	}
	return false
}

func (e ContractDataFilterOperator) String() string {
	return string(e)
}

func (e *ContractDataFilterOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContractDataFilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContractDataFilterOperator", str)
	}
	return nil
}

func (e ContractDataFilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ContractSortBy string

const (
	ContractSortByCreatedAtDesc ContractSortBy = "CREATED_AT_DESC"
	ContractSortByCreatedAtAsc  ContractSortBy = "CREATED_AT_ASC"
)

var AllContractSortBy = []ContractSortBy{
	ContractSortByCreatedAtDesc,
	ContractSortByCreatedAtAsc,
}

func (e ContractSortBy) IsValid() bool {
	switch e {
	case ContractSortByCreatedAtDesc, ContractSortByCreatedAtAsc:
		return true
	}
	return false
}

func (e ContractSortBy) String() string {
	return string(e)
}

func (e *ContractSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ContractSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ContractSortBy", str)
	}
	return nil
}

func (e ContractSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DateGranularity string

const (
	DateGranularityHour        DateGranularity = "HOUR"
	DateGranularityFourHours   DateGranularity = "FOUR_HOURS"
	DateGranularityTwelveHours DateGranularity = "TWELVE_HOURS"
	DateGranularityDay         DateGranularity = "DAY"
	DateGranularityWeek        DateGranularity = "WEEK"
	DateGranularityMonth       DateGranularity = "MONTH"
	DateGranularityQuarter     DateGranularity = "QUARTER"
	DateGranularityYear        DateGranularity = "YEAR"
)

var AllDateGranularity = []DateGranularity{
	DateGranularityHour,
	DateGranularityFourHours,
	DateGranularityTwelveHours,
	DateGranularityDay,
	DateGranularityWeek,
	DateGranularityMonth,
	DateGranularityQuarter,
	DateGranularityYear,
}

func (e DateGranularity) IsValid() bool {
	switch e {
	case DateGranularityHour, DateGranularityFourHours, DateGranularityTwelveHours, DateGranularityDay, DateGranularityWeek, DateGranularityMonth, DateGranularityQuarter, DateGranularityYear:
		return true
	}
	return false
}

func (e DateGranularity) String() string {
	return string(e)
}

func (e *DateGranularity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DateGranularity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DateGranularity", str)
	}
	return nil
}

func (e DateGranularity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DatePreset string

const (
	DatePresetLast24Hours DatePreset = "LAST_24_HOURS"
	DatePresetLast7Days   DatePreset = "LAST_7_DAYS"
	DatePresetLast30Days  DatePreset = "LAST_30_DAYS"
	DatePresetLast90Days  DatePreset = "LAST_90_DAYS"
	DatePresetLastYear    DatePreset = "LAST_YEAR"
	DatePresetAllTime     DatePreset = "ALL_TIME"
)

var AllDatePreset = []DatePreset{
	DatePresetLast24Hours,
	DatePresetLast7Days,
	DatePresetLast30Days,
	DatePresetLast90Days,
	DatePresetLastYear,
	DatePresetAllTime,
}

func (e DatePreset) IsValid() bool {
	switch e {
	case DatePresetLast24Hours, DatePresetLast7Days, DatePresetLast30Days, DatePresetLast90Days, DatePresetLastYear, DatePresetAllTime:
		return true
	}
	return false
}

func (e DatePreset) String() string {
	return string(e)
}

func (e *DatePreset) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DatePreset(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DatePreset", str)
	}
	return nil
}

func (e DatePreset) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventAttributeFilterOperator string

const (
	EventAttributeFilterOperatorAnd EventAttributeFilterOperator = "AND"
	EventAttributeFilterOperatorOr  EventAttributeFilterOperator = "OR"
)

var AllEventAttributeFilterOperator = []EventAttributeFilterOperator{
	EventAttributeFilterOperatorAnd,
	EventAttributeFilterOperatorOr,
}

func (e EventAttributeFilterOperator) IsValid() bool {
	switch e {
	case EventAttributeFilterOperatorAnd, EventAttributeFilterOperatorOr:
		return true
	}
	return false
}

func (e EventAttributeFilterOperator) String() string {
	return string(e)
}

func (e *EventAttributeFilterOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventAttributeFilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventAttributeFilterOperator", str)
	}
	return nil
}

func (e EventAttributeFilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventDataOperator string

const (
	EventDataOperatorEqual          EventDataOperator = "EQUAL"
	EventDataOperatorGreater        EventDataOperator = "GREATER"
	EventDataOperatorGreaterOrEqual EventDataOperator = "GREATER_OR_EQUAL"
	EventDataOperatorLower          EventDataOperator = "LOWER"
	EventDataOperatorLowerOrEqual   EventDataOperator = "LOWER_OR_EQUAL"
)

var AllEventDataOperator = []EventDataOperator{
	EventDataOperatorEqual,
	EventDataOperatorGreater,
	EventDataOperatorGreaterOrEqual,
	EventDataOperatorLower,
	EventDataOperatorLowerOrEqual,
}

func (e EventDataOperator) IsValid() bool {
	switch e {
	case EventDataOperatorEqual, EventDataOperatorGreater, EventDataOperatorGreaterOrEqual, EventDataOperatorLower, EventDataOperatorLowerOrEqual:
		return true
	}
	return false
}

func (e EventDataOperator) String() string {
	return string(e)
}

func (e *EventDataOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventDataOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventDataOperator", str)
	}
	return nil
}

func (e EventDataOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EventSortBy string

const (
	EventSortByBlockHeightAsc  EventSortBy = "BLOCK_HEIGHT_ASC"
	EventSortByBlockHeightDesc EventSortBy = "BLOCK_HEIGHT_DESC"
)

var AllEventSortBy = []EventSortBy{
	EventSortByBlockHeightAsc,
	EventSortByBlockHeightDesc,
}

func (e EventSortBy) IsValid() bool {
	switch e {
	case EventSortByBlockHeightAsc, EventSortByBlockHeightDesc:
		return true
	}
	return false
}

func (e EventSortBy) String() string {
	return string(e)
}

func (e *EventSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventSortBy", str)
	}
	return nil
}

func (e EventSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Filter string

const (
	FilterSales                      Filter = "SALES"
	FilterNameSales                  Filter = "NAME_SALES"
	FilterMints                      Filter = "MINTS"
	FilterAsks                       Filter = "ASKS"
	FilterNameAsks                   Filter = "NAME_ASKS"
	FilterBids                       Filter = "BIDS"
	FilterNameBids                   Filter = "NAME_BIDS"
	FilterAirdrops                   Filter = "AIRDROPS"
	FilterCollectionBids             Filter = "COLLECTION_BIDS"
	FilterReceivedOffersOnOwnedNft   Filter = "RECEIVED_OFFERS_ON_OWNED_NFT"
	FilterReceivedOffersOnOwnedNames Filter = "RECEIVED_OFFERS_ON_OWNED_NAMES"
	FilterNameEvents                 Filter = "NAME_EVENTS"
	FilterFilteredNameEvents         Filter = "FILTERED_NAME_EVENTS"
	FilterSentNameOffers             Filter = "SENT_NAME_OFFERS"
	FilterTokenMetadatas             Filter = "TOKEN_METADATAS"
)

var AllFilter = []Filter{
	FilterSales,
	FilterNameSales,
	FilterMints,
	FilterAsks,
	FilterNameAsks,
	FilterBids,
	FilterNameBids,
	FilterAirdrops,
	FilterCollectionBids,
	FilterReceivedOffersOnOwnedNft,
	FilterReceivedOffersOnOwnedNames,
	FilterNameEvents,
	FilterFilteredNameEvents,
	FilterSentNameOffers,
	FilterTokenMetadatas,
}

func (e Filter) IsValid() bool {
	switch e {
	case FilterSales, FilterNameSales, FilterMints, FilterAsks, FilterNameAsks, FilterBids, FilterNameBids, FilterAirdrops, FilterCollectionBids, FilterReceivedOffersOnOwnedNft, FilterReceivedOffersOnOwnedNames, FilterNameEvents, FilterFilteredNameEvents, FilterSentNameOffers, FilterTokenMetadatas:
		return true
	}
	return false
}

func (e Filter) String() string {
	return string(e)
}

func (e *Filter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Filter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Filter", str)
	}
	return nil
}

func (e Filter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MessageSortBy string

const (
	MessageSortByBlockHeightAsc  MessageSortBy = "BLOCK_HEIGHT_ASC"
	MessageSortByBlockHeightDesc MessageSortBy = "BLOCK_HEIGHT_DESC"
)

var AllMessageSortBy = []MessageSortBy{
	MessageSortByBlockHeightAsc,
	MessageSortByBlockHeightDesc,
}

func (e MessageSortBy) IsValid() bool {
	switch e {
	case MessageSortByBlockHeightAsc, MessageSortByBlockHeightDesc:
		return true
	}
	return false
}

func (e MessageSortBy) String() string {
	return string(e)
}

func (e *MessageSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MessageSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MessageSortBy", str)
	}
	return nil
}

func (e MessageSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NameSortBy string

const (
	NameSortByMintedAtDesc NameSortBy = "MINTED_AT_DESC"
	NameSortByMintedAtAsc  NameSortBy = "MINTED_AT_ASC"
	NameSortByNameDesc     NameSortBy = "NAME_DESC"
	NameSortByNameAsc      NameSortBy = "NAME_ASC"
	NameSortByOffersDesc   NameSortBy = "OFFERS_DESC"
	NameSortByOffersAsc    NameSortBy = "OFFERS_ASC"
)

var AllNameSortBy = []NameSortBy{
	NameSortByMintedAtDesc,
	NameSortByMintedAtAsc,
	NameSortByNameDesc,
	NameSortByNameAsc,
	NameSortByOffersDesc,
	NameSortByOffersAsc,
}

func (e NameSortBy) IsValid() bool {
	switch e {
	case NameSortByMintedAtDesc, NameSortByMintedAtAsc, NameSortByNameDesc, NameSortByNameAsc, NameSortByOffersDesc, NameSortByOffersAsc:
		return true
	}
	return false
}

func (e NameSortBy) String() string {
	return string(e)
}

func (e *NameSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NameSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NameSortBy", str)
	}
	return nil
}

func (e NameSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SaleType string

const (
	SaleTypeFixedPrice  SaleType = "FIXED_PRICE"
	SaleTypeAuction     SaleType = "AUCTION"
	SaleTypeLiveAuction SaleType = "LIVE_AUCTION"
	SaleTypeAny         SaleType = "ANY"
	SaleTypeUnlisted    SaleType = "UNLISTED"
	SaleTypeListed      SaleType = "LISTED"
	SaleTypeExpired     SaleType = "EXPIRED"
)

var AllSaleType = []SaleType{
	SaleTypeFixedPrice,
	SaleTypeAuction,
	SaleTypeLiveAuction,
	SaleTypeAny,
	SaleTypeUnlisted,
	SaleTypeListed,
	SaleTypeExpired,
}

func (e SaleType) IsValid() bool {
	switch e {
	case SaleTypeFixedPrice, SaleTypeAuction, SaleTypeLiveAuction, SaleTypeAny, SaleTypeUnlisted, SaleTypeListed, SaleTypeExpired:
		return true
	}
	return false
}

func (e SaleType) String() string {
	return string(e)
}

func (e *SaleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SaleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SaleType", str)
	}
	return nil
}

func (e SaleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SalesSaleType string

const (
	SalesSaleTypeFixedPrice      SalesSaleType = "FIXED_PRICE"
	SalesSaleTypeAuction         SalesSaleType = "AUCTION"
	SalesSaleTypeOffer           SalesSaleType = "OFFER"
	SalesSaleTypeCollectionOffer SalesSaleType = "COLLECTION_OFFER"
)

var AllSalesSaleType = []SalesSaleType{
	SalesSaleTypeFixedPrice,
	SalesSaleTypeAuction,
	SalesSaleTypeOffer,
	SalesSaleTypeCollectionOffer,
}

func (e SalesSaleType) IsValid() bool {
	switch e {
	case SalesSaleTypeFixedPrice, SalesSaleTypeAuction, SalesSaleTypeOffer, SalesSaleTypeCollectionOffer:
		return true
	}
	return false
}

func (e SalesSaleType) String() string {
	return string(e)
}

func (e *SalesSaleType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SalesSaleType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SalesSaleType", str)
	}
	return nil
}

func (e SalesSaleType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SalesSortBy string

const (
	SalesSortBySaleTimeDesc   SalesSortBy = "SALE_TIME_DESC"
	SalesSortBySaleTimeAsc    SalesSortBy = "SALE_TIME_ASC"
	SalesSortByStarsPriceDesc SalesSortBy = "STARS_PRICE_DESC"
	SalesSortByStarsPriceAsc  SalesSortBy = "STARS_PRICE_ASC"
	SalesSortByUsdPriceDesc   SalesSortBy = "USD_PRICE_DESC"
	SalesSortByUsdPriceAsc    SalesSortBy = "USD_PRICE_ASC"
	SalesSortByRarityDesc     SalesSortBy = "RARITY_DESC"
	SalesSortByRarityAsc      SalesSortBy = "RARITY_ASC"
	SalesSortByTokenIDAsc     SalesSortBy = "TOKEN_ID_ASC"
	SalesSortByTokenIDDesc    SalesSortBy = "TOKEN_ID_DESC"
)

var AllSalesSortBy = []SalesSortBy{
	SalesSortBySaleTimeDesc,
	SalesSortBySaleTimeAsc,
	SalesSortByStarsPriceDesc,
	SalesSortByStarsPriceAsc,
	SalesSortByUsdPriceDesc,
	SalesSortByUsdPriceAsc,
	SalesSortByRarityDesc,
	SalesSortByRarityAsc,
	SalesSortByTokenIDAsc,
	SalesSortByTokenIDDesc,
}

func (e SalesSortBy) IsValid() bool {
	switch e {
	case SalesSortBySaleTimeDesc, SalesSortBySaleTimeAsc, SalesSortByStarsPriceDesc, SalesSortByStarsPriceAsc, SalesSortByUsdPriceDesc, SalesSortByUsdPriceAsc, SalesSortByRarityDesc, SalesSortByRarityAsc, SalesSortByTokenIDAsc, SalesSortByTokenIDDesc:
		return true
	}
	return false
}

func (e SalesSortBy) String() string {
	return string(e)
}

func (e *SalesSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SalesSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SalesSortBy", str)
	}
	return nil
}

func (e SalesSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TokenSortBy string

const (
	TokenSortByRarityDesc                TokenSortBy = "RARITY_DESC"
	TokenSortByRarityAsc                 TokenSortBy = "RARITY_ASC"
	TokenSortByPriceDesc                 TokenSortBy = "PRICE_DESC"
	TokenSortByPriceAsc                  TokenSortBy = "PRICE_ASC"
	TokenSortByListingType               TokenSortBy = "LISTING_TYPE"
	TokenSortByTokenIDAsc                TokenSortBy = "TOKEN_ID_ASC"
	TokenSortByTokenIDDesc               TokenSortBy = "TOKEN_ID_DESC"
	TokenSortByNameAsc                   TokenSortBy = "NAME_ASC"
	TokenSortByNameDesc                  TokenSortBy = "NAME_DESC"
	TokenSortByCollectionAddrTokenIDAsc  TokenSortBy = "COLLECTION_ADDR_TOKEN_ID_ASC"
	TokenSortByMintedAsc                 TokenSortBy = "MINTED_ASC"
	TokenSortByMintedDesc                TokenSortBy = "MINTED_DESC"
	TokenSortByLiveAuctionEndTimeDesc    TokenSortBy = "LIVE_AUCTION_END_TIME_DESC"
	TokenSortByLiveAuctionEndTimeAsc     TokenSortBy = "LIVE_AUCTION_END_TIME_ASC"
	TokenSortByLiveAuctionStartTimeDesc  TokenSortBy = "LIVE_AUCTION_START_TIME_DESC"
	TokenSortByLiveAuctionStartTimeAsc   TokenSortBy = "LIVE_AUCTION_START_TIME_ASC"
	TokenSortByLiveAuctionHighestBidDesc TokenSortBy = "LIVE_AUCTION_HIGHEST_BID_DESC"
	TokenSortByLiveAuctionHighestBidAsc  TokenSortBy = "LIVE_AUCTION_HIGHEST_BID_ASC"
	TokenSortByListedAsc                 TokenSortBy = "LISTED_ASC"
	TokenSortByListedDesc                TokenSortBy = "LISTED_DESC"
	TokenSortByAcquiredAsc               TokenSortBy = "ACQUIRED_ASC"
	TokenSortByAcquiredDesc              TokenSortBy = "ACQUIRED_DESC"
)

var AllTokenSortBy = []TokenSortBy{
	TokenSortByRarityDesc,
	TokenSortByRarityAsc,
	TokenSortByPriceDesc,
	TokenSortByPriceAsc,
	TokenSortByListingType,
	TokenSortByTokenIDAsc,
	TokenSortByTokenIDDesc,
	TokenSortByNameAsc,
	TokenSortByNameDesc,
	TokenSortByCollectionAddrTokenIDAsc,
	TokenSortByMintedAsc,
	TokenSortByMintedDesc,
	TokenSortByLiveAuctionEndTimeDesc,
	TokenSortByLiveAuctionEndTimeAsc,
	TokenSortByLiveAuctionStartTimeDesc,
	TokenSortByLiveAuctionStartTimeAsc,
	TokenSortByLiveAuctionHighestBidDesc,
	TokenSortByLiveAuctionHighestBidAsc,
	TokenSortByListedAsc,
	TokenSortByListedDesc,
	TokenSortByAcquiredAsc,
	TokenSortByAcquiredDesc,
}

func (e TokenSortBy) IsValid() bool {
	switch e {
	case TokenSortByRarityDesc, TokenSortByRarityAsc, TokenSortByPriceDesc, TokenSortByPriceAsc, TokenSortByListingType, TokenSortByTokenIDAsc, TokenSortByTokenIDDesc, TokenSortByNameAsc, TokenSortByNameDesc, TokenSortByCollectionAddrTokenIDAsc, TokenSortByMintedAsc, TokenSortByMintedDesc, TokenSortByLiveAuctionEndTimeDesc, TokenSortByLiveAuctionEndTimeAsc, TokenSortByLiveAuctionStartTimeDesc, TokenSortByLiveAuctionStartTimeAsc, TokenSortByLiveAuctionHighestBidDesc, TokenSortByLiveAuctionHighestBidAsc, TokenSortByListedAsc, TokenSortByListedDesc, TokenSortByAcquiredAsc, TokenSortByAcquiredDesc:
		return true
	}
	return false
}

func (e TokenSortBy) String() string {
	return string(e)
}

func (e *TokenSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TokenSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TokenSortBy", str)
	}
	return nil
}

func (e TokenSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionSortBy string

const (
	TransactionSortByBlockHeightAsc  TransactionSortBy = "BLOCK_HEIGHT_ASC"
	TransactionSortByBlockHeightDesc TransactionSortBy = "BLOCK_HEIGHT_DESC"
)

var AllTransactionSortBy = []TransactionSortBy{
	TransactionSortByBlockHeightAsc,
	TransactionSortByBlockHeightDesc,
}

func (e TransactionSortBy) IsValid() bool {
	switch e {
	case TransactionSortByBlockHeightAsc, TransactionSortByBlockHeightDesc:
		return true
	}
	return false
}

func (e TransactionSortBy) String() string {
	return string(e)
}

func (e *TransactionSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionSortBy", str)
	}
	return nil
}

func (e TransactionSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
