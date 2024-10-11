package irepositories

import "v3/dbo"

type IRoleRepo interface {
	GetAllRoles() (*[]dbo.Role, error)
	GetRolesByName(name string) (*[]dbo.Role, error)
	GetRolesByStatus(status bool) (*[]dbo.Role, error)
	GetRoleById(id string) (*dbo.Role, error)
	CreateRole(r dbo.Role) error
	RemoveRole(id string) error
	UpdateRole(r dbo.Role) error
	ActivateRole(id string) error
}
