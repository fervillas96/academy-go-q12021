package main

import "fmt"

// Pokemon struct for creating pokemons
type Pokemon struct {
	name   string
	attack string
}

// Attack function for a pokemon to attack
func Attack(p Pokemon) string {
	return fmt.Sprintf("%s use %s!", p.name, p.attack)
}

// CreatePokemon creates a new Pokemon
func CreatePokemon(name, attack string) Pokemon {
	p := Pokemon{name: name, attack: attack}
	return p
}
