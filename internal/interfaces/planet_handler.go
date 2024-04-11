package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"main/internal/domain/planet"
	"net/http"
)

var Planets []planet.Planet

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func ReturnAllPlanets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Planets)
}

func ReturnSinglePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, planet := range Planets {
		if planet.ID == id {
			json.NewEncoder(w).Encode(planet)
		}
	}
}

func CreateNewPlanet(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var planet planet.Planet

	json.Unmarshal(reqBody, &planet)

	Planets = append(Planets, planet)
	json.NewEncoder(w).Encode(planet)
}

func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, planet := range Planets {
		if planet.ID == id {
			Planets = append(Planets[:index], Planets[index+1:]...)
		}
	}
}

func UpdatePlanet(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var newPlanet planet.Planet
	json.Unmarshal(reqBody, &newPlanet)

	vars := mux.Vars(r)
	id := vars["id"]

	if newPlanet.ID != id {
		return
	}

	for index, planet := range Planets {
		if planet.ID == id {
			Planets[index] = newPlanet
		}
	}
}
