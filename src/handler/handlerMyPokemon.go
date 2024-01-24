package handler

import (
	"github.com/labstack/echo/v4"
	"miniProject/src/model"
	"miniProject/src/service"
	"miniProject/src/utils"
	"net/http"
	"strconv"
)

type (
	MyPokemonHandler interface {
		MyPokemonList(ctx echo.Context) error
		AddMyPokemon(ctx echo.Context) error
		ReleasePokemon(ctx echo.Context) error
	}

	MyPokemonHandlerImpl struct {
		myPokemonService service.MyPokemonService
	}
)

func NewMyPokemonHandler(p service.MyPokemonService) MyPokemonHandler {
	return &MyPokemonHandlerImpl{myPokemonService: p}
}

func (p *MyPokemonHandlerImpl) MyPokemonList(ctx echo.Context) error {
	paramQuery := utils.SetParamQuery(ctx)
	results, err := p.myPokemonService.MyPokemonList(ctx.Request().Context(), *paramQuery)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, results)
}

func (p *MyPokemonHandlerImpl) AddMyPokemon(ctx echo.Context) error {
	requestPoke := new(model.AddMyPokemonRequest)
	if err := ctx.Bind(requestPoke); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}
	err := p.myPokemonService.AddMyPokemon(ctx.Request().Context(), *requestPoke)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, "ok")
}

func (p *MyPokemonHandlerImpl) ReleasePokemon(ctx echo.Context) error {
	paramId := ctx.Param("id")
	id, _ := strconv.Atoi(paramId)
	results, err := p.myPokemonService.ReleasePokemon(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, results)
}
