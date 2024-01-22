package memory

import (
	"ddd_tavern/aggregate"
	billDomain "ddd_tavern/domain/bill"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetBill(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	// Create a fake customer to add to repository
	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()

	// Create a fake bill for this customer to add to repository
	bill, err := aggregate.NewBill(100, id)
	if err != nil {
		t.Fatal(err)
	}

	// Create the repo to use, and add some test Data to it for testing
	// Skip Factory for this
	repo := MemoryRepository{
		bills: map[uuid.UUID]aggregate.Bill{
			bill.GetID(): bill,
		},
	}

	testCases := []testCase{
		{
			name:        "No Bill By ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: billDomain.ErrBillNotFound,
		}, {
			name:        "Bill By ID",
			id:          bill.GetID(),
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

func TestMemory_AddBill(t *testing.T) {
	type testCase struct {
		name        string
		price       float64
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "First Customer",
			price:       100,
			expectedErr: nil,
		},
	}

	repo := MemoryRepository{
		bills: map[uuid.UUID]aggregate.Bill{},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			cust, err := aggregate.NewCustomer(tc.name)
			if err != nil {
				t.Fatal(err)
			}
			custId := cust.GetID()

			// Create a fake bill for this customer to add to repository
			bill, err := aggregate.NewBill(100, custId)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(bill)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			found, err := repo.Get(bill.GetID())
			if err != nil {
				t.Fatal(err)
			}
			if found.GetID() != bill.GetID() {
				t.Errorf("Expected %v, got %v", bill.GetID(), found.GetID())
			}

			if len(repo.bills) != 1 {
				t.Errorf("Expected 1 product, got %d", len(repo.bills))
			}
		})
	}
}
