package star_system

import "testing"

func TestStarSystem_NewStarSystem(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty Star Name",
			name:        "",
			expectedErr: ErrInvalidStar,
		},
		{
			test:        "Valid Star Name",
			name:        "Algedi",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := NewStarSystem(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
