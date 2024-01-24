package service

import (
	"encoding/json"
	"fmt"
	"miniProject/src/model"
	"miniProject/src/utils"
	"net/http"
)

type (
	// PokemonService
	PokemonListService interface {
		GetPokemon(paramQuery utils.ParamQuery) ([]*model.PokemonList, error)
	}

	PokemonListServiceImpl struct {
	}
)

func NewPokemonListService() PokemonListService {
	return &PokemonListServiceImpl{}
}

func (p *PokemonListServiceImpl) GetPokemon(paramQuery utils.ParamQuery) ([]*model.PokemonList, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/?offset=%v&limit=%v", paramQuery.Offset, paramQuery.Limit)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemonApiResponse *model.PokemonAPIResponse
	errApi := json.NewDecoder(resp.Body).Decode(&pokemonApiResponse)
	if errApi != nil {
		return nil, err
	}
	return pokemonApiResponse.Results, nil
}
