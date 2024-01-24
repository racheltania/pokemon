package handler

import (
	"github.com/labstack/echo/v4"
	"miniProject/src/service"
	"miniProject/src/utils"
	"net/http"
)

type (
	PokemonListHandler interface {
		GetPokemonList(ctx echo.Context) error
	}

	PokemonListHandlerImpl struct {
		pokemonListService service.PokemonListService
	}
)

func NewPokemonListHandler(p service.PokemonListService) PokemonListHandler {
	return &PokemonListHandlerImpl{pokemonListService: p}
}

func (p *PokemonListHandlerImpl) GetPokemonList(ctx echo.Context) error {
	paramQuery := utils.SetParamQuery(ctx)
	results, err := p.pokemonListService.GetPokemon(*paramQuery)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, results)
}
