package memory

import (
	"github.com/google/uuid"
	"main/internal/domain/aggregates/star_system"
	"main/internal/domain/repositories/star_system_repository"
	"testing"
)

func TestMemory_GetStarSystem(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a new star system to add to the repository
	starSystem, err := star_system.NewStarSystem("Antares")
	if err != nil {
		t.Fatal(err)
	}
	id := starSystem.GetID()

	// Create repository and add test data
	repo := MemoryRepository{
		starSystems: map[uuid.UUID]star_system.StarSystem{
			id: starSystem,
		},
	}

	testCases := []testCase{
		{
			name:        "No StarSystem By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: star_system_repository.ErrStarSystemNotFound,
		}, {
			name:        "StarSystem By ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			_, err := repo.Get(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

func TestMemory_AddStarSystem(t *testing.T) {
	type testCase struct {
		name        string
		starSystem  string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add Customer",
			starSystem:  "Betelgeuse",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				starSystems: map[uuid.UUID]star_system.StarSystem{},
			}

			cust, err := star_system.NewStarSystem(tc.starSystem)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(cust)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(cust.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != cust.GetID() {
				t.Errorf("Expected %v, got %v", cust.GetID(), found.GetID())
			}
		})
	}
}

// TODO: Implement Update and Delete Tests
