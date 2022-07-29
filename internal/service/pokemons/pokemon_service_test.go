package pokemons

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/pedrohsilvaml/pokemon-api/internal/client/poke_api"
	"github.com/pedrohsilvaml/pokemon-api/internal/client/poke_api/mock"
)

type getPokemonTest struct {
	ID      int
	Name    string
	Partial bool
}

var getPokemonTests = []getPokemonTest{
	{ID: 4, Name: "charmander", Partial: false},
	{ID: 0, Name: "xxx", Partial: true},
	{ID: 0, Name: "", Partial: true},
}

func TestGetPokemonTableDriven(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, test := range getPokemonTests {
		data := poke_api.PokeAPIData{ID: test.ID, Name: test.Name}
		var errorResponse error
		if test.ID == 0 {
			errorResponse = errors.New("not found")
		}

		mockPokeClient := mock.NewMockPokeAPI(ctrl)
		gomock.InOrder(
			mockPokeClient.EXPECT().GetPokemon(data.Name).Return(&data, errorResponse),
		)

		service := NewPokemonService(mockPokeClient)
		response, _ := service.GetPokemon(test.Name)

		if response.Partial != test.Partial {
			t.Errorf("[pokemon_service] Error: expect partial equal to %v but got %v for test: %v", test.Partial, response.Partial, test)
		}

		if response.Data.ID != test.ID {
			t.Errorf("[pokemon_service] Error: expect ID equal to %v but got %v for test: %v", test.ID, response.Data.ID, test)
		}

		if response.Data.Name != test.Name {
			t.Errorf("[pokemon_service] Error: expect name equal to %v but got %v for test: %v", test.Name, response.Data.Name, test)
		}
	}
}

var getInitialPokemonsTestResponse = []poke_api.PokeAPIData{
	{ID: 4, Name: "charmander"},
	{ID: 7, Name: "squirtle"},
	{ID: 1, Name: "bulbasaur"},
}

func TestGetInitialPokemons(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPokeClient := mock.NewMockPokeAPI(ctrl)
	gomock.InAnyOrder(mockPokeClient.EXPECT().GetPokemon(getInitialPokemonsTestResponse[0].Name).Return(&getInitialPokemonsTestResponse[0], nil))
	gomock.InAnyOrder(mockPokeClient.EXPECT().GetPokemon(getInitialPokemonsTestResponse[1].Name).Return(&getInitialPokemonsTestResponse[1], nil))
	gomock.InAnyOrder(mockPokeClient.EXPECT().GetPokemon(getInitialPokemonsTestResponse[2].Name).Return(&getInitialPokemonsTestResponse[2], nil))

	service := NewPokemonService(mockPokeClient)
	response := service.GetInitialPokemons()

	for _, pokeData := range response {
		if pokeData.Data.ID == 0 {
			t.Errorf("[pokemon_service] Error: expect ID to be present but got blank for data: %v", pokeData.Data)
		}

		if len(pokeData.Data.Name) == 0 {
			t.Errorf("[pokemon_service] Error: expect name to be present but got blank for data: %v", pokeData.Data)
		}

		if pokeData.Partial {
			t.Errorf("[pokemon_service] Error: expect partial to be false but got true for data: %v", pokeData.Data)
		}
	}
}
