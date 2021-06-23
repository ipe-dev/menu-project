package service

import (
	"log"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/entity/value"
	"github.com/ipe-dev/menu_project/domain/repository"
)

type UserService interface {
	CheckUserExists(LoginID string) (bool, error)
	LoginAuthentication(LoginID string, Password value.Password) (*entity.User, error)
}
type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{r}
}

func (s userService) CheckUserExists(LoginID string) (bool, error) {
	user, err := s.UserRepository.GetByLoginID(LoginID)
	if err != nil {
		log.Println(err)
		return false, err
	}

	return user != nil, err
}
func (s userService) LoginAuthentication(LoginID string, Password value.Password) (*entity.User, error) {
	GetUser, err := s.UserRepository.GetByLoginIDAndPassword(LoginID, Password)
	return GetUser, err
}
