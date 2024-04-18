package star_system

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain/entities/planet"
	"main/internal/domain/entities/star"
	"main/internal/domain/value_objects/star_type"
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

	s := &star.Star{
		ID:   uuid.New(),
		Name: starName,
	}

	return StarSystem{
		star:    s,
		planets: make([]*planet.Planet, 0),
	}, nil
}

func (s *StarSystem) GetID() uuid.UUID {
	return s.star.ID
}

func (s *StarSystem) SetID(id uuid.UUID) {
	if s.star == nil {
		s.star = &star.Star{}
	}
	s.star.ID = id
}

func (s *StarSystem) GetName() string {
	return s.star.Name
}

func (s *StarSystem) SetName(name string) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.Name = name

	return nil
}

func (s *StarSystem) GetStarType() star_type.StarType {
	return s.star.StarType
}

func (s *StarSystem) SetStarType(st star_type.StarType) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.StarType = st

	return nil
}

func (s *StarSystem) GetMass() float64 {
	return s.star.Mass
}

func (s *StarSystem) SetMass(mass float64) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.Mass = mass

	return nil
}

func (s *StarSystem) GetRadius() float64 {
	return s.star.Radius
}

func (s *StarSystem) SetRadius(r float64) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.Radius = r

	return nil
}

func (s *StarSystem) GetTemperature() float64 {
	return s.star.Temperature
}

func (s *StarSystem) SetTemperature(t float64) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.Temperature = t

	return nil
}

func (s *StarSystem) GetLuminosity() float64 {
	return s.star.Luminosity
}

func (s *StarSystem) SetLuminosity(l float64) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.Mass = l

	return nil
}

func (s *StarSystem) GetMetallicity() float64 {
	return s.star.Metallicity
}

func (s *StarSystem) SetMetallicity(m float64) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.star.Mass = m

	return nil
}

func (s *StarSystem) GetPlanets() []*planet.Planet { return s.planets }

func (s *StarSystem) AddPlanet(p *planet.Planet) error {
	if s.star == nil {
		return ErrInvalidStar
	}
	s.planets = append(s.planets, p)

	return nil
}
