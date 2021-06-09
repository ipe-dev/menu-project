package entity

import "time"

type FoodStuff struct {
	ID        int
	MenuID    int
	Name      string
	CreateAt  time.Time
	UpdatedAt time.Time
}
