package services

import (
	"main/internal/domain/aggregates/star_system"
	"main/internal/domain/repositories/star_system_repository"
)

// TODO: Add config

type StarSystemService struct {
	starSystems star_system_repository.StarSystemRepository
}

func NewStarSystemService() *StarSystemService {
	s := &StarSystemService{}

	return s
}

func (s *StarSystemService) CreateStarSystem(name string) error {
	starSystem, err := star_system.NewStarSystem(name)
	if err != nil {
		return err
	}

	err = s.starSystems.Add(starSystem)
	return err
}
