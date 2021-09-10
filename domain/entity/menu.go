package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/errors"
)

type Menu struct {
	ID        int
	MemoID    int
	Name      string
	Date      time.Time
	Kind      int
	URL       string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

const MenuKindLunch = 1  // 昼ご飯
const MenuKindDinner = 2 // 夜ご飯

type MenuOption func(*Menu) error

func (menu Menu) CheckMenuDate(memo Memo) error {
	if menu.Date.Before(memo.StartDate) || menu.Date.After(memo.EndDate) {
		return errors.NewCustomError("Menu作成エラー：日付が不正です", menu.Date, memo.StartDate, memo.EndDate)
	}
	return nil
}
func MenuIDOption(ID int) MenuOption {
	return func(m *Menu) error {
		if ID != 0 {
			m.ID = ID
		}
		return errors.NewCustomError("Menu作成エラー：IDがありません")
	}
}
func MenuMemoIDOption(MemoID int) MenuOption {
	return func(m *Menu) error {
		if MemoID != 0 {
			m.MemoID = MemoID
		}
		return errors.NewCustomError("Menu作成エラー：メモIDがありません")
	}
}
func MenuNameOption(Name string) MenuOption {
	return func(m *Menu) error {
		if Name != "" {
			m.Name = Name
		}
		return errors.NewCustomError("Menu作成エラー：メニュー名がありません")

	}
}
func MenuDateOption(Timestamp int64) MenuOption {
	return func(m *Menu) error {
		if Timestamp != 0 {
			m.Date = value.NewDate(Timestamp)
		}
		return errors.NewCustomError("Menu作成エラー：IDがありません")
	}
}
func MenuKindOption(Kind int) MenuOption {
	return func(m *Menu) error {
		if Kind != 0 {
			m.Kind = Kind
		}
		return errors.NewCustomError("Menu作成エラー：昼夜種別がありません")
	}
}
func MenuUrlOption(Url string) MenuOption {
	return func(m *Menu) error {
		if Url != "" {
			m.URL = Url
		}
		return nil
	}
}

func NewMenu(opts ...MenuOption) *Menu {
	menu := new(Menu)
	for _, opt := range opts {
		opt(menu)
	}
	return menu
}
