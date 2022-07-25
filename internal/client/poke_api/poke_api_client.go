package poke_api

import (
	"encoding/json"
	"net/http"
)

const BaseUri = "https://pokeapi.co/api/v2"

type PokeAPIData struct {
	ID                     int         `json:"id,omitempty"`
	Name                   string      `json:"name,omitempty"`
	LocationAreaEncounters string      `json:"location_area_encounters,omitempty"`
	Weight                 int         `json:"weight,omitempty"`
	Height                 int         `json:"height,omitempty"`
	Order                  int         `json:"order,omitempty"`
	BaseExperience         int         `json:"base_experience,omitempty"`
	IsDefault              bool        `json:"is_default,omitempty"`
	Abilities              []Abilities `json:"abilities,omitempty"`
	Species                *Species    `json:"species,omitempty"`
	Sprites                *Sprites    `json:"sprites,omitempty"`
	Types                  []Types     `json:"types,omitempty"`
}
type Ability struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Abilities struct {
	Ability  Ability `json:"ability,omitempty"`
	IsHidden bool    `json:"is_hidden,omitempty"`
	Slot     int     `json:"slot,omitempty"`
}
type Species struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Sprites struct {
	BackDefault      string      `json:"back_default,omitempty"`
	BackFemale       interface{} `json:"back_female,omitempty"`
	BackShiny        string      `json:"back_shiny,omitempty"`
	BackShinyFemale  interface{} `json:"back_shiny_female,omitempty"`
	FrontDefault     string      `json:"front_default,omitempty"`
	FrontFemale      interface{} `json:"front_female,omitempty"`
	FrontShiny       string      `json:"front_shiny,omitempty"`
	FrontShinyFemale interface{} `json:"front_shiny_female,omitempty"`
}
type Type struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
type Types struct {
	Slot int  `json:"slot,omitempty"`
	Type Type `json:"type,omitempty"`
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
