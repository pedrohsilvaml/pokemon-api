package server

import (
	"github.com/gin-gonic/gin"

	"github.com/pedrohsilvaml/pokemon-api/internal/client/poke_api"
	"github.com/pedrohsilvaml/pokemon-api/internal/server/pokemons"
	poke_service "github.com/pedrohsilvaml/pokemon-api/internal/service/pokemons"
)

type Config struct {
}

type Server struct {
	*Config
	Engine *gin.Engine
}

func NewServer(cfg *Config) *Server {
	return &Server{Config: cfg}
}

func (sv *Server) Setup() {
	sv.Engine = gin.Default()

	pokeAPIClient := poke_api.NewPokeAPIClient()
	pokemonService := poke_service.NewPokemonService(pokeAPIClient)
	pokemonClt := pokemons.NewPokemonController(pokemonService)

	sv.Engine.GET("/pokemon", pokemonClt.GetPokemon)
	sv.Engine.GET("/pokemons/initial", pokemonClt.GetInitialPokemons)
}

func (sv *Server) Start() {
	sv.Engine.Run()
}
