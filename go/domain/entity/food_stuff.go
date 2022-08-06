package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/errors"
)

type FoodStuff struct {
	ID        int
	Name      string
	MemoID    int
	BuyStatus int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const FoodStuffBuyTrue = 1
const FoodStuffBuyFalse = 0

type FoodStuffOption func(*FoodStuff) error

func FoodStuffIDOption(ID int) FoodStuffOption {
	return func(f *FoodStuff) error {
		if ID != 0 {
			f.ID = ID
		}
		return errors.NewCustomError("FoodStuff作成エラー：IDがありません", ID)
	}
}
func FoodStuffMenuNameOption(Name string) FoodStuffOption {
	return func(f *FoodStuff) error {
		if Name != "" {
			f.Name = Name
		}
		return errors.NewCustomError("FoodStuff作成エラー：食材名がありません", Name)
	}
}

func NewFoodStuff(opts ...FoodStuffOption) *FoodStuff {
	foodStuff := new(FoodStuff)
	for _, opt := range opts {
		opt(foodStuff)
	}
	return foodStuff
}
