package main

import (
	"github.com/helmigandi/learn-go-pokedex-cli/internal/pokeapi"
	"time"
)

func main() {
	client := pokeapi.NewClient(15 * time.Second)
	cfg := &config{
		pokeapiClient: client,
	}
	startRepl(cfg)
}
