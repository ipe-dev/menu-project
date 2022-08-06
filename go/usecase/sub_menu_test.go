package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

func TestCreateSubMenu(t *testing.T) {
	database.Connect()
	r := requests.CreateSubMenuRequest{
		Name:   "大根おろし",
		MemoID: 3,
	}
	r2 := requests.CreateSubMenuRequest{
		Name:   "卵焼き",
		MemoID: 1,
	}
	var req []requests.CreateSubMenuRequest
	req = append(req, r)
	req = append(req, r2)
	bulkRequest := requests.BulkCreateSubMenuRequest{req}
	p := persistence.NewSubMenuPersistence()
	menuUseCase := NewSubMenuUseCase(p)
	e := menuUseCase.BulkCreate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println("success")
	}
}

func TestGetSubMenu(t *testing.T) {
	database.Connect()
	r := requests.GetSubMenuRequest{
		ID:     1,
		MemoID: 1,
	}
	p := persistence.NewSubMenuPersistence()
	menuUseCase := NewSubMenuUseCase(p)
	m, e := menuUseCase.Get(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		t.Log(m)
	}
}
func TestGetSubMenuList(t *testing.T) {
	database.Connect()
	r := requests.GetSubMenuListRequest{
		MemoIDList: []int{1, 2, 3},
	}
	p := persistence.NewSubMenuPersistence()
	menuUseCase := NewSubMenuUseCase(p)
	m, e := menuUseCase.GetList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
