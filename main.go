package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Home Page Hit")
}

func getPokemonById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var pokedex = ReadCsvFile("/Users/jorge.villanueva/Downloads/pokedex.csv")
	fmt.Println(pokedex)
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
