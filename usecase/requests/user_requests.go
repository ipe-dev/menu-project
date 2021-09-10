package requests

type GetUserRequest struct {
	ID int `json:"id"`
}
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UpdateUserRequest struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginRequest struct {
	LoginID  string `json:"login_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LogoutRequest struct {
	ID int `json"id" binding:"required"`
}
