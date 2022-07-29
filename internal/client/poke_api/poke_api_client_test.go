package poke_api

import "testing"

type getPokemonTest struct {
	ID   int
	Name string
}

var getPokemonTests = []getPokemonTest{
	{ID: 4, Name: "charmander"},
	{ID: 0, Name: "xxx"},
	{ID: 0, Name: ""},
}

func TestGetPokemonTableDriven(t *testing.T) {
	client := NewClient()

	for _, test := range getPokemonTests {
		response, err := client.GetPokemon(test.Name)

		if test.ID != 0 && err != nil {
			t.Errorf("[poke_api_client] Error: %s\n for test: %v", err, test)
		}
		if response.ID != test.ID {
			t.Errorf("[poke_api_client] Error: invalid pokemon for test: %v", test)
		}
	}
}

func BenchmarkGetPokemon(b *testing.B) {
	client := NewClient()

	for i := 0; i < b.N; i++ {
		client.GetPokemon("charmander")
	}
}
