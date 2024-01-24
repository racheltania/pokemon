package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"miniProject/src/model"
	"net/http"
	"strconv"
)

type (
	// PokemonService
	PokemonDetailService interface {
		DetailPokemon(pokemonName string) (*model.PokemonDetailResponse, error)
		CatchPokemon() bool
		RenamePokemon(pokemonName string, hitRename int) string
	}

	PokemonDetailServiceImpl struct {
	}
)

func NewPokemonDetailService() PokemonDetailService {
	return &PokemonDetailServiceImpl{}
}

func (*PokemonDetailServiceImpl) DetailPokemon(pokemonName string) (*model.PokemonDetailResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
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

	var pokemonApiResponse *model.PokemonDetail
	errApi := json.NewDecoder(resp.Body).Decode(&pokemonApiResponse)
	if errApi != nil {
		fmt.Printf("Error decoding response into JSON: %s\n", err)
		return nil, err
	}

	var moveNames = []string{}
	for _, move := range pokemonApiResponse.Moves {
		moveNames = append(moveNames, move.Move.Name)
	}
	var typeNames = []string{}
	for _, typePoke := range pokemonApiResponse.Types {
		typeNames = append(typeNames, typePoke.Type.Name)
	}
	pokeResponse := model.PokemonDetailResponse{
		Name:    pokemonApiResponse.Name,
		Picture: pokemonApiResponse.Picture.FrontDefault,
		Types:   typeNames,
		Moves:   moveNames,
	}
	return &pokeResponse, nil
}

func (*PokemonDetailServiceImpl) CatchPokemon() bool {
	randomNumber := rand.Intn(2)

	if randomNumber == 0 {
		return true
	} else {
		return false
	}
}

func (i *PokemonDetailServiceImpl) RenamePokemon(pokemonName string, hitRename int) string {
	newName := pokemonName + "-"
	if hitRename <= 1 {
		newName = newName + strconv.Itoa(hitRename)
		return newName
	}

	a, b := 0, 1
	for i := 2; i <= hitRename; i++ {
		a, b = b, a+b
	}

	newName = newName + strconv.Itoa(b)

	return newName
}
