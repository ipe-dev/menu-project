package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
)

func TestCreateFoodStuff(t *testing.T) {
	database.Connect()
	r := CreateFoodStuffRequest{
		Name:   "プチトマト",
		MenuID: 1,
	}
	r2 := CreateFoodStuffRequest{
		Name:   "ひき肉",
		MenuID: 1,
	}
	var requests []CreateFoodStuffRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkCreateFoodStuffRequest{requests}
	p := persistence.NewFoodStuffPersistence()
	foodStuffUseCase := NewFoodStuffUseCase(p)
	m, e := foodStuffUseCase.BulkCreate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestUpdateFoodStuff(t *testing.T) {
	database.Connect()
	r := UpdateFoodStuffRequest{
		ID:   1,
		Name: "レタス",
	}
	r2 := UpdateFoodStuffRequest{
		ID:   2,
		Name: "ニンジン",
	}
	var requests []UpdateFoodStuffRequest
	requests = append(requests, r)
	requests = append(requests, r2)
	bulkRequest := BulkUpdateFoodStuffRequest{requests}
	p := persistence.NewFoodStuffPersistence()
	menuUseCase := NewFoodStuffUseCase(p)
	m, e := menuUseCase.BulkUpdate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestGetFoodStuff(t *testing.T) {
	database.Connect()
	r := GetFoodStuffRequest{
		MenuID: 1,
	}
	p := persistence.NewFoodStuffPersistence()
	menuUseCase := NewFoodStuffUseCase(p)
	m, e := menuUseCase.Get(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		t.Log(m)
	}
}

func TestGetFoodStuffList(t *testing.T) {
	database.Connect()
	r := GetFoodStuffListRequest{
		MenuIDList: []int{1, 2, 3},
	}
	p := persistence.NewFoodStuffPersistence()
	menuUseCase := NewFoodStuffUseCase(p)
	m, e := menuUseCase.GetList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
func TestChangeFoodStuffStatus(t *testing.T) {
	database.Connect()
	r := ChangeFoodStuffStatusRequest{
		ID:     1,
		Status: 1,
	}
	p := persistence.NewFoodStuffPersistence()
	menuUseCase := NewFoodStuffUseCase(p)
	e := menuUseCase.ChangeBuyStatus(r)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	}
}
