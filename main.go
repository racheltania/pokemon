package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"miniProject/db"
	"miniProject/middlewares"
	"miniProject/src/handler"
	"miniProject/src/repository"
	"miniProject/src/route"
	"miniProject/src/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	db.InitDB()
	//list of Repos
	myPokemonRepo := repository.NewMyPokemonRepo(db.DB)
	// list of services
	pokemonListService := service.NewPokemonListService()
	pokemonDetailService := service.NewPokemonDetailService()
	myPokemonService := service.NewMyPokemonService(myPokemonRepo)
	//list of handlers
	pokemonListHandler := handler.NewPokemonListHandler(pokemonListService)
	pokemonDetailHandler := handler.NewPokemonDetailHandler(pokemonDetailService)
	myPokemonHandler := handler.NewMyPokemonHandler(myPokemonService)

	e := echo.New()
	e.Use(middlewares.EnableCORS())
	route.Route(
		e,
		pokemonListHandler,
		pokemonDetailHandler,
		myPokemonHandler,
	)
	e.Logger.Fatal(e.Start(":1323"))
}
