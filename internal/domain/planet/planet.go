package planet

import "github.com/google/uuid"

type Planet struct {
	ID               uuid.UUID  `json:"Id"`
	Name             string     `json:"Name"`
	PlanetType       PlanetType `json:"PlanetType"`
	Mass             float64    `json:"Mass"`
	Radius           float64    `json:"Radius"`
	OrbitalPeriod    float64    `json:"OrbitalPeriod"`
	DistanceFromStar float64    `json:"DistanceFromStar"`
}

func NewPlanet(name string, planetType PlanetType,
	mass float64, radius float64, orbitalPeriod float64, distanceFromStar float64) *Planet {
	return &Planet{
		ID:               uuid.New(),
		Name:             name,
		PlanetType:       planetType,
		Mass:             mass,
		Radius:           radius,
		OrbitalPeriod:    orbitalPeriod,
		DistanceFromStar: distanceFromStar,
	}
}
