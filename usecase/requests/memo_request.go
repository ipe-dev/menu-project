package requests

type CreateMemoRequest struct {
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
}
type UpdateMemoRequest struct {
	ID        int   `json:"id" binding:"required"`
	StartDate int64 `json:"start_date"`
	EndDate   int64 `json:"end_date"`
}
type GetMemoListRequest struct {
	UserID int `json:"user_id"`
}
type GetMemoRequest struct {
	ID     int `json:"id" binding:"required"`
	UserID int `json:"user_id"`
}
