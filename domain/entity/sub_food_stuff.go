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

type SubFoodStuffOption func(*SubFoodStuff)

func SubFoodStuffIDOption(ID int) SubFoodStuffOption {
	return func(s *SubFoodStuff) {
		if ID != 0 {
			s.ID = ID
		}
	}
}
func SubFoodStuffMenuIDOption(SubMenuID int) SubFoodStuffOption {
	return func(s *SubFoodStuff) {
		if SubMenuID != 0 {
			s.SubMenuID = SubMenuID
		}
	}
}
func SubFoodStuffNameOption(Name string) SubFoodStuffOption {
	return func(s *SubFoodStuff) {
		if Name != "" {
			s.Name = Name
		}
	}
}
func NewSubFoodStuff(opts ...SubFoodStuffOption) *SubFoodStuff {
	subfoodstuff := new(SubFoodStuff)
	for _, opt := range opts {
		opt(subfoodstuff)
	}
	return subfoodstuff
}
