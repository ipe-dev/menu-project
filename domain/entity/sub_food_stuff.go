package entity

import "time"

type SubFoodStuff struct {
	ID        int
	SubMenuID int
	Name      string
	BuyStatus int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const BuySubFoodStatusTrue = 1
const BuySubFoodStatusFalse = 0
