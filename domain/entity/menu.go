package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity/value"
)

type Menu struct {
	ID        int
	MemoID    int
	Name      string
	Date      string
	Kind      int
	URL       string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const MenuKindLunch = 1  // 昼ご飯
const MenuKindDinner = 2 // 夜ご飯

type MenuOption func(*Menu)

func MenuIDOption(ID int) MenuOption {
	return func(m *Menu) {
		if ID != 0 {
			m.ID = ID
		}
	}
}
func MenuMemoIDOption(MemoID int) MenuOption {
	return func(m *Menu) {
		if MemoID != 0 {
			m.MemoID = MemoID
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
