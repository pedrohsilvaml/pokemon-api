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

	pokemonCtl := BuildPokemonController()

	sv.Engine.GET("/pokemon", pokemonCtl.GetPokemon)
	sv.Engine.GET("/pokemons/initial", pokemonCtl.GetInitialPokemons)
}

func (sv *Server) Start() {
	sv.Engine.Run()
}

func BuildPokemonController() *pokemons.PokemonController {
	pokeAPIClient := poke_api.NewClient()
	pokemonService := poke_service.NewPokemonService(pokeAPIClient)

	return pokemons.NewPokemonController(pokemonService)
}
