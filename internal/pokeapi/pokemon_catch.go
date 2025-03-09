package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(foundPokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + foundPokemon

	// Fetch data from the API
	req, err :=  http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return pokemon, err
	}

	return pokemon, nil
}