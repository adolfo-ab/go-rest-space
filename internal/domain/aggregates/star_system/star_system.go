package star_system

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain/entities/planet"
	"main/internal/domain/entities/star"
)

var (
	ErrInvalidStar = errors.New("A star system needs a valid star")
)

type StarSystem struct {
	star    *star.Star
	planets []*planet.Planet
}

func NewStarSystem(starName string) (StarSystem, error) {
	if starName == "" {
		return StarSystem{}, ErrInvalidStar
	}

	star := &star.Star{
		ID:   uuid.New(),
		Name: starName,
	}

	return StarSystem{
		star:    star,
		planets: make([]*planet.Planet, 0),
	}, nil
}
