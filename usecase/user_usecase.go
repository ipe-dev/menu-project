package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/domain/service"
)

type UserUseCase interface {
	Get(GetUserRequest) (*entity.User, error)
	Create(CreateUserRequest) error
	Update(UpdateUserRequest) error
	LoginAuthenticate(LoginRequest) (*entity.User, bool)
}
type userUseCase struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewUserUseCase(r repository.UserRepository, s service.UserService) userUseCase {
	return &userUseCase{r, s}
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

func (u userUseCase) Get(r GetUserRequest) (*entity.User, error) {
	user, err := u.userRepository.Get(r.ID)
	return user, err
}

func (u userUseCase) Create(r CreateUserRequest) error {

	// ログインID使用済みチェック
	exists, err := u.userService.CheckUserExists(r.LoginID)
	if err != nil {
		return err
	}
	if exists {
		// TODO: カスタムエラー作る
		return err
	}
	user := entity.NewUser(
		entity.UserNameOption(r.Name),
		entity.LoginIDOption(r.LoginID),
		entity.PasswordOption(r.Password),
	)
	err = u.userRepository.Create(*user)
	return err
}
func (u userUseCase) Update(r UpdateUserRequest) error {
	// ログインID使用済みチェック
	exists, err := u.userService.CheckUserExists(r.LoginID)
	if err != nil {
		return err
	}
	if exists {
		// TODO: カスタムエラー作る
		return err
	}
	user := entity.NewUser(
		entity.UserIDOption(r.ID),
		entity.UserNameOption(r.Name),
		entity.LoginIDOption(r.LoginID),
		entity.PasswordOption(r.Password),
	)

	err = u.userRepository.Update(*user)
	return err
}
func (u userUseCase) LoginAuthenticate(r LoginRequest) (*entity.User, bool) {
	Password := value.NewPassword(r.Password)
	GetUser, err := u.userService.LoginAuthentication(r.LoginID, Password)
	return GetUser, err

}
