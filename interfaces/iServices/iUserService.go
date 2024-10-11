package iservices

import (
	"v3/dbo"
	"v3/spModels"
)

type IUserService interface {
	GetAllUsers() (*[]dbo.User, error)
	GetUsersByRole(role string) (*[]dbo.User, error)
	GetUsersByStatus(rawStatus string) (*[]dbo.User, error)
	GetUserById(id string) (*dbo.User, error)
	AddUser(u spModels.SignUpModel, actorId string) (error, string)
	UpdateUser(user spModels.UserNormalModel, actorId string) (string, error)
	ChangeUserStatus(rawStatus, userId, actorId string) (error, string)
	Login(email string, password string) (string, string, error)
	LogOut(userId string) error
	VerifyAction(rawToken string) (error, string)
	VerifyResetPassword(newPass, re_newPass, token string) (string, error)
	RecoverAccountByCustomer(email string) (string, error)
}
