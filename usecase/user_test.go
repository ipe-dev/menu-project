package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/domain/service"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
)

func TestCreateUser(t *testing.T) {
	database.Connect()
	r := CreateUserRequest{
		Name:     "あああああ",
		LoginID:  "test_id0000000000000003",
		Password: "test_password3",
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
	r := UpdateUserRequest{
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
	r := GetUserRequest{
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
	r := LoginRequest{
		LoginID:  "test_id222",
		Password: "test_password",
	}
	p := persistence.NewUserPersistence()
	s := service.NewUserService(p)
	menuUseCase := NewUserUseCase(p, s)
	m, b := menuUseCase.LoginAuthenticate(r)
	if !b {
		t.Fatalf("failed test %#v", b)
	} else {
		t.Log(m)
	}
}
