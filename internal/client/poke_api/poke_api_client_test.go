package poke_api

import "testing"

type getPokemonTest struct {
	Name string
	ID   int
}

var getPokemonTests = []getPokemonTest{
	{"charmander", 4},
	{"xxx", 0},
	{"", 0},
}

func TestGetPokemonTableDriven(t *testing.T) {
	client := NewClient()

	for _, test := range getPokemonTests {
		response, err := client.GetPokemon(test.Name)

		if test.ID != 0 && err != nil {
			t.Errorf("[GetPokemon] Error: %s\n for test: %v", err, test)
		}
		if response.ID != test.ID {
			t.Errorf("[GetPokemon] Error: invalid pokemon for test: %v", test)
		}
	}
}

func BenchmarkGetPokemon(b *testing.B) {
	client := NewClient()

	for i := 0; i < b.N; i++ {
		client.GetPokemon("charmander")
	}
}
