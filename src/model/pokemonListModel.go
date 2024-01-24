package model

type (
	PokemonList struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	PokemonAPIResponse struct {
		Results []*PokemonList `json:"results"`
	}
)
