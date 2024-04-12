package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/internal/domain/planet"
	"main/internal/interfaces"
	"net/http"
)

var Planets []planet.Planet

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", interfaces.HomePage)
	router.HandleFunc("/planets", interfaces.ReturnAllPlanets)
	router.HandleFunc("/planet", interfaces.CreateNewPlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", interfaces.ReturnSinglePlanet).Methods("GET")
	router.HandleFunc("/planet/{id}", interfaces.DeletePlanet).Methods("DELETE")
	router.HandleFunc("/planet/{id}", interfaces.UpdatePlanet).Methods("PUT")

	log.Fatalln(http.ListenAndServe(":10000", router))
}

func main() {
	handleRequests()
}
