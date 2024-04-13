package repositories

import (
	"errors"
	"github.com/google/uuid"
	"main/internal/domain/aggregates/star_system"
)

var (
	ErrStarSystemNotFound = errors.New("The star system could not be found in the repository")
	ErrAddStarSystem      = errors.New("Failed to add the star system to the repository")
	ErrUpdateStarSystem   = errors.New("Failed to update star system in the repository")
	ErrDeleteStarSystem   = errors.New("Failed to delete star system from the repository")
)

type StarSystemRepository interface {
	Get(uuid.UUID) (star_system.StarSystem, error)
	Add(star_system.StarSystem) error
	Update(system star_system.StarSystem) error
	Delete(uuid2 uuid.UUID) error
}
