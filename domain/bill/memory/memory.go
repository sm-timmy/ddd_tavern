package memory

import (
	"ddd_tavern/aggregate"
	"ddd_tavern/domain/bill"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	bills map[uuid.UUID]aggregate.Bill
	sync.RWMutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		bills: make(map[uuid.UUID]aggregate.Bill, 0),
	}
}

// Get all bills
func (mr *MemoryRepository) GetAll() ([]aggregate.Bill, error) {
	// Collect all Products from map
	var bills []aggregate.Bill
	for _, bill := range mr.bills {
		bills = append(bills, bill)
	}
	return bills, nil
}

// Get finds a bill by ID
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Bill, error) {
	if b, ok := mr.bills[id]; ok {
		return b, nil
	}
	return aggregate.Bill{}, bill.ErrBillNotFound
}

func (mr *MemoryRepository) Add(c aggregate.Bill) error {
	if mr.bills == nil {
		// Saftey check if bills is not create, shouldn't happen if using the Factory, but you never know
		mr.Lock()
		mr.bills = make(map[uuid.UUID]aggregate.Bill)
		mr.Unlock()
	}

	if _, ok := mr.bills[c.GetID()]; ok {
		return fmt.Errorf("bill already exists %w", bill.ErrFailedToAddBill)
	}

	mr.Lock()
	mr.bills[c.GetID()] = c
	mr.Unlock()
	return nil
}

// Update will replace an existing bill information with the new bill information
func (mr *MemoryRepository) Update(c aggregate.Bill) error {
	// Make sure Bill is in the repository
	if _, ok := mr.bills[c.GetID()]; !ok {
		return fmt.Errorf("bill does not exist: %w", bill.ErrUpdateBill)
	}

	mr.Lock()
	mr.bills[c.GetID()] = c
	mr.Unlock()
	return nil
}
