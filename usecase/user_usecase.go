package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type UserUseCase interface {
	Get(GetUserRequest) entity.User
	Create(CreateUserRequest) error
	Update(UpdateUserRequest) error
	Login(LoginRequest) error
	Logout(LogoutRequest) error
}
type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) userUseCase {
	return &userUseCase{r}
}

type GetUserRequest struct {
	ID int `json:"id"`
}
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type UpdateUserRequest struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginRequest struct {
	LoginID  string `json:"login_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LogoutRequest struct {
	ID int `json"id" validate:"required"`
}

func (u userUseCase) Get(r GetUserRequest) entity.User {
	user := u.userRepository.Get(r.ID)
	return user
}

func (u userUseCase) Create(r CreateUserRequest) error {
	user := entity.User{
		Name:     r.Name,
		LoginID:  r.LoginID,
		Password: r.Password,
	}

	var err error
	err = u.userRepository.Create(user)
	return err
}
func (u userUseCase) Update(r UpdateUserRequest) error {
	user := entity.User{
		ID:       r.ID,
		Name:     r.Name,
		LoginID:  r.LoginID,
		Password: r.Password,
	}

	var err error
	err = u.userRepository.Update(user)
	return err
}
