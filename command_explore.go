package main

import (
    "fmt"
)

func commandExplore(cfg *config, arguments []string) error {
    if len(arguments) < 1 {
        return fmt.Errorf("Inalid number of arguments")
    }
    pokemons, err := cfg.pokeapiClient.ExploreLocations(arguments[0])
    if err != nil {
        return err
    }
    fmt.Printf("Exploring %s...\n", arguments[0])
    fmt.Println("Found Pokemon:")
    for _, pokemon := range pokemons.PokemonEncounters {
        fmt.Println(" -",  pokemon.Pokemon.Name)
    }

    return nil
}
