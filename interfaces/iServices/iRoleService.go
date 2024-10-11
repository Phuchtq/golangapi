package iservices

import "v3/dbo"

type IRoleService interface {
	GetAllRoles() (*[]dbo.Role, error)
	GetRolesByName(name string) (*[]dbo.Role, error)
	GetRolesByStatus(rawStatus string) (*[]dbo.Role, error)
	GetRoleById(id string) (*dbo.Role, error)
	CreateRole(name string) error
	UpdateRole(x dbo.Role) error
	RemoveRole(id string) error
	ActivateRole(id string) error
}
