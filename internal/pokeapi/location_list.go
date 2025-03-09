package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// List Locations
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// If cache exists for the request use it and return
	if cachedData, found := c.Cache.Get(url); found {
		var locationsResp RespShallowLocations
		err := json.Unmarshal(cachedData, &locationsResp)
		return locationsResp, err
	}

	// Otherwise fetch the data from the API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// add the response data to the cache
	c.Cache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}