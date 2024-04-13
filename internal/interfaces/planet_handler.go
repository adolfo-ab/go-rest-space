package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"main/internal/domain/entity/planet"
	"main/internal/domain/value_object/planet_type"
	"net/http"
)

var Planets = []*planet.Planet{
	planet.NewPlanet("Mercury", planet_type.Terrestrial, 0.330, 2439.7, 87.969, 57.9),
	planet.NewPlanet("Venus", planet_type.Terrestrial, 4.87, 6051.8, 224.701, 108.2),
	planet.NewPlanet("Earth", planet_type.Terrestrial, 5.97, 6371.0, 365.256, 149.6),
	planet.NewPlanet("Mars", planet_type.Terrestrial, 0.642, 3389.5, 686.971, 227.9),
	planet.NewPlanet("Jupiter", planet_type.GasGiant, 1898, 69911, 4332.59, 778.6),
	planet.NewPlanet("Saturn", planet_type.GasGiant, 568, 58232, 10759.22, 1433.5),
	planet.NewPlanet("Uranus", planet_type.IceGiant, 86.8, 25362, 30688.5, 2872.5),
	planet.NewPlanet("Neptune", planet_type.IceGiant, 102, 24622, 60182, 4495.1),
}

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
		if planet.ID.String() == id {
			json.NewEncoder(w).Encode(planet)
		}
	}
}

func CreateNewPlanet(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var planet planet.Planet

	json.Unmarshal(reqBody, &planet)

	Planets = append(Planets, &planet)
	json.NewEncoder(w).Encode(planet)
}

func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, planet := range Planets {
		if planet.ID.String() == id {
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

	if newPlanet.ID.String() != id {
		return
	}

	for index, planet := range Planets {
		if planet.ID.String() == id {
			Planets[index] = &newPlanet
		}
	}
}
