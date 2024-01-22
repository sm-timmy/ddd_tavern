package aggregate_test

import (
	"ddd_tavern/aggregate"
	"testing"
)

func TestCustomer_NewCustomer(t *testing.T) {

	type testCase struct {
		name   string
		in     string
		expErr error
	}

	cases := []*testCase{
		{
			name:   "empty name test",
			in:     "",
			expErr: aggregate.ErrInvalidPerson,
		},
		{
			name:   "valid name test",
			in:     "Valid Name",
			expErr: nil,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			// Create a new customer
			_, err := aggregate.NewCustomer(testCase.in)
			// Check if the error matches the expected error
			if err != testCase.expErr {
				t.Errorf("Expected error %v, got %v", testCase.expErr, err)
			}
		})
	}
}
