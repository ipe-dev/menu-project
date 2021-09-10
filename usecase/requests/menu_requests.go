package requests

type CreateMenuRequest struct {
	Name string `json:"name"`
	Date int64  `json:"date" binding:"required"`
	Kind int    `json:"kind" binding:"required"`
	URL  string `json:"url"`
}
type UpdateMenuRequest struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name"`
	Date int64  `json:"date" binding:"required"`
	Kind int    `json:"kind" binding:"required"`
	URL  string `json:"url"`
}
type GetMenuListRequest struct {
	MemoID int `json:"memo_id" binding:"required"`
}
type GetMenuRequest struct {
	ID int `json:"id" binding:"required"`
}
type BulkCreateMenuRequest struct {
	MemoID         int `json:"memo_id"`
	CreateRequests []CreateMenuRequest
}
type BulkUpdateMenuRequest struct {
	MemoID         int `json:"memo_id"`
	UpdateRequests []UpdateMenuRequest
}
