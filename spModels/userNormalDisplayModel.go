package spModels

type UserNormalModel struct {
	UserId       string `json:"user_id"`
	RoleId       string `json:"role_id"`
	Email        string `json:"email" validate:"email, required"`
	Pasword      string `json:"pasword" validate:"required, min=8"`
	ActiveStatus string `json:"active_status"`
}
