// Product is an aggregate that represents a product.
package aggregate

import (
	"ddd_tavern/entity"
	"github.com/google/uuid"
)

// Bill is an aggregate that combines item with a price and quantity
type Bill struct {
	// item is the root entity which is an item
	item  *entity.Item
	price float64
	// customer is the person to pay the bill
	customer uuid.UUID
}

func NewBill(price float64, customerId uuid.UUID) (Bill, error) {
	if price == 0 {
		return Bill{}, ErrMissingValues
	}

	return Bill{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        "Bill",
			Description: "Bill for customer to pay",
		},
		price:    price,
		customer: customerId,
	}, nil
}

func (b Bill) GetID() uuid.UUID {
	return b.item.ID
}

func (b Bill) GetItem() *entity.Item {
	return b.item
}

func (b Bill) GetPrice() float64 {
	return b.price
}
