package service

import (
	"fmt"
	"log"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/repository"
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
		log.Println(err)
		return false, err
	}
	if ID == 0 {
		// Create
		return user.ID != 0, err
	} else {
		// Update
		return user.ID != 0 && user.ID != ID, err
	}
}
func (s userService) LoginAuthentication(LoginID string, Password string) (*entity.User, error) {
	GetUser, err := s.UserRepository.GetByLoginID(LoginID)
	hash := GetUser.Password
	compareError := bcrypt.CompareHashAndPassword([]byte(hash), []byte(Password))
	fmt.Println(GetUser.Password)
	if compareError != nil {
		return GetUser, compareError
	}
	return GetUser, err
}
