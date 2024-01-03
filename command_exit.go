package main

import (
    "os"
)

func commandExit(cfg *config, arguments []string) error {
    os.Exit(0)
    return nil
}
