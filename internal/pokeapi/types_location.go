package pokeapi

// RespLocationArea
type RespLocationArea struct {
	ID					int		`json:"id"`
	GameIndex			int 	`json:"game_index"`
	Name				string  `json:"name"`
	PokemonEncounters 	[]PokemonEncounter	`json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon		Pokemon	`json:"pokemon"`
}

type Pokemon struct {
	Name		string	`json:"name"`
	URL			string  `json:"url"`
}