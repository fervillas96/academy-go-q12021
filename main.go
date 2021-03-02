package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func createPokedexMap(pokedexArray [][]string) map[int]string {
	var pokedexmap = make(map[int]string)

	for index, element := range pokedexArray {
		if index != 0 {
			pokemonId, err := strconv.Atoi(element[0])
			if err != nil {
				fmt.Println(err)
			}
			pokedexmap[pokemonId] = element[1]
		}
	}

	return pokedexmap

}

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Home Page Hit")
}

func getPokemonById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var csvData = ReadCsvFile("/Users/jorge.villanueva/Downloads/pokedex.csv")

	var pokedex = createPokedexMap(csvData)
	fmt.Printf("%v", pokedex[1])
	fmt.Printf("%v", pokedex[2])
	fmt.Printf("%v", pokedex[5687])
}

func handleRequest() {
	router := httprouter.New()

	router.GET("/", homePage)
	router.GET("/getname/:id", getPokemonById)

	log.Fatal(http.ListenAndServe(":4000", router))
}

func main() {
	handleRequest()
}
