package repository

type UserRepository interface {
	Create() error
	Update() error
}
