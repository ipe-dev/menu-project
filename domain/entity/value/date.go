package value

import "time"

type Date string

func NewDate(timestamp int64) Date {
	d := time.Unix(timestamp, 0).Format("2006/01/02")
	return Date(d)
}
