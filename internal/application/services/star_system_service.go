package services

import "main/internal/domain/repositories/star_system_repository"

// TODO: Add config

type StarSystemService struct {
	starSystem star_system_repository.StarSystemRepository
}

func NewStarSystemService() *StarSystemService {
	s := &StarSystemService{}

	return s
}
