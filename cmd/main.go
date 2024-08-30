package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/dgraph-io/ristretto"
	"github.com/go-chi/chi/v5"
	gqlTypes "github.com/luke-ae/sg-whitelist/gql"
	"golang.org/x/exp/slog"
)

func main() {
	r := chi.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 33, // maximum cost of cache (8GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		panic("error creating cache layer")
	}

	c := NewHTTPClient("http://rest.stargaze-apis.com", cache)
	gql := NewGQLClient("https://constellations-api.mainnet.stargaze-apis.com/graphql", c, cache)

	r.Get("/on_whitelist/{wallet}", func(w http.ResponseWriter, r *http.Request) {
		walletAddr := chi.URLParam(r, "wallet")

		offsetInt, _ := strconv.Atoi(r.URL.Query().Get("offset"))
		limitInt, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		if limitInt == 0 {
			limitInt = 10
		}

		offset := int64(offsetInt)
		limit := int64(limitInt)

		resp, err := gql.gqlClient.Collections(context.Background(), &offset, &limit)

		if err != nil {
			json.NewEncoder(w).Encode(Error{Error: err.Error()})
		}
		collections := resp.Collections.Collections
		memberChan := make(chan *MemberResponse, len(collections))
		var wg sync.WaitGroup

		for _, coll := range collections {
			wg.Add(1)
			go func(coll *gqlTypes.Collections_Collections_Collections) {
				defer wg.Done()
				resp, err := c.FetchMinterAddr(r.Context(), coll.CollectionAddr)
				if err != nil {
					fmt.Println(err)
					return
				}

				minterResp, err := c.FetchMinterConfig(r.Context(), resp.Minter)
				if err != nil {
					fmt.Println(err)
					return
				}

				wlResp, err := c.FetchIsWhitelistMember(r.Context(), minterResp.Whitelist, walletAddr)
				if err != nil {
					memberChan <- &MemberResponse{
						Minter:     resp.Minter,
						Collection: coll.CollectionAddr,
						Name:       coll.Name,
						Image:      coll.Image,
					}
					return
				}
				memberChan <- &MemberResponse{
					Minter:     resp.Minter,
					Collection: coll.CollectionAddr,
					Name:       coll.Name,
					Image:      coll.Image,
					HasMember:  wlResp.IsMember,
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

		WriteJSON(w, http.StatusOK, members)

	})

	slog.Info(fmt.Sprintf("starting server on port %v...", port))
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), r)
	if err != nil {
		panic(fmt.Sprintf("error starting server on port %v", port))
	}

}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
