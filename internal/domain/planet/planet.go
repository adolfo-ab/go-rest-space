package planet

type Planet struct {
	ID               string     `json:"Id"`
	Name             string     `json:"Name"`
	PlanetType       PlanetType `json:"PlanetType"`
	Mass             float64    `json:"Mass"`
	Radius           float64    `json:"Radius"`
	OrbitalPeriod    float64    `json:"OrbitalPeriod"`
	DistanceFromStar float64    `json:"DistanceFromStar"`
}

func NewPlanet(id string, name string, planetType PlanetType,
	mass float64, radius float64, orbitalPeriod float64, distanceFromStar float64) *Planet {
	return &Planet{
		ID:               id,
		Name:             name,
		PlanetType:       planetType,
		Mass:             mass,
		Radius:           radius,
		OrbitalPeriod:    orbitalPeriod,
		DistanceFromStar: distanceFromStar,
	}
}
