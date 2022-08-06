package value

import "time"

func NewTitle(startDate int64, endDate int64) string {
	title := time.Unix(startDate, 0).Format("2001/01/02") + "~" + time.Unix(endDate, 0).Format("2001/01/02")
	return title
}
