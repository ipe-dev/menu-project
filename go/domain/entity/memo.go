package entity

import (
	"time"

	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/errors"
)

type Memo struct {
	ID        int
	UserID    int
	Title     string
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MemoOption func(*Memo) error

func MemoIDOption(ID int) MemoOption {
	return func(m *Memo) error {
		if ID != 0 {
			m.ID = ID
			return nil
		}
		return errors.NewCustomError("Memo作成エラー：IDがありません")
	}
}
func MemoUserIDOption(UserID int) MemoOption {
	return func(m *Memo) error {
		if UserID != 0 {
			m.ID = UserID
			return nil
		}
		return errors.NewCustomError("Memo作成エラー：ユーザーIDがありません")
	}
}
func MemoTitleOption(Title string) MemoOption {
	return func(m *Memo) error {
		if Title != "" {
			m.Title = Title
		}

		return errors.NewCustomError("Memo作成エラー：メモタイトルがありません")
	}
}
func MemoStartDateOption(Timestamp int64) MemoOption {
	return func(m *Memo) error {
		if Timestamp != 0 {
			m.StartDate = value.NewDate(Timestamp)
			return nil
		}
		return errors.NewCustomError("Memo作成エラー：開始日がありません")
	}
}
func MemoEndDateOption(Timestamp int64) MemoOption {
	return func(m *Memo) error {
		if Timestamp != 0 {
			m.EndDate = value.NewDate(Timestamp)
			return nil
		}
		return errors.NewCustomError("Memo作成エラー：終了日がありません")
	}
}
func NewMemo(opts ...MemoOption) (*Memo, error) {
	memo := new(Memo)
	var err error
	for _, opt := range opts {
		if err = opt(memo); err != nil {
			return memo, err
		}
	}
	return memo, nil
}
