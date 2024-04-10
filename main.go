package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Planet struct {
	ID               string     `json:"Id"`
	Name             string     `json:"Name"`
	PlanetType       PlanetType `json:"PlanetType"`
	Mass             float64    `json:"Mass"`
	Radius           float64    `json:"Radius"`
	OrbitalPeriod    float64    `json:"OrbitalPeriod"`
	DistanceFromStar float64    `json:"DistanceFromStar"`
}

type PlanetType int

const (
	Undefined PlanetType = iota
	Terrestrial
	GasGiant
	IceGiant
	DwarfPlanet
)

var Planets []Planet

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/planets", returnAllPlanets)
	router.HandleFunc("/planet", createNewPlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", returnSinglePlanet).Methods("GET")
	router.HandleFunc("/planet/{id}", deletePlanet).Methods("DELETE")

	log.Fatalln(http.ListenAndServe(":10000", router))
}

func main() {
	Planets = []Planet{
		Planet{ID: "1", Name: "Mercury", PlanetType: Terrestrial, Mass: 0.330, Radius: 2439.7, OrbitalPeriod: 87.969, DistanceFromStar: 57.9},
		Planet{ID: "2", Name: "Venus", PlanetType: Terrestrial, Mass: 4.87, Radius: 6051.8, OrbitalPeriod: 224.701, DistanceFromStar: 108.2},
		Planet{ID: "3", Name: "Earth", PlanetType: Terrestrial, Mass: 5.97, Radius: 6371.0, OrbitalPeriod: 365.256, DistanceFromStar: 149.6},
		Planet{ID: "4", Name: "Mars", PlanetType: Terrestrial, Mass: 0.642, Radius: 3389.5, OrbitalPeriod: 686.971, DistanceFromStar: 227.9},
		Planet{ID: "5", Name: "Jupiter", PlanetType: GasGiant, Mass: 1898, Radius: 69911, OrbitalPeriod: 4332.59, DistanceFromStar: 778.6},
		Planet{ID: "6", Name: "Saturn", PlanetType: GasGiant, Mass: 568, Radius: 58232, OrbitalPeriod: 10759.22, DistanceFromStar: 1433.5},
		Planet{ID: "7", Name: "Uranus", PlanetType: IceGiant, Mass: 86.8, Radius: 25362, OrbitalPeriod: 30688.5, DistanceFromStar: 2872.5},
		Planet{ID: "8", Name: "Neptune", PlanetType: IceGiant, Mass: 102, Radius: 24622, OrbitalPeriod: 60182, DistanceFromStar: 4495.1},
	}

	handleRequests()
}

func returnAllPlanets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPlanets")
	json.NewEncoder(w).Encode(Planets)
}

func returnSinglePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, planet := range Planets {
		if planet.ID == key {
			json.NewEncoder(w).Encode(planet)
		}
	}
}

func createNewPlanet(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var planet Planet

	json.Unmarshal(reqBody, &planet)

	Planets = append(Planets, planet)
	json.NewEncoder(w).Encode(planet)
}

func deletePlanet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	for index, planet := range Planets {
		if planet.ID == id {
			Planets = append(Planets[:index], Planets[index+1:]...)
		}
	}
}
