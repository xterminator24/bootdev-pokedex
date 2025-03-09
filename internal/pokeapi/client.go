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
}

// New Client
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		Cache:		pokecache.NewCache(cacheInterval),
	}
}