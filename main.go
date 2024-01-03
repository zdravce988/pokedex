package main

import (
    "github.com/zdravce988/pokedex/internal/pokeapi"
)

type config struct {
    pokeapiClient pokeapi.Client
    nextLocationAreaURL *string
    previousLocationAreaURL *string
    pokedex map[string]pokeapi.Pokemon
}

func main() {
    cfg := config {
        pokeapiClient: pokeapi.NewClient(),
        pokedex: make(map[string]pokeapi.Pokemon),
    }

    startRepl(&cfg)
}


