package main

import (
    "fmt"
)

func commandInspect(cfg *config, arguments []string) error {
    pokemonName := arguments[0]

    pokemon, exists := cfg.pokedex[pokemonName]
    if !exists {
        return fmt.Errorf("%s not caught!", pokemonName)
    }

    fmt.Println("Name:", pokemon.Name)
    fmt.Println("Height:", pokemon.Height)
    fmt.Println("Weight:", pokemon.Weight)

    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, t := range pokemon.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }

    return nil
}
