package dto

import (
	"time"
)

type Memo struct {
	ID         int         `json:"id"`
	UserID     int         `json:"user_id"`
	Title      string      `json:"title"`
	StartDate  time.Time   `json:"start_date"`
	EndDate    time.Time   `json:"end_date"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Menus      []Menu      `json:"menus"`
	SubMenus   []SubMenu   `json:"sub_menus"`
	FoodStuffs []FoodStuff `json:"food_stuffs"`
}
