package entity

type Menu struct {
	ID         int    `json:"id"`
	WeekID     int    `json:"week_id"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	Kind       int    `json:"kind"`
	URL        string `json:"url"`
	UserID     int    `json:"user_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	FoodStuffs []*FoodStuff
	SubMenu    []*SubMenu
}
