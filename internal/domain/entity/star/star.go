package star

import (
	"github.com/google/uuid"
	"main/internal/domain/value_object/star_type"
)

type Star struct {
	ID          uuid.UUID          `json:"Id"`
	Name        string             `json:"Name"`
	StarType    star_type.StarType `json:"StarType"`
	Mass        float64            `json:"Mass"`
	Radius      float64            `json:"Radius"`
	Temperature float64            `json:"Temperature"`
	Luminosity  float64            `json:"Luminosity"`
	Metallicity float64            `json:"Metallicity"`
}
