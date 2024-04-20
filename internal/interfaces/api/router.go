package api

import (
	"github.com/gorilla/mux"
	"log"
	"main/internal/interfaces/api/handlers"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	// Home Page
	router.HandleFunc("/", handlers.HomePage)

	// Star System Routes
	h := handlers.NewStarSystemHandler()
	router.HandleFunc("/star/{id}", h.GetStarSystem).Methods("GET")
	router.HandleFunc("/star", h.CreateStarSystem).Methods("POST")
	router.HandleFunc("/star/{id}", h.UpdateStarSystem).Methods("PUT")
	router.HandleFunc("/star/{id}", h.DeleteStarSystem).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":10000", router))
}
