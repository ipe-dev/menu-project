package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/errors"
)

type SubMenu struct {
	ID        int
	Name      string
	MemoID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
type SubMenuOption func(*SubMenu) error

func SubMenuIDOption(ID int) SubMenuOption {
	return func(m *SubMenu) error {
		if ID != 0 {
			m.ID = ID
		}
		return errors.NewCustomError("SubMenu作成エラー：IDがありません")
	}
}
func SubMenuNameOption(Name string) SubMenuOption {
	return func(m *SubMenu) error {
		if Name != "" {
			m.Name = Name
		}
		return errors.NewCustomError("SubMenu作成エラー：サブメニュー名がありません")
	}
}
func SubMenuMemoIDOption(MemoID int) SubMenuOption {
	return func(m *SubMenu) error {
		if MemoID != 0 {
			m.MemoID = MemoID
		}
		return errors.NewCustomError("SubMenu作成エラー：メモIDがありません")
	}
}
func NewSubMenu(opts ...SubMenuOption) *SubMenu {
	submenu := new(SubMenu)
	for _, opt := range opts {
		opt(submenu)
	}
	return submenu
}
