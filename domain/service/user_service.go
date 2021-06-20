package service

import (
	"log"

	"github.com/ipe-dev/menu_project/domain/repository"
)

type UserService interface {
	CheckUserExists(LoginID string) bool
}
type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{r}
}

func (s userService) CheckUserExists(LoginID string) bool {
	user, err := s.UserRepository.GetByLoginID(LoginID)
	if err != nil {
		log.Println(err)
		return false
	}

	return user != nil
}
