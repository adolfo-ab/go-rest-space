package planet

import (
	"github.com/google/uuid"
	"main/internal/domain/value_object/planet_type"
)

type Planet struct {
	ID               uuid.UUID              `json:"Id"`
	Name             string                 `json:"Name"`
	PlanetType       planet_type.PlanetType `json:"PlanetType"`
	Mass             float64                `json:"Mass"`
	Radius           float64                `json:"Radius"`
	OrbitalPeriod    float64                `json:"OrbitalPeriod"`
	DistanceFromStar float64                `json:"DistanceFromStar"`
}
