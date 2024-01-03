package pokeapi

import (
    "encoding/json"
)

func (c *Client) GetLocations(url *string) (Result, error) {
    endpoint := "location-area"
    fullURL := baseURL + endpoint

    if url != nil {
        fullURL = *url
    }

    data, err := c.getData(fullURL)
    if err != nil {
        return Result{}, err 
    }

    result := Result{}
    err = json.Unmarshal(data, &result)

    if err != nil {
        return Result{}, err
    }

    return result, nil 
}

func (c *Client) ExploreLocations(area string) (LocationAreaPokemons, error) {
    endpoint := "location-area/"
    fullURL := baseURL + endpoint + area

    data, err := c.getData(fullURL)
    if err != nil {
        return LocationAreaPokemons{}, err 
    }

    pokemons := LocationAreaPokemons{}
    err = json.Unmarshal(data, &pokemons)

    if err != nil {
        return LocationAreaPokemons{}, err
    }

    return pokemons, nil
}
