package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you are on first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locations.Next
	cfg.prevLocationsURL = locations.Previous

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
