package services

import (
	"github.com/google/uuid"
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

func NewStarSystemService(cfg StarSystemConfig) (*StarSystemService, error) {
	s := &StarSystemService{}

	err := cfg(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *StarSystemService) CreateStarSystem(name string) error {
	starSystem, err := star_system.NewStarSystem(name)
	if err != nil {
		return err
	}

	err = s.starSystems.Add(starSystem)
	return err
}

func (s *StarSystemService) GetStarSystem(id uuid.UUID) (*star_system.StarSystem, error) {
	st, err := s.starSystems.Get(id)
	if err != nil {
		return nil, err
	}
	return &st, nil
}

func (s *StarSystemService) UpdateStarSystem(starSystem star_system.StarSystem) error {
	return s.starSystems.Update(starSystem)
}

func (s *StarSystemService) DeleteStarSystem(id uuid.UUID) error {
	return s.starSystems.Delete(id)
}
