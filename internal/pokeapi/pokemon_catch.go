package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func (c *Client) CatchPokemon(foundPokemon string) (Pokemon, error) {
	pokemon, err := c.GetPokemon(foundPokemon)
	if err != nil {
		return Pokemon{}, err
	}

	// Generate a random number between 0 and 100
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	randomInt := r.Intn(101) // random integer between 0 and 100

	// if the random roll is less than or equal to the base experience of the pokemon, it escapes
	if randomInt <= pokemon.BaseExperience {
		return Pokemon{}, fmt.Errorf("%s excaped", foundPokemon)
	}

	return pokemon, nil
}

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