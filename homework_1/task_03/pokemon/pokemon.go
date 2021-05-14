package pokemon

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
)

const pokemonAPI = "https://pokeapi.co/api/v2/pokemon/"

type Location struct {
	Name string
}

type Encounter struct {
	Location Location `json:"location_area"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type Output struct {
	Name string
	Locations []string
}

func GetData(url string) ([]byte, error) {
	httpClient := pester.New()

	httpResponse, err := httpClient.Get(url)
	if err != nil {
		return nil, errors.New("HTTP get towards pokeapi")
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, errors.New("error while reading pokeapi")
	}

	return bodyContent, nil
}

func FindLocations(input string) ([]byte, error) {
	data, err := GetData(pokemonAPI + input)
	if err != nil {
		return nil, err
	}
	var decoded Pokemon
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		return nil, errors.New("unmarshalling the JSON body content")
	}

	data, err = GetData(pokemonAPI + input + "/encounters")
	if err != nil {
		return nil, err
	}
	var areas []Encounter
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return nil, errors.New("unmarshalling the JSON body content")
	}


	var output Output
	output.Name = decoded.Name
	for _, value := range areas {
		output.Locations = append(output.Locations, value.Location.Name)
	}

	result, _ := json.MarshalIndent(output, "", "\t")
	return result, err
}