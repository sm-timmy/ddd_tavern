package customer

import (
	"ddd_tavern/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(customer aggregate.Customer) error
	Update(customer aggregate.Customer) error
}
