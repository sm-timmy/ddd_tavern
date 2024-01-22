package aggregate_test

import (
	"ddd_tavern/aggregate"
	"github.com/google/uuid"
	"testing"
)

func TestBill_NewBill(t *testing.T) {
	type testCase struct {
		testName    string
		price       float64
		customerId  uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			testName:    "should return error if price is zero",
			price:       0,
			expectedErr: aggregate.ErrMissingValues,
		},
		{
			testName:    "valid values",
			price:       1.0,
			customerId:  uuid.New(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, err := aggregate.NewBill(tc.price, tc.customerId)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
