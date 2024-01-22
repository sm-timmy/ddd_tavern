package bill

import (
	"ddd_tavern/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrBillNotFound    = errors.New("the Bill was not found in the repository")
	ErrFailedToAddBill = errors.New("failed to add the Bill to the repository")
	ErrUpdateBill      = errors.New("failed to update the Bill in the repository")
)

type BillingRepository interface {
	GetAll() ([]aggregate.Bill, error)
	Get(uuid uuid.UUID) (aggregate.Bill, error)
	Add(bill aggregate.Bill) error
	Update(bill aggregate.Bill) error
}
