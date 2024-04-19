package services

import (
	"testing"
)

func TestStarSystemService_CreateStarSystem(t *testing.T) {
	s, err := NewStarSystemService(
		WithMemoryStarSystemRepository(),
	)

	if err != nil {
		t.Error(err)
	}

	names := []string{"s1", "s2", "s3"}

	for _, name := range names {
		err = s.CreateStarSystem(name)
		if err != nil {
			t.Error(err)
		}
	}

	starSystems, err := s.starSystems.GetAll()
	if err != nil {
		t.Error(err)
	}

	if len(starSystems) != len(names) {
		t.Errorf("Incorrect number of star systems in repository. Expected %v, got %v", len(names), len(starSystems))
	}

}
