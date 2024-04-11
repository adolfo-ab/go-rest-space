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
	router.HandleFunc("/planet/{id}", interfaces.ReturnAllPlanets).Methods("GET")
	router.HandleFunc("/planet/{id}", interfaces.DeletePlanet).Methods("DELETE")
	router.HandleFunc("/planet/{id}", interfaces.UpdatePlanet).Methods("PUT")

	log.Fatalln(http.ListenAndServe(":10000", router))
}

func main() {
	Planets = []planet.Planet{
		planet.Planet{ID: "1", Name: "Mercury", PlanetType: planet.Terrestrial, Mass: 0.330, Radius: 2439.7, OrbitalPeriod: 87.969, DistanceFromStar: 57.9},
		planet.Planet{ID: "2", Name: "Venus", PlanetType: planet.Terrestrial, Mass: 4.87, Radius: 6051.8, OrbitalPeriod: 224.701, DistanceFromStar: 108.2},
		planet.Planet{ID: "3", Name: "Earth", PlanetType: planet.Terrestrial, Mass: 5.97, Radius: 6371.0, OrbitalPeriod: 365.256, DistanceFromStar: 149.6},
		planet.Planet{ID: "4", Name: "Mars", PlanetType: planet.Terrestrial, Mass: 0.642, Radius: 3389.5, OrbitalPeriod: 686.971, DistanceFromStar: 227.9},
		planet.Planet{ID: "5", Name: "Jupiter", PlanetType: planet.GasGiant, Mass: 1898, Radius: 69911, OrbitalPeriod: 4332.59, DistanceFromStar: 778.6},
		planet.Planet{ID: "6", Name: "Saturn", PlanetType: planet.GasGiant, Mass: 568, Radius: 58232, OrbitalPeriod: 10759.22, DistanceFromStar: 1433.5},
		planet.Planet{ID: "7", Name: "Uranus", PlanetType: planet.IceGiant, Mass: 86.8, Radius: 25362, OrbitalPeriod: 30688.5, DistanceFromStar: 2872.5},
		planet.Planet{ID: "8", Name: "Neptune", PlanetType: planet.IceGiant, Mass: 102, Radius: 24622, OrbitalPeriod: 60182, DistanceFromStar: 4495.1},
	}

	handleRequests()
}
