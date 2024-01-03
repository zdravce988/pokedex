package main

import (
    "fmt"
)

func commandHelp(cfg *config, arguments []string) error {
    commands := commands()
    fmt.Println()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
    for _, command := range commands {
        fmt.Printf("%s: %s\n", command.name, command.description)
    }
    fmt.Println()
    return nil
}
