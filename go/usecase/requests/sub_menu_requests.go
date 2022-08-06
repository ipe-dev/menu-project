package requests

type GetSubMenuRequest struct {
	ID     int `json:"id"　validate:"required"`
	MemoID int `json:"memo_id"　validate:"required"`
}
type GetSubMenuListRequest struct {
	MemoIDList []int `json:"memo_id_list"`
}
type CreateSubMenuRequest struct {
	ID     int    `json:"id"　validate:"required"`
	Name   string `json:"name"　validate:"required"`
	MemoID int    `json:"memo_id"　validate:"required"`
}
type BulkCreateSubMenuRequest struct {
	CreateRequests []CreateSubMenuRequest
}
type UpdateSubMenuRequest struct {
	ID     int    `json:"id"　validate:"required"`
	Name   string `json:"name"`
	MemoID int    `json:"memo_id"　validate:"required"`
}
type BulkUpdateSubMenuRequest struct {
	UpdateRequests []UpdateSubMenuRequest
}
