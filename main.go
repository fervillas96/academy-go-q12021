package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const file = "/Users/jorge.villanueva/Downloads/pokedex.csv"

var pokedex = make(map[int]string)

func fillPokedex(pokedexArray [][]string) {

	for index, element := range pokedexArray {
		if index != 0 {
			pokemonID, err := strconv.Atoi(element[0])
			if err != nil {
				fmt.Println(err)
			}
			pokedex[pokemonID] = element[1]
		}
	}
}

func createPokedex() {
	var csvData = ReadCsvFile(file)
	fillPokedex(csvData)
}

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Pokedex API!\nAvailable endpoints:\n -> /getname/:id - for retrieve a pokemon name")
}

func getPokemonByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pokemonID, err := strconv.Atoi(ps.ByName("id"))

	if err == nil {
		fmt.Fprint(w, pokedex[pokemonID])
	}
}

func handleRequest() {
	router := httprouter.New()

	router.GET("/", homePage)
	router.GET("/getname/:id", getPokemonByID)

	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	createPokedex()

	handleRequest()
}
