package main

import(
    "fmt"
    "math/rand"
)

func commandCatch(cfg *config, arguments []string) error {
    pokemonName := arguments[0]

    pokemon, exists := cfg.pokedex[pokemonName]
    if exists {
        fmt.Println("You've already caught", pokemonName)
        return nil
    }

    pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
    if err != nil {
        return err
    }
    fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
    base := int(100 * (1  - (1 - float64(pokemon.BaseExperience) / 1000)))
    seed := rand.Intn(100) + 1 
    if seed >= base {
        fmt.Println(pokemonName, "was caught!")
        cfg.pokedex[pokemonName] = pokemon
    } else {
        fmt.Println(pokemonName, "escaped!")
    }

    return nil
}
