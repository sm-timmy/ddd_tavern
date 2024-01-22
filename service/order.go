package service

import (
	"ddd_tavern/aggregate"
	"ddd_tavern/domain/bill"
	billingMemory "ddd_tavern/domain/bill/memory"
	"ddd_tavern/domain/customer"
	customerMemory "ddd_tavern/domain/customer/memory"
	"ddd_tavern/domain/product"
	productMemory "ddd_tavern/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

// OrderConfiguration is an alias for a function that will take in a pointer to an OrderService and modify it
type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
	bills     bill.BillingRepository
}

// NewOrderService takes a variable amount of OrderConfiguration functions and returns a new OrderService
// Each OrderConfiguration will be called in the order they are passed in
func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a given customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

// WithMemoryProductRepository adds in memory product repo and adds all input products
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
		pr := productMemory.New()

		// Add Items to repo
		for _, p := range products {
			err := pr.Add(p)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

// WithMemoryCustomerRepository applies a memory customer repository to the OrderService
func WithMemoryCustomerRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	cr := customerMemory.New()
	return WithCustomerRepository(cr)
}

// WithMemoryBillingRepository applies a memory customer repository to the BillingService
func WithMemoryBillingRepository() OrderConfiguration {
	// Create the memory repo, if we needed parameters, such as connection strings they could be inputted here
	br := billingMemory.New()
	return WithBillingRepository(br)
}

func WithBillingRepository(br bill.BillingRepository) OrderConfiguration {
	// return a function that matches the BillingConfiguration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(os *OrderService) error {
		os.bills = br
		return nil
	}
}

// CreateOrder will chain together all repositories to create an order for a customer
func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Get the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product
	var products []aggregate.Product
	var price float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		price += p.GetPrice()
	}

	// All Products exists in store, now we can create the order
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))

	return price, nil
}

// // CreateBilling will chain together all repositories to create an Billing for a customer
func (os *OrderService) Bill(customerID uuid.UUID, price float64) error {
	// Get the customer
	cust, err := os.customers.Get(customerID)
	if err != nil {
		return err
	}

	newBill, err := aggregate.NewBill(price, cust.GetID())
	if err != nil {
		return err
	}

	err = os.bills.Add(newBill)
	if err != nil {
		return err
	}

	// All Products exists in store, now we can create the Billing
	log.Printf("Customer: %s has been Billed", cust.GetID())
	return nil
}
