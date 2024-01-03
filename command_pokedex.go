package main

import (
    "fmt"
)

func commandPokedex(cfg *config, arguments []string) error {
    fmt.Println("Your Pokedex:")
    for name := range cfg.pokedex {
        fmt.Printf("  - %s\n", name)
    }
    return nil
}
