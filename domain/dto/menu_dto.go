package dto

import "time"

type Menu struct {
	ID       int       `json:"id"`
	MemoID   int       `json:"memo_id"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	UnixDate int64     `json:"unix_date"`
	Kind     int       `json:"kind"`
	URL      string    `json:"url"`
}
