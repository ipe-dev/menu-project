package entity

import "time"

type User struct {
	ID        int
	Name      string
	LoginID   string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u User) Login() {

}

func (u User) Logout() {

}
