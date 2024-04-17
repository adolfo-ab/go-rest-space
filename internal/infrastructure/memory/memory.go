package memory

import (
	"github.com/google/uuid"
	"main/internal/domain/aggregates/star_system"
	"main/internal/domain/repositories/star_system_repository"
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
	if starSystem, ok := mr.starSystems[id]; ok {
		return starSystem, nil
	}

	return star_system.StarSystem{}, star_system_repository.ErrStarSystemNotFound
}

func (mr *MemoryRepository) GetAll() ([]star_system.StarSystem, error) {
	if len(mr.starSystems) == 0 {
		return nil, star_system_repository.ErrStarSystemRepoEmpty
	}
	var starSystems []star_system.StarSystem
	for _, s := range mr.starSystems {
		starSystems = append(starSystems, s)
	}
	return starSystems, nil
}

func (mr *MemoryRepository) Add(s star_system.StarSystem) error {
	if mr.starSystems == nil {
		mr.Lock()
		mr.starSystems = make(map[uuid.UUID]star_system.StarSystem)
		mr.Unlock()
	}

	if _, ok := mr.starSystems[s.GetID()]; ok {
		return star_system_repository.ErrAddStarSystem
	}

	mr.Lock()
	mr.starSystems[s.GetID()] = s
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(s star_system.StarSystem) error {
	if _, ok := mr.starSystems[s.GetID()]; !ok {
		return star_system_repository.ErrUpdateStarSystem
	}

	mr.Lock()
	mr.starSystems[s.GetID()] = s
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
	if _, ok := mr.starSystems[id]; !ok {
		return star_system_repository.ErrDeleteStarSystem
	}
	mr.Lock()
	delete(mr.starSystems, id)
	mr.Unlock()
	return nil
}
