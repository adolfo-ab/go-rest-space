package main

import (
	"log"
	"net/http"
)

type Planet struct {
	id               string
	name             string
	planetType       PlanetType
	mass             float64
	radius           float64
	orbitalPeriod    float64
	distanceFromStar float64
}

type PlanetType int

const (
	Undefined PlanetType = iota
	Terrestrial
	GasGiant
	IceGiant
	DwarfPlanet
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello world"}`))
}

func main() {
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
