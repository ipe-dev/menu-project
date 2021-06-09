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
	m, e := foodStuffUseCase.BulkCreateFoodStuff(bulkRequest)
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
	m, e := menuUseCase.BulkUpdateFoodStuff(bulkRequest)
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
	m, e := menuUseCase.GetFoodStuff(r)
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
	m, e := menuUseCase.GetFoodStuffList(r)
	t.Log(m)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println(m)
	}
}
