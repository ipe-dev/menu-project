package entity

import "time"

type FoodStuff struct {
	ID        int
	MenuID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
