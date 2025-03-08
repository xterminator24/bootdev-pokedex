package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count		int		`json:"count"`
	Next		string	`json:"next"`
	Previous	string	`json:"previous"`
	Results		[]struct {
		Name string `json:"name"`
		URL	 string `json:"url"`
	} `json:"results"`
}

func getLocationAreas(url string) (LocationAreas, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	res, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreas{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	var locationAreas LocationAreas
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreas{}, err
	}
	
	return locationAreas, nil
}

func printLocationAreas(locationAreas LocationAreas) {
	for _,location := range locationAreas.Results {
		fmt.Println(location.Name)
	}
}