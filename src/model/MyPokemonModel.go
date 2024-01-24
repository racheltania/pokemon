package model

type MyPokemon struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type AddMyPokemonRequest struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
