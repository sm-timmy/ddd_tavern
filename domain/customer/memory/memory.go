package memory

import (
	"ddd_tavern/aggregate"
	"ddd_tavern/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.RWMutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregate.Customer, 0),
	}
}

// Get finds a customer by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if c, ok := mr.customers[id]; ok {
		return c, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.customers == nil {
		// Saftey check if customers is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists %w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing customer information with the new customer information
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	// Make sure Customer is in the repository
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
