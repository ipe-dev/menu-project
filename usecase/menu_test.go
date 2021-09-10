package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/ipe-dev/menu_project/domain/entity"
	"github.com/ipe-dev/menu_project/domain/service"
	"github.com/ipe-dev/menu_project/infrastructure/database"
	persistence "github.com/ipe-dev/menu_project/infrastructure/persistence"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

func TestCreateMenu(t *testing.T) {
	database.Connect()
	// date := time.Now().AddDate(0, 0, 1).Unix()
	date := time.Now().Unix()
	memoID := 1
	r := requests.CreateMenuRequest{
		Name: "ハンバーグ",
		Date: date,
		Kind: entity.MenuKindLunch,
		URL:  "https://www.google.com/",
	}
	r2 := requests.CreateMenuRequest{
		Name: "寿司",
		Date: date,
		Kind: entity.MenuKindDinner,
		URL:  "https://www.google.com/",
	}
	var req []requests.CreateMenuRequest
	req = append(req, r)
	req = append(req, r2)
	bulkRequest := requests.BulkCreateMenuRequest{memoID, req}
	p := persistence.NewMenuPersistence()
	p2 := persistence.NewMemoPersistence()
	s := service.NewMemoService(p2)
	menuUseCase := NewMenuUseCase(p, p2, s)
	e := menuUseCase.BulkCreate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println("success")
	}
}
func TestUpdateMenu(t *testing.T) {
	database.Connect()
	r := requests.UpdateMenuRequest{
		ID:   1,
		Name: "タコライス",
		Kind: entity.MenuKindLunch,
		URL:  "https://www.google.com/",
	}
	var req []requests.UpdateMenuRequest
	memoID := 1
	req = append(req, r)
	bulkRequest := requests.BulkUpdateMenuRequest{memoID, req}
	p := persistence.NewMenuPersistence()
	p2 := persistence.NewMemoPersistence()
	s := service.NewMemoService(p2)
	menuUseCase := NewMenuUseCase(p, p2, s)

	e := menuUseCase.BulkUpdate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println("success")
	}
}
func TestGetMenu(t *testing.T) {
	database.Connect()
	r := requests.GetMenuRequest{
		ID: 1,
	}
	p := persistence.NewMenuPersistence()
	p2 := persistence.NewMemoPersistence()
	s := service.NewMemoService(p2)
	menuUseCase := NewMenuUseCase(p, p2, s)
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
	r := requests.GetMenuListRequest{
		MemoID: 1,
	}
	p := persistence.NewMenuPersistence()
	p2 := persistence.NewMemoPersistence()
	s := service.NewMemoService(p2)
	menuUseCase := NewMenuUseCase(p, p2, s)
	m, e := menuUseCase.GetList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
