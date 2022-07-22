package pokemons

import (
	"github.com/pedrohsilvaml/pokemon-api/internal/client/poke_api"
)

type PokemonService struct{}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

type GetPokemonResponse struct {
	Data    interface{}
	Partial bool
}

func (ps *PokemonService) GetPokemon(name string) (*GetPokemonResponse, error) {
	response, err := poke_api.GetPokemon(name)
	partial := (err != nil)
	data := response.Data

	if partial {
		data = &poke_api.PokeAPIData{Name: name}
	}

	return &GetPokemonResponse{
		Data:    data,
		Partial: partial,
	}, err
}
