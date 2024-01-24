package model

type PokemonDetail struct {
	Name    string        `json:"name"`
	Picture PictureDetail `json:"sprites"`
	Moves   []Moves       `json:"moves"`
	Types   []Types       `json:"types"`
}

type PictureDetail struct {
	FrontDefault string `json:"front_default"`
}

type Moves struct {
	Move Move `json:"move"`
}

type Move struct {
	Name string `json:"name"`
}

type Types struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}

type PokemonDetailResponse struct {
	Name    string   `json:"name"`
	Picture string   `json:"sprites"`
	Moves   []string `json:"moves"`
	Types   []string `json:"types"`
}
