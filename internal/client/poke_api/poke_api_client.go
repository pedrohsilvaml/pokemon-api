package poke_api

import (
	"encoding/json"
	"net/http"
)

const BaseUri = "https://pokeapi.co/api/v2"

type PokeAPIData struct {
	ID                     int    `json:"id,omitempty"`
	Name                   string `json:"name,omitempty"`
	Order                  int    `json:"order,omitempty"`
	Weight                 int    `json:"weight,omitempty"`
	Height                 int    `json:"height,omitempty"`
	IsDefault              bool   `json:"is_default,omitempty"`
	BaseExperience         int    `json:"base_experience,omitempty"`
	LocationAreaEncounters string `json:"location_area_encounters,omitempty"`
	Species                struct {
		URL  string `json:"url,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"species,omitempty"`
	Forms []struct {
		URL  string `json:"url,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"forms,omitempty"`
	Stats []struct {
		Effort   int `json:"effort,omitempty"`
		BaseStat int `json:"base_stat,omitempty"`
		Stat     struct {
			URL  string `json:"url,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"stat,omitempty"`
	} `json:"stats,omitempty"`
	Types []struct {
		Slot int `json:"slot,omitempty"`
		Type struct {
			URL  string `json:"url,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"type,omitempty"`
	} `json:"types,omitempty"`
}

type HttpResponse struct {
	Data       *PokeAPIData
	StatusCode int
}

type PokeAPIClient struct{}

func NewPokeAPIClient() PokeAPIClient {
	return PokeAPIClient{}
}

func (client PokeAPIClient) GetPokemon(name string) (*HttpResponse, error) {
	url := client.buildUrl("/pokemon/" + name)
	req, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	var data PokeAPIData
	err = json.NewDecoder(req.Body).Decode(&data)

	return &HttpResponse{Data: &data, StatusCode: req.StatusCode}, err
}

func (client PokeAPIClient) buildUrl(path string) string {
	return BaseUri + path
}
