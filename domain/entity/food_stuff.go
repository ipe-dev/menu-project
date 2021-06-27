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

type FoodStuffOption func(*FoodStuff)

func FoodStuffIDOption(ID int) FoodStuffOption {
	return func(f *FoodStuff) {
		if ID != 0 {
			f.ID = ID
		}
	}
}
func FoodStuffMenuIDOption(MenuID int) FoodStuffOption {
	return func(f *FoodStuff) {
		if MenuID != 0 {
			f.MenuID = MenuID
		}
	}
}
func FoodStuffMenuNameOption(Name string) FoodStuffOption {
	return func(f *FoodStuff) {
		if Name != "" {
			f.Name = Name
		}
	}
}

func NewFoodStuff(opts ...FoodStuffOption) *FoodStuff {
	foodStuff := new(FoodStuff)
	for _, opt := range opts {
		opt(foodStuff)
	}
	return foodStuff
}
