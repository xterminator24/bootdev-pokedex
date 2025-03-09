package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Explore(locationName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + locationName

	// If cache exists for the request use it and return
	if cachedData, found := c.Cache.Get(url); found {
		fmt.Println("using cache...")
		var locationArea	RespLocationArea
		err := json.Unmarshal(cachedData, &locationArea)
		return locationArea, err
	}
	
	// Otherwise fetch the data from the API
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}

	// add the response data to the cache
	c.Cache.Add(url, dat)

	locationAreaResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return RespLocationArea{}, err
	}

	return locationAreaResp, nil
}