package requests

type GetFoodStuffRequest struct {
	MenuID int `json:"menu_id" validate:"required"`
}
type GetFoodStuffListRequest struct {
	MenuIDList []int `json:"menu_id_list"`
}
type CreateFoodStuffRequest struct {
	Name   string `json:"name" validate:"required"`
	MenuID int    `json:"menu_id" validate:"required"`
}
type BulkCreateFoodStuffRequest struct {
	CreateRequests []CreateFoodStuffRequest
}
type UpdateFoodStuffRequest struct {
	ID     int    `json:"id" validate:"required"`
	Name   string `json:"name"`
	MenuID int    `json:"menu_id" validate:"required"`
}
type BulkUpdateFoodStuffRequest struct {
	UpdateRequests []UpdateFoodStuffRequest
}
type ChangeFoodStuffStatusRequest struct {
	ID     int
	Status int
}
