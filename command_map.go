package main

import (
    "fmt"
)

func commandMap(cfg *config, arguments []string) error {
    result, err := cfg.pokeapiClient.GetLocations(cfg.nextLocationAreaURL)
    if err != nil {
        return err
    }

    for _, value := range result.Results {
        fmt.Println(value.Name)
    }
    cfg.nextLocationAreaURL = result.Next
    cfg.previousLocationAreaURL = result.Previous
    return nil
}

func commandMapb(cfg *config, arguments []string) error {
    result,err := cfg.pokeapiClient.GetLocations(cfg.previousLocationAreaURL)
    if err != nil {
        return err
    }

    for _, value := range result.Results {
        fmt.Println(value.Name)
    }
    cfg.nextLocationAreaURL = result.Next
    cfg.previousLocationAreaURL = result.Previous
    return nil
}
