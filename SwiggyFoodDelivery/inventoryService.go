package main

import "time"

type InventoryItem struct {
	id         string
	itemId     string
	totalQty   int
	reserveQty int
	updatedAt  time.Time
}

type Inventory struct {
	inventoryItemList map[string]*InventoryItem // itemId -> inventoryItem
}

type InventoryService struct {
	inventoryList map[string]*Inventory // restaurantId -> Inventory
}

func (s *InventoryService) reserve(restaurantId string, userItemId string) bool {
	return true
}
