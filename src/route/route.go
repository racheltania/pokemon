package route

import (
	"github.com/labstack/echo/v4"
	"miniProject/src/handler"
)

func Route(
	e *echo.Echo,
	PokemonListHandler handler.PokemonListHandler,
	PokemonDetailHandler handler.PokemonDetailHandler,
	MyPokemonHandler handler.MyPokemonHandler,
) {
	e.GET("/", PokemonListHandler.GetPokemonList)
	e.GET("/:pokemon", PokemonDetailHandler.GetPokemonDetail)
	e.GET("/catch", PokemonDetailHandler.CatchPokemon)
	e.DELETE("/release/:id", MyPokemonHandler.ReleasePokemon)
	e.GET("/my", MyPokemonHandler.MyPokemonList)
	e.GET("/rename", PokemonDetailHandler.HitRename)
	e.POST("/add", MyPokemonHandler.AddMyPokemon)

}
