package entity

import "github.com/ipe-dev/menu_project/domain/entity/value"

type Memo struct {
	ID        int
	UserID    int
	Title     string
	StartDate string
	EndDate   string
}

type MemoOption func(*Memo)

func MemoIDOption(ID int) MemoOption {
	return func(m *Memo) {
		if ID != 0 {
			m.ID = ID
		}
	}
}
func MemoUserIDOption(ID int) MemoOption {
	return func(m *Memo) {
		if ID != 0 {
			m.ID = ID
		}
	}
}
func MemoTitleOption(Title string) MemoOption {
	return func(m *Memo) {
		if Title != "" {
			m.Title = Title
		}
	}
}
func MemoStartDateOption(Timestamp int64) MemoOption {
	return func(m *Memo) {
		if Timestamp != 0 {
			m.StartDate = value.NewDate(Timestamp)
		}
	}
}
func MemoEndDateOption(Timestamp int64) MemoOption {
	return func(m *Memo) {
		if Timestamp != 0 {
			m.EndDate = value.NewDate(Timestamp)
		}
	}
}
func NewMemo(opts ...MemoOption) *Memo {
	memo := new(Memo)
	for _, opt := range opts {
		opt(memo)
	}
	return memo
}
