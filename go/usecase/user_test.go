package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/domain/service"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

func TestCreateUser(t *testing.T) {
	database.Connect()
	r := requests.CreateUserRequest{
		Name:     "田中太郎",
		LoginID:  "new_login",
		Password: "taro_password",
	}
	p := persistence.NewUserPersistence()
	s := service.NewUserService(p)
	userUseCase := NewUserUseCase(p, s)
	e := userUseCase.Create(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	}
}
func TestUpdateUser(t *testing.T) {
	database.Connect()
	r := requests.UpdateUserRequest{
		ID:       1,
		Name:     "user_name_update",
		LoginID:  "test_id222",
		Password: "test_password",
	}
	p := persistence.NewUserPersistence()
	s := service.NewUserService(p)
	userUseCase := NewUserUseCase(p, s)
	e := userUseCase.Update(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	}
}
func TestGetUser(t *testing.T) {
	database.Connect()
	r := requests.GetUserRequest{
		ID: 1,
	}
	p := persistence.NewUserPersistence()
	s := service.NewUserService(p)
	userUseCase := NewUserUseCase(p, s)
	m, e := userUseCase.Get(r)
	fmt.Println(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		t.Log(m)
	}
}

func TestLoginUser(t *testing.T) {
	database.Connect()
	r := requests.LoginRequest{
		LoginID:  "test_id222",
		Password: "test_password",
	}
	p := persistence.NewUserPersistence()
	s := service.NewUserService(p)
	menuUseCase := NewUserUseCase(p, s)
	m, e := menuUseCase.LoginAuthenticate(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		t.Log(m)
	}
}
