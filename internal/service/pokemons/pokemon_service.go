package pokemons

import (
	"sync"

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

func (s PokemonService) GetInitialPokemons() []GetPokemonResponse {
	pokemon_names := getInitialPokemons()
	jobs := len(pokemon_names)
	var pokemons []GetPokemonResponse

	channel := make(chan GetPokemonResponse, jobs)
	var waitGroup sync.WaitGroup

	for _, name := range pokemon_names {
		waitGroup.Add(1)
		go s.getPokemonJob(name, channel, &waitGroup)
	}
	waitGroup.Wait()
	close(channel)

	for pokemonResponse := range channel {
		pokemons = append(pokemons, pokemonResponse)
	}

	return pokemons
}

func (s PokemonService) getPokemonJob(name string, channel chan GetPokemonResponse, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	defer s.getPokemonJobRecover(name, channel)

	response, _ := s.GetPokemon(name)
	channel <- *response
}

func (PokemonService) getPokemonJobRecover(name string, channel chan GetPokemonResponse) {
	if r := recover(); r != nil {
		data := &poke_api.PokeAPIData{Name: name}
		channel <- GetPokemonResponse{Data: data, Partial: true}
	}
}

func getInitialPokemons() []string {
	return []string{"charmander", "squirtle", "bulbasaur"}
}
