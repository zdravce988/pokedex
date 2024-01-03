package pokeapi

import (
    "encoding/json"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
    endpoint := "pokemon/"
    fullURL := baseURL + endpoint + pokemonName

    data, err := c.getData(fullURL)
    if err != nil {
        return Pokemon{}, err 
    }

    pokemon := Pokemon{}
    err = json.Unmarshal(data, &pokemon)
    if err != nil {
        return Pokemon{}, err 
    }

    return pokemon, nil
}
