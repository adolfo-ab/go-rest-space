package memory

import (
	"github.com/google/uuid"
	"main/internal/domain/aggregates/star_system"
	"main/internal/domain/repositories/star_system_repository"
	"testing"
)

const fakeId = "f47ac10b-58cc-0372-8567-0e02b2c3d479"

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
			name:        "No StarSystem by ID",
			id:          uuid.MustParse(fakeId),
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

func TestMemory_AddStarSystems(t *testing.T) {
	type testCase struct {
		name        string
		starSystems []string
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Add valid StarSystems",
			starSystems: []string{"S1", "S2"},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repo := MemoryRepository{
				starSystems: map[uuid.UUID]star_system.StarSystem{},
			}

			for _, s := range tc.starSystems {
				starSystem, err := star_system.NewStarSystem(s)
				if err != nil {
					t.Fatal(err)
				}

				err = repo.Add(starSystem)
				if err != tc.expectedErr {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}

				found, err := repo.Get(starSystem.GetID())
				if err != nil {
					t.Fatal(err)
				}

				if found.GetID() != starSystem.GetID() {
					t.Errorf("Expected %v, got %v", starSystem.GetID(), found.GetID())
				}
			}

			if len(tc.starSystems) != len(repo.starSystems) {
				t.Errorf("Repo doesn't added all the star systems. Expected %v, got %v", len(tc.starSystems), len(repo.starSystems))
			}
		})
	}
}

func TestMemory_UpddateStarSystem(t *testing.T) {
	type testCase struct {
		name         string
		id           uuid.UUID
		expectedMass float64
		expectedErr  error
	}

	starSystem, err := star_system.NewStarSystem("Arcturus")
	if err != nil {
		t.Fatal(err)
	}
	id := starSystem.GetID()

	repo := MemoryRepository{
		starSystems: map[uuid.UUID]star_system.StarSystem{
			id: starSystem,
		},
	}

	testCases := []testCase{
		{
			name:         "Update valid star system",
			id:           id,
			expectedMass: 1.5,
			expectedErr:  nil,
		},
		{
			name:         "Try to update invalid star system",
			id:           uuid.MustParse(fakeId),
			expectedMass: 0.0,
			expectedErr:  star_system_repository.ErrUpdateStarSystem,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			starSystem.SetMass(tc.expectedMass)
			starSystem.SetID(tc.id)
			err = repo.Update(starSystem)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(id)
			if err != nil {
				t.Fatal(err)
			}
			if found.GetMass() != tc.expectedMass {
				t.Errorf("Expected %v, got %v", tc.expectedMass, found.GetMass())
			}

		})
	}
}

func TestMemory_DeleteStarSystem(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a new star system to add to the repo
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
			name:        "Try to delete invalid StarSystem",
			id:          uuid.MustParse(fakeId),
			expectedErr: star_system_repository.ErrDeleteStarSystem,
		}, {
			name:        "Delte StarSystem by ID",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Delete(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
