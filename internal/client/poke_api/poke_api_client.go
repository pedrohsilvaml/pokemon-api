package poke_api

import (
	"encoding/json"

	"github.com/pedrohsilvaml/pokemon-api/pkg/request"
)

const BaseUri = "https://pokeapi.co/api/v2"

type PokeAPIClient struct{}

func NewClient() PokeAPIClient {
	return PokeAPIClient{}
}

func (client PokeAPIClient) GetPokemon(name string) (*PokeAPIData, error) {
	url := client.buildUrl("/pokemon/" + name)

	response, err := request.Get(url)

	if err != nil {
		return nil, err
	}

	var data PokeAPIData
	err = json.Unmarshal(response.Body, &data)

	return &data, err
}

func (client PokeAPIClient) buildUrl(path string) string {
	return BaseUri + path
}
