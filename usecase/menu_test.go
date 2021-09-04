package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	persistence "github.com/ipe-dev/menu_project/infrastructure/persistence"
)

func TestCreateMenu(t *testing.T) {
	database.Connect()
	// date := time.Now().AddDate(0, 0, 1).Unix()
	date := time.Now().Unix()
	r := CreateMenuRequest{
		Name: "ハンバーグ",
		Date: date,
		Kind: entity.MenuKindLunch,
		URL:  "https://www.google.com/",
	}
	r2 := CreateMenuRequest{
		Name: "寿司",
		Date: date,
		Kind: entity.MenuKindDinner,
		URL:  "https://www.google.com/",
	}
	var requests []CreateMenuRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkCreateMenuRequest{requests}
	p := persistence.NewMenuPersistence()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.BulkCreate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestUpdateMenu(t *testing.T) {
	database.Connect()
	r := UpdateMenuRequest{
		ID:   1,
		Name: "タコライス",
		Kind: entity.MenuKindLunch,
		URL:  "https://www.google.com/",
	}
	var requests []UpdateMenuRequest
	requests = append(requests, r)
	bulkRequest := BulkUpdateMenuRequest{requests}
	p := persistence.NewMenuPersistence()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.BulkUpdate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestGetMenu(t *testing.T) {
	database.Connect()
	time := time.Date(2021, 6, 4, 0, 0, 0, 0, time.Local).Unix()
	r := GetMenuRequest{
		Date:   time,
		MemoID: 1,
	}
	p := persistence.NewMenuPersistence()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.Get(r)
	fmt.Println(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestGetMenuList(t *testing.T) {
	database.Connect()
	r := GetMenuListRequest{
		MemoID: 1,
		UserID: 1,
	}
	p := persistence.NewMenuPersistence()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.GetList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
