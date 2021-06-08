package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/infrastructure/database"
	persistance "github.com/ipe-dev/menu_project/infrastructure/persistence"
)

func TestCreateSubMenu(t *testing.T) {
	database.Connect()
	r := CreateSubMenuRequest{
		Name:   "大根おろし",
		MenuID: 3,
	}
	r2 := CreateSubMenuRequest{
		Name:   "卵焼き",
		MenuID: 1,
	}
	var requests []CreateSubMenuRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkCreateSubMenuRequest{requests}
	p := persistance.NewSubMenuPersistance()
	menuUseCase := NewSubMenuUseCase(p)
	m, e := menuUseCase.BulkCreateSubMenu(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestUpdateSubMenu(t *testing.T) {
	database.Connect()
	r := UpdateSubMenuRequest{
		ID:   1,
		Name: "aaa",
	}
	r2 := UpdateSubMenuRequest{
		ID:   2,
		Name: "お茶漬け",
	}
	var requests []UpdateSubMenuRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkUpdateSubMenuRequest{requests}
	p := persistance.NewSubMenuPersistance()
	menuUseCase := NewSubMenuUseCase(p)
	m, e := menuUseCase.BulkUpdateSubMenu(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestGetSubMenu(t *testing.T) {
	database.Connect()
	r := GetSubMenuRequest{
		ID:     1,
		MenuID: 1,
	}
	p := persistance.NewSubMenuPersistance()
	menuUseCase := NewSubMenuUseCase(p)
	m, e := menuUseCase.GetSubMenu(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		t.Log(m)
	}
}
func TestGetSubMenuList(t *testing.T) {
	database.Connect()
	r := GetSubMenuListRequest{
		MenuIDList: []int{1, 2, 3},
	}
	p := persistance.NewSubMenuPersistance()
	menuUseCase := NewSubMenuUseCase(p)
	m, e := menuUseCase.GetSubMenuList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
