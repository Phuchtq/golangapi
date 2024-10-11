package dbo

import (
	"time"
)

type User struct {
	UserId          string     `json:"user_id"`
	RoleId          string     `json:"role_id"`
	Email           string     `json:"email" validate:"email, required"`
	Pasword         string     `json:"pasword" validate:"required, min=8"`
	ActiveStatus    bool       `json:"active_status"`
	FailAccess      int        `json:"fail_access"`
	LastFail        *time.Time `json:"last_fail"`
	AccessToken     *string    `json:"access_token"`
	RefreshToken    *string    `json:"refresh_token"`
	ActionToken     *string    `json:"action_token"`
	ActionPeriod    *time.Time `json:"action_period"`
	IsHaveToResetPw *bool      `json:"is_have_to_reset_password"`
}
