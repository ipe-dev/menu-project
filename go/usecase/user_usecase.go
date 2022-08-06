package usecase

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/domain/service"
	"github.com/ipe-dev/menu_project/errors"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

type UserUseCase interface {
	Get(requests.GetUserRequest) (*entity.User, error)
	Create(requests.CreateUserRequest) error
	Update(requests.UpdateUserRequest) error
	LoginAuthenticate(requests.LoginRequest) (*entity.User, error)
}
type userUseCase struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewUserUseCase(r repository.UserRepository, s service.UserService) UserUseCase {
	return &userUseCase{r, s}
}

func (u userUseCase) Get(r requests.GetUserRequest) (*entity.User, error) {
	user, err := u.userRepository.Get(r.ID)
	return user, err
}

func (u userUseCase) Create(r requests.CreateUserRequest) error {

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
func (u userUseCase) Update(r requests.UpdateUserRequest) error {
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
func (u userUseCase) LoginAuthenticate(r requests.LoginRequest) (*entity.User, error) {
	GetUser, err := u.userService.LoginAuthentication(r.LoginID, r.Password)
	if err != nil {
		return GetUser, err
	}
	if GetUser.ID == 0 {
		return GetUser, errors.NewLoginNotFoundError("ユーザーが見つかりませんでした", r.LoginID, r.Password)
	}
	return GetUser, nil

}
