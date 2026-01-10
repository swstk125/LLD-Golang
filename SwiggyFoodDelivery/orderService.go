package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type OrderStatus string

const (
	CREATED         OrderStatus = "CREATED"
	PAYMENT_PENDING OrderStatus = "PAYMENT_PENDING"
	CONFIRMED       OrderStatus = "CONFIRMED"
	PREPARING       OrderStatus = "PREPARING"
	ON_THE_WAY      OrderStatus = "ON_THE_WAY"
	DELIVERED       OrderStatus = "DELIVERED"
	CANCELLED       OrderStatus = "CANCELLED"
)

type Order struct {
	Id         string
	CustomerId string
	amount     int
	userCart   *Cart
	status     OrderStatus
	Agent      *DeliveryAgent
}

type OrderService struct {
	OrdersById map[string]*Order
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrdersById: make(map[string]*Order),
	}
}

func (o *OrderService) Create(customerId string, c *Cart) *Order {
	newOrder := &Order{
		Id:         "order" + strconv.Itoa(rand.Intn(10000)),
		CustomerId: customerId,
		amount:     c.Amount,
		userCart:   c,
		status:     CREATED,
	}
	o.OrdersById[newOrder.Id] = newOrder

	fmt.Println("OrderCreated : ", newOrder)
	return newOrder
}

func (o *OrderService) Place(customerId string, c *Cart) *Order {
	newOrder := o.Create(customerId, c)
	newOrder.status = PAYMENT_PENDING

	fmt.Println("MakePayment to place Order : ", newOrder)
	return newOrder
}

func (o *OrderService) AssignDeliveryAgent(orderId string, a *DeliveryAgent) {
	o.OrdersById[orderId].Agent = a
}
