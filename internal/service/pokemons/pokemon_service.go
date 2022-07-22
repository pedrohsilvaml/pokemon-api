package pokemons

import (
	"github.com/pedrohsilvaml/pokemon-api/internal/client/poke_api"
)

type PokemonService struct{}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

type GetPokemonResponse struct {
	Data    interface{} `json:"data"`
	Partial bool        `json:"partial"`
}

func (PokemonService) GetPokemon(name string) (*GetPokemonResponse, error) {
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

func (ps PokemonService) GetInitialPokemons() (*[]*GetPokemonResponse, error) {
	pokemon_names := getInitialPokemons()
	pokemons := make([]*GetPokemonResponse, len(pokemon_names))

	for i, name := range pokemon_names {
		response, _ := ps.GetPokemon(name)
		pokemons[i] = response
	}

	return &pokemons, nil
}

func getInitialPokemons() []string {
	return []string{"charmander", "squirtle", "bulbasaur"}
}
