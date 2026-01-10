package main

import (
	"math/rand"
	"strconv"
)

type Cart struct {
	Id         string
	CustomerId string
	itemList   []*Item
	Amount     int
}

type CartService struct {
	carts map[string]*Cart // userId -> cart
}

func NewCartService() *CartService {
	return &CartService{
		carts: make(map[string]*Cart),
	}
}

func (c *CartService) AddItem(customerId string, item *Item) {
	userCart, ok := c.carts[customerId]

	if !ok {
		userCart = &Cart{
			Id:         "cart" + strconv.Itoa(rand.Intn(1000)),
			CustomerId: customerId,
			itemList:   make([]*Item, 0),
			Amount:     0,
		}
		c.carts[customerId] = userCart
	}

	userCart.itemList = append(userCart.itemList, item)
	userCart.Amount += item.price
}

func (c *CartService) get(customerId string) *Cart {
	userCart, ok := c.carts[customerId]

	if !ok {
		userCart = &Cart{
			Id:         "cart" + strconv.Itoa(rand.Intn(1000)),
			CustomerId: customerId,
			itemList:   make([]*Item, 0),
		}
		c.carts[customerId] = userCart
	}
	return userCart
}
