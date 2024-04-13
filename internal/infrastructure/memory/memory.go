package memory

import (
	"github.com/google/uuid"
	"main/internal/domain/aggregates/star_system"
	"sync"
)

type MemoryRepository struct {
	starSystems map[uuid.UUID]star_system.StarSystem
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		starSystems: make(map[uuid.UUID]star_system.StarSystem),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (star_system.StarSystem, error) {
	return star_system.StarSystem{}, nil
}

func (mr *MemoryRepository) Add(starSystem star_system.StarSystem) error {
	return nil
}

func (mr *MemoryRepository) Update(starSystem star_system.StarSystem) error {
	return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	return nil
}
