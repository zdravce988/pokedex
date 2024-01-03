package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)


type cliCommand struct {
    name string
    description string
    callback func(*config, []string) error
}

func commands() map[string]cliCommand {
    return map[string]cliCommand {
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "map": {
            name: "map",
            description: "Displays the names of 20 location areas in the Pokemon world",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb",
            description: "Displays previous 20 names of location areas in the Pokemon world",
            callback: commandMapb,
        },
        "explore": {
            name: "explore",
            description: "Explore the pokemons in the given area",
            callback: commandExplore,
        },
        "catch": {
            name: "catch",
            description: "Command to catch a pokemon",
            callback: commandCatch,
        },
        "inspect": {
            name: "inspect",
            description: "Inspect the pokemon in your deck",
            callback: commandInspect,
        },
        "pokedex": {
            name: "pokedex",
            description: "Shows the pokemons in your deck",
            callback: commandPokedex,
        },
    }
}

func cleanInput(input string) (string, []string) {
    contents := strings.ToLower(strings.TrimSpace(input))
    args := strings.Fields(contents) 
    command := args[0]
    args = args[1:]
    return command, args
}

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
    commands := commands()

    fmt.Print("pokedex> ")
    for scanner.Scan() {
        command, args := cleanInput(scanner.Text())

        c, exists := commands[command]
        if !exists {
            fmt.Println("Unknown command")
        } else {
            err := c.callback(cfg, args)
            if err != nil {
                fmt.Println(err)
            }
        }
        fmt.Print("pokedex> ")
    }
}
