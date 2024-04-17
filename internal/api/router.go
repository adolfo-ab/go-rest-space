package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", HomePage)
	router.HandleFunc("/planets", ReturnAllPlanets)
	router.HandleFunc("/planet", CreateNewPlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", ReturnSinglePlanet).Methods("GET")
	router.HandleFunc("/planet/{id}", DeletePlanet).Methods("DELETE")
	router.HandleFunc("/planet/{id}", UpdatePlanet).Methods("PUT")

	log.Fatalln(http.ListenAndServe(":10000", router))
}
