package star_system_repository

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain/aggregates/star_system"
)

var (
	ErrStarSystemNotFound  = errors.New("the star system could not be found in the repository")
	ErrStarSystemRepoEmpty = errors.New("empty repository: no star systems found")
	ErrAddStarSystem       = errors.New("failed to add the star system to the repository")
	ErrUpdateStarSystem    = errors.New("failed to update star system in the repository")
	ErrDeleteStarSystem    = errors.New("failed to delete star system from the repository")
)

type StarSystemRepository interface {
	Get(uuid.UUID) (star_system.StarSystem, error)
	GetAll() ([]star_system.StarSystem, error)
	Add(star_system.StarSystem) error
	Update(star_system.StarSystem) error
	Delete(uuid.UUID) error
}
