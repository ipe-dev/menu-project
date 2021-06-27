package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity/value"
)

type Menu struct {
	ID        int
	WeekID    int
	Name      string
	Date      value.Date
	Kind      int
	URL       string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const MenuKindLunch = 1
const MenuKindDinner = 2

type MenuOption func(*Menu)

func MenuIDOption(ID int) MenuOption {
	return func(m *Menu) {
		if ID != 0 {
			m.ID = ID
		}
	}
}
func MenuWeekIDOption(WeekID int) MenuOption {
	return func(m *Menu) {
		if WeekID != 0 {
			m.WeekID = WeekID
		}
	}
}
func MenuNameOption(Name string) MenuOption {
	return func(m *Menu) {
		if Name != "" {
			m.Name = Name
		}
	}
}
func MenuDateOption(Timestamp int64) MenuOption {
	return func(m *Menu) {
		if Timestamp != 0 {
			m.Date = value.NewDate(Timestamp)
		}
	}
}
func MenuKindOption(Kind int) MenuOption {
	return func(m *Menu) {
		if Kind != 0 {
			m.Kind = Kind
		}
	}
}
func MenuUrlOption(Url string) MenuOption {
	return func(m *Menu) {
		if Url != "" {
			m.URL = Url
		}
	}
}
func MenuUserIDOption(UserID int) MenuOption {
	return func(m *Menu) {
		if UserID != 0 {
			m.UserID = UserID
		}
	}
}

func NewMenu(opts ...MenuOption) *Menu {
	menu := new(Menu)
	for _, opt := range opts {
		opt(menu)
	}
	return menu
}
