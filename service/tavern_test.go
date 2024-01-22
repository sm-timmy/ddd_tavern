package service

import (
	"ddd_tavern/aggregate"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
		WithMemoryBillingRepository(),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

	bills, err := os.bills.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(bills) != 1 {
		t.Error(fmt.Errorf("err bill count %w", errors.New("bill count error")))
	}
}
