package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/luke-ae/sg-whitelist/gql"
)

func (c *GQLClient) WhitelistedCollections(w http.ResponseWriter, r *http.Request) {
	walletAddr := chi.URLParam(r, "wallet")

	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if limit == 0 {
		limit = 25
	}

	resp, err := c.gqlClient.LatestCollections(r.Context(), offset, limit)
	if err != nil {
		json.NewEncoder(w).Encode(Error{Error: err.Error()})
	}

	var (
		collections = resp.Collections.Collections
		memberChan  = make(chan *MemberResponse, len(collections))
		wg          sync.WaitGroup
	)

	for _, coll := range collections {
		wg.Add(1)
		go func(coll *gql.LatestCollections_Collections_Collections) {
			defer wg.Done()
			resp, err := c.client.FetchMinterAddr(r.Context(), coll.CollectionAddr)
			if err != nil {
				fmt.Println(err)
				return
			}

			minterResp, err := c.client.FetchMinterConfig(r.Context(), resp.MinterAddr)
			if err != nil {
				fmt.Println(err)
				return
			}

			if minterResp.WhitelistAddr == "" {
				memberChan <- &MemberResponse{
					Minter:     resp.MinterAddr,
					Collection: coll.CollectionAddr,
					Name:       coll.Name,
					Image:      coll.Image,
					MintedAt:   coll.MintedAt,
				}
				return
			}

			wlResp, err := c.client.FetchIsWhitelistMember(r.Context(), minterResp.WhitelistAddr, walletAddr)
			if err != nil {
				memberChan <- &MemberResponse{
					Minter:     resp.MinterAddr,
					Collection: coll.CollectionAddr,
					Name:       coll.Name,
					Image:      coll.Image,
					MintedAt:   coll.MintedAt,
				}
				return
			}
			memberChan <- &MemberResponse{
				Minter:     resp.MinterAddr,
				Collection: coll.CollectionAddr,
				Name:       coll.Name,
				Image:      coll.Image,
				HasMember:  wlResp.IsMember,
				MintedAt:   coll.MintedAt,
			}
		}(coll)
	}

	go func() {
		wg.Wait()
		close(memberChan)
	}()

	var members []*MemberResponse
	for m := range memberChan {
		members = append(members, m)
	}

	start := time.Now()
	sort.Sort(ByDate(members))
	duration := time.Until(start)
	fmt.Printf("Sorting took %v\n", duration)
	WriteJSON(w, http.StatusOK, members)
}

type ByDate []*MemberResponse

func (a ByDate) Len() int { return len(a) }
func (a ByDate) Less(i, j int) bool {
	layout := "2006-01-02T15:04:05.000000Z07:00"
	date1, _ := time.Parse(layout, *a[i].MintedAt)
	date2, _ := time.Parse(layout, *a[j].MintedAt)
	return date1.After(date2)
}
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
