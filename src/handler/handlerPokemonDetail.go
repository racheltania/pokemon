package handler

import (
	"github.com/labstack/echo/v4"
	"miniProject/src/service"
	"net/http"
	"strconv"
)

type (
	PokemonDetailHandler interface {
		GetPokemonDetail(ctx echo.Context) error
		CatchPokemon(ctx echo.Context) error
		HitRename(ctx echo.Context) error
	}

	PokemonDetailHandlerImpl struct {
		pokemonDetailService service.PokemonDetailService
	}
)

func NewPokemonDetailHandler(p service.PokemonDetailService) PokemonDetailHandler {
	return &PokemonDetailHandlerImpl{pokemonDetailService: p}
}

func (p *PokemonDetailHandlerImpl) GetPokemonDetail(ctx echo.Context) error {
	pokemonName := ctx.Param("pokemon")
	results, err := p.pokemonDetailService.DetailPokemon(pokemonName)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, results)
}

func (p *PokemonDetailHandlerImpl) CatchPokemon(ctx echo.Context) error {
	results := p.pokemonDetailService.CatchPokemon()
	return ctx.JSON(http.StatusOK, results)
}

func (p *PokemonDetailHandlerImpl) HitRename(ctx echo.Context) error {
	name := ctx.QueryParam("name")
	hitRename := ctx.QueryParam("hit")
	intHit, _ := strconv.Atoi(hitRename)
	results := p.pokemonDetailService.RenamePokemon(name, intHit)
	return ctx.JSON(http.StatusOK, results)
}
