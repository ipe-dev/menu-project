package entity

import "time"

type SubMenu struct {
	ID        int
	Name      string
	MenuID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
