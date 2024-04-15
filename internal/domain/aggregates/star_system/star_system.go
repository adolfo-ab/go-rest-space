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

// StarSystem is an aggregate that includes a Star (root entity) and Planets orbiting around it
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

// Returns the ID of the star (root entity) of the star system aggregate
func (s *StarSystem) GetID() uuid.UUID {
	return s.star.ID
}

// Set the ID of the star (root entity) of the star system aggregate
func (s *StarSystem) SetID(id uuid.UUID) {
	if s.star == nil {
		s.star = &star.Star{}
	}
	s.star.ID = id
}

// Returns the Name of the star (root entity) of the star system aggregate
func (s *StarSystem) GetName() string {
	return s.star.Name
}

// Set the Name of the star (root entity) of the star system aggregate
func (s *StarSystem) SetName(name string) {
	if s.star == nil {
		s.star = &star.Star{}
	}
	s.star.Name = name
}

// Returns the Mass of the star (root entity) of the star system aggregate
func (s *StarSystem) GetMass() float64 {
	return s.star.Mass
}

// Set the Mass of the star (root entity) of the star system aggregate
func (s *StarSystem) SetMass(mass float64) {
	if s.star == nil {
		s.star = &star.Star{}
	}
	s.star.Mass = mass
}
