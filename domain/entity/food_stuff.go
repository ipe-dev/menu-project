package entity

import "time"

type FoodStuff struct {
	ID        int
	MenuID    int
	Name      string
	BuyStatus int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const FoodStuffBuyTrue = 1
const FoodStuffBuyFalse = 0
