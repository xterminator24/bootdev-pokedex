package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	// If cache exists for the request use it and return
	if cachedData, found := c.Cache.Get(url); found {
		fmt.Println("using cache...")
		var location	Location
		err := json.Unmarshal(cachedData, &location)
		return location, err
	}
	
	// Otherwise fetch the data from the API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	// add the response data to the cache
	c.Cache.Add(url, dat)

	location := Location{}
	err = json.Unmarshal(dat, &location)
	if err != nil {
		return Location{}, err
	}

	return location, nil
}