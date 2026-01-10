package main

type Customer struct {
	Id      string
	address string
}

type CustomerService struct {
	customers map[string]*Customer
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		customers: make(map[string]*Customer),
	}
}

func (c *CustomerService) Get(id string) *Customer {
	return c.customers[id]
}
