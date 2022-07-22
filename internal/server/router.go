package server

import (
	"github.com/gin-gonic/gin"

	"github.com/pedrohsilvaml/pokemon-api/internal/server/pokemons"
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

	pokemonClt := pokemons.NewPokemonController()

	sv.Engine.GET("/pokemon", pokemonClt.GetPokemon)
	sv.Engine.GET("/pokemons/initial", pokemonClt.GetInitialPokemons)
}

func (sv *Server) Start() {
	sv.Engine.Run()
}
