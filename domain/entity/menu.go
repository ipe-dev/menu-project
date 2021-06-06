package entity

import (
	"time"
)

type Menu struct {
	ID        int
	WeekID    int
	Name      string
	Date      string
	Kind      int
	URL       string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const MenuKindLunch = 1
const MenuKindDinner = 2
