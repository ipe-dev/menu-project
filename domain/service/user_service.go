package service

import (
	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
	"github.com/ipe-dev/menu_project/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CheckUserExists(ID int, LoginID string) (bool, error)
	LoginAuthentication(LoginID string, Password string) (*entity.User, error)
}
type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{r}
}

func (s userService) CheckUserExists(ID int, LoginID string) (bool, error) {
	user, err := s.UserRepository.GetByLoginID(LoginID)
	if err != nil {
		return false, errors.NewInfraError(err, LoginID)
	}
	if ID == 0 {
		// Create
		return user.ID != 0, nil
	} else {
		// Update
		return user.ID != 0 && user.ID != ID, nil
	}
}
func (s userService) LoginAuthentication(LoginID string, Password string) (*entity.User, error) {
	GetUser, err := s.UserRepository.GetByLoginID(LoginID)
	hash := GetUser.Password
	compareError := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	if compareError != nil {
		return GetUser, errors.NewLoginPasswordError(compareError, hash, Password)
	}
	return GetUser, err
}
