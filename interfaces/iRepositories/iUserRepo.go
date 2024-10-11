package irepositories

import "v3/dbo"

type IUserRepo interface {
	GetAllUsers() (*[]dbo.User, error)
	GetUsersByRole(id string) (*[]dbo.User, error)
	GetUsersByStatus(status bool) (*[]dbo.User, error)
	GetUserById(id string) (*dbo.User, error)
	GetUserByEmail(email string) (*dbo.User, error)
	AddUser(u dbo.User) error
	UpdateUser(u dbo.User) error
	ChangeUserStatus(status bool, id string) error
}
