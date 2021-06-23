package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
)

func TestCreateSubFoodStuff(t *testing.T) {
	database.Connect()
	r := CreateSubFoodStuffRequest{
		Name:      "じゃがいも",
		SubMenuID: 1,
	}
	r2 := CreateSubFoodStuffRequest{
		Name:      "にんじん",
		SubMenuID: 1,
	}
	var requests []CreateSubFoodStuffRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkCreateSubFoodStuffRequest{requests}
	p := persistence.NewSubFoodStuffPersistence()
	subFoodStuffUseCase := NewSubFoodStuffUseCase(p)
	m, e := subFoodStuffUseCase.BulkCreate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestUpdateSubFoodStuff(t *testing.T) {
	database.Connect()
	r := UpdateSubFoodStuffRequest{
		ID:   1,
		Name: "たまねぎ",
	}
	r2 := UpdateSubFoodStuffRequest{
		ID:   2,
		Name: "ちくわ",
	}
	var requests []UpdateSubFoodStuffRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkUpdateSubFoodStuffRequest{requests}
	p := persistence.NewSubFoodStuffPersistence()
	menuUseCase := NewSubFoodStuffUseCase(p)
	m, e := menuUseCase.BulkUpdate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestGetSubFoodStuff(t *testing.T) {
	database.Connect()
	r := GetSubFoodStuffRequest{
		SubMenuID: 1,
	}
	p := persistence.NewSubFoodStuffPersistence()
	menuUseCase := NewSubFoodStuffUseCase(p)
	m, e := menuUseCase.Get(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		t.Log(m)
	}
}
func TestGetSubFoodStuffList(t *testing.T) {
	database.Connect()
	r := GetSubFoodStuffListRequest{
		SubMenuIDList: []int{1, 2, 3},
	}
	p := persistence.NewSubFoodStuffPersistence()
	menuUseCase := NewSubFoodStuffUseCase(p)
	m, e := menuUseCase.GetList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestChangeBuyStatus(t *testing.T) {
	database.Connect()
	r := ChangeSubBuyStatusRequest{
		ID:     1,
		Status: 1,
	}
	p := persistence.NewSubFoodStuffPersistence()
	usecase := NewSubFoodStuffUseCase(p)
	e := usecase.ChangeStatus(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(e)
	}
}
