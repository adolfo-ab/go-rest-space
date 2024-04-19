package services

import (
	"main/internal/domain/aggregates/star_system"
	"main/internal/domain/repositories/star_system_repository"
	"main/internal/infrastructure/memory"
)

type StarSystemConfig func(ss *StarSystemService) error

func WithStarSystemRepository(sr star_system_repository.StarSystemRepository) StarSystemConfig {
	return func(s *StarSystemService) error {
		s.starSystems = sr
		return nil
	}
}

func WithMemoryStarSystemRepository() StarSystemConfig {
	sr := memory.New()
	return WithStarSystemRepository(sr)
}

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
