package entity

import "time"

type SubMenu struct {
	ID        int
	Name      string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type SubMenuOption func(*SubMenu)

func SubMenuIDOption(ID int) SubMenuOption {
	return func(m *SubMenu) {
		if ID != 0 {
			m.ID = ID
		}
	}
}
func SubMenuNameOption(Name string) SubMenuOption {
	return func(m *SubMenu) {
		if Name != "" {
			m.Name = Name
		}
	}
}

func NewSubMenu(opts ...SubMenuOption) *SubMenu {
	submenu := new(SubMenu)
	for _, opt := range opts {
		opt(submenu)
	}
	return submenu
}
