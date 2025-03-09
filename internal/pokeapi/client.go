package pokeapi

import (
	"net/http"
	"time"

	"github.com/xterminator24/bootdev-pokedex/internal/pokecache"
)

// Client
type Client struct {
	httpClient 	http.Client
	Cache		*pokecache.Cache
	Pokedex     map[string]Pokemon
}

// New Client
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache:		pokecache.NewCache(timeout),
		Pokedex:    map[string]Pokemon{},
	}
}