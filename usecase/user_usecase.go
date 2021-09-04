package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/domain/service"
	"github.com/ipe-dev/menu_project/errors"
)

type UserUseCase interface {
	Get(GetUserRequest) (*entity.User, error)
	Create(CreateUserRequest) error
	Update(UpdateUserRequest) error
	LoginAuthenticate(LoginRequest) (*entity.User, error)
}
type userUseCase struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewUserUseCase(r repository.UserRepository, s service.UserService) UserUseCase {
	return &userUseCase{r, s}
}

type GetUserRequest struct {
	ID int `json:"id"`
}
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginRequest struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LogoutRequest struct {
	ID int `json"id" binding:"required"`
}

func (u userUseCase) Get(r GetUserRequest) (*entity.User, error) {
	user, err := u.userRepository.Get(r.ID)
	return user, err
}

func (u userUseCase) Create(r CreateUserRequest) error {

	// ログインID使用済みチェック
	exists, err := u.userService.CheckUserExists(0, r.LoginID)
	if err != nil {
		return err
	}
	if exists {
		return errors.NewExistError("ログインIDが存在しています", r)
	}
	user, err := entity.NewUser(
		entity.UserNameOption(r.Name),
		entity.LoginIDOption(r.LoginID),
		entity.PasswordOption(r.Password),
	)
	if err != nil {
		return err
	}
	err = u.userRepository.Create(*user)
	return err
}
func (u userUseCase) Update(r UpdateUserRequest) error {
	// ログインID使用済みチェック
	exists, err := u.userService.CheckUserExists(r.ID, r.LoginID)
	if err != nil {
		return err
	}
	if exists {
		return errors.NewExistError("ログインIDが存在しています", r)
	}
	user, err := entity.NewUser(
		entity.UserIDOption(r.ID),
		entity.UserNameOption(r.Name),
		entity.LoginIDOption(r.LoginID),
		entity.PasswordOption(r.Password),
	)

	if err != nil {
		return err
	}

	err = u.userRepository.Update(*user)
	return err
}
func (u userUseCase) LoginAuthenticate(r LoginRequest) (*entity.User, error) {
	GetUser, err := u.userService.LoginAuthentication(r.LoginID, r.Password)
	if err != nil {
		return GetUser, err
	}
	if GetUser.ID == 0 {
		return GetUser, errors.NewLoginNotFoundError("ユーザーが見つかりませんでした", r.LoginID, r.Password)
	}
	return GetUser, nil

}
