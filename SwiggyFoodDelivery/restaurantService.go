package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Item struct {
	Id          string
	Name        string
	IsAvailable bool
	price       int
}

type Restaurant struct {
	id        string
	menuItems []*Item
}

type RestaurantService struct {
	restaurantList []*Restaurant
}

func NewRestaurantService() *RestaurantService {
	return &RestaurantService{
		restaurantList: make([]*Restaurant, 0),
	}
}

func (s *RestaurantService) CreateResaturant() *Restaurant {
	newRestaurant := &Restaurant{
		id:        "rest" + strconv.Itoa(rand.Intn(1000)),
		menuItems: make([]*Item, 0),
	}

	return newRestaurant
}

func (s *RestaurantService) AddItems(restaurantId string) {
	for _, r := range s.restaurantList {
		if r.id == restaurantId {
			newItem := &Item{
				Id:          "item" + strconv.Itoa(rand.Intn(1000)),
				Name:        "chicken" + strconv.Itoa(rand.Intn(1000)),
				IsAvailable: true,
				price:       rand.Intn(1000),
			}
			r.menuItems = append(r.menuItems, newItem)
			fmt.Println("added a new item : ", newItem, " to the restaurant : ", r)
		}
	}
}

func (s *RestaurantService) getAll() []*Restaurant {
	return s.restaurantList
}

func (s *RestaurantService) getMenu(restaurantId string) []*Item {
	for _, r := range s.restaurantList {
		if r.id == restaurantId {
			return r.menuItems
		}
	}
	return nil
}
