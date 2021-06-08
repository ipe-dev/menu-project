package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	persistance "github.com/ipe-dev/menu_project/infrastructure/persistence"
)

func TestCreateMenu(t *testing.T) {
	database.Connect()
	date := time.Now().AddDate(0, 0, 1).Unix()
	r := CreateMenuRequest{
		// Name:   "ハンバーグ",
		Date:   date,
		Kind:   entity.MenuKindLunch,
		URL:    "https://www.google.com/",
		UserID: 1,
		WeekID: 1,
	}
	r2 := CreateMenuRequest{
		Name:   "寿司",
		Date:   date,
		Kind:   entity.MenuKindDinner,
		URL:    "https://www.google.com/",
		UserID: 1,
		WeekID: 1,
	}
	var requests []CreateMenuRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkCreateMenuRequest{requests}
	p := persistance.NewMenuPersistance()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.BulkCreateMenu(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestUpdateMenu(t *testing.T) {
	database.Connect()
	r := UpdateMenuRequest{
		ID:     1,
		Name:   "タコライス",
		Kind:   entity.MenuKindLunch,
		URL:    "https://www.google.com/",
		UserID: 1,
		WeekID: 1,
	}
	var requests []UpdateMenuRequest
	requests = append(requests, r)
	bulkRequest := BulkUpdateMenuRequest{requests}
	p := persistance.NewMenuPersistance()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.BulkUpdateMenu(bulkRequest)
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
		UserID: 1,
	}
	p := persistance.NewMenuPersistance()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.GetMenu(r)
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
		WeekID: 1,
		UserID: 1,
	}
	p := persistance.NewMenuPersistance()
	menuUseCase := NewMenuUseCase(p)
	m, e := menuUseCase.GetMenuList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
