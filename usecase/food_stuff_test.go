package usecase

import (
	"fmt"
	"testing"

	"github.com/ipe-dev/menu_project/infrastructure/database"
	"github.com/ipe-dev/menu_project/infrastructure/persistence"
	"github.com/ipe-dev/menu_project/usecase/requests"
)

func TestCreateFoodStuff(t *testing.T) {
	database.Connect()
	r := requests.CreateFoodStuffRequest{
		Name:   "プチトマト",
		MenuID: 1,
	}
	r2 := requests.CreateFoodStuffRequest{
		Name:   "ひき肉",
		MenuID: 1,
	}
	var req []requests.CreateFoodStuffRequest
	req = append(req, r)
	req = append(req, r2)
	bulkRequest := requests.BulkCreateFoodStuffRequest{req}
	p := persistence.NewFoodStuffPersistence()
	foodStuffUseCase := NewFoodStuffUseCase(p)
	e := foodStuffUseCase.BulkCreate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println("success")
	}
}
func TestUpdateFoodStuff(t *testing.T) {
	database.Connect()
	r := requests.UpdateFoodStuffRequest{
		ID:   1,
		Name: "レタス",
	}
	r2 := requests.UpdateFoodStuffRequest{
		ID:   2,
		Name: "ニンジン",
	}
	var req []requests.UpdateFoodStuffRequest
	req = append(req, r)
	req = append(req, r2)
	bulkRequest := requests.BulkUpdateFoodStuffRequest{req}
	p := persistence.NewFoodStuffPersistence()
	menuUseCase := NewFoodStuffUseCase(p)
	e := menuUseCase.BulkUpdate(bulkRequest)
	if e != nil {
		t.Fatalf("failed test %#v", e)
	} else {
		fmt.Println("success")
	}
}
func TestGetFoodStuff(t *testing.T) {
	database.Connect()
	r := requests.GetFoodStuffRequest{
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
	r := requests.GetFoodStuffListRequest{
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
	r := requests.ChangeFoodStuffStatusRequest{
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
