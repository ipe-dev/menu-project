package value

import "time"

func NewDate(timestamp int64) time.Time {
	d := time.Unix(timestamp, 0)
	return d
}
