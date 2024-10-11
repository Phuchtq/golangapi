package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"v3/constants/notis"
	"v3/dbo"
	irepositories "v3/interfaces/iRepositories"
	iservices "v3/interfaces/iServices"
	"v3/repositories"
	"v3/utils"
)

type roleService struct {
	roleRepo irepositories.IRoleRepo
}

// func InitializeRoleService(roleRepo irepositories.IRoleRepo) iservices.IRoleService {
// 	return &roleService{
// 		roleRepo: roleRepo,
// 	}
// }

func InitializeRoleService(db *sql.DB) iservices.IRoleService {
	return &roleService{
		roleRepo: repositories.InitializeRoleRepo(db),
	}
}
func (tr *roleService) GetAllRoles() (*[]dbo.Role, error) {
	return tr.roleRepo.GetAllRoles()
}

func (tr *roleService) GetRolesByName(name string) (*[]dbo.Role, error) {
	if trimStr := strings.TrimSpace(name); trimStr != "" {
		return tr.roleRepo.GetRolesByName(trimStr)
	}
	return tr.roleRepo.GetAllRoles()
}

func (tr *roleService) GetRolesByStatus(rawStatus string) (*[]dbo.Role, error) {
	status, err := utils.IsStatusValid(rawStatus)
	if err != nil {
		return nil, errors.New(notis.InvalidStatusWarnMsg)
	}
	return tr.roleRepo.GetRolesByStatus(status)
}

func (tr *roleService) GetRoleById(id string) (*dbo.Role, error) {
	if id == "" {
		return nil, errors.New(notis.GenericsErrorWarnMsg)
	}
	//---------------------------------------
	res, err := tr.roleRepo.GetRoleById(id)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New(notis.UndefinedRoleWarnMsg)
	}
	return res, nil
}

func (tr *roleService) CreateRole(name string) error {
	if name := strings.TrimSpace(name); name == "" {
		return errors.New(notis.NameEmptyWarnMsg)
	}
	//---------------------------------------
	id, err := generateRoleId(&tr.roleRepo)
	if err != nil {
		return err
	}
	//---------------------------------------
	return tr.roleRepo.CreateRole(dbo.Role{
		RoleId:       id,
		RoleName:     name,
		ActiveStatus: true,
	})
}

func (tr *roleService) UpdateRole(x dbo.Role) error {
	res, err := tr.roleRepo.GetRoleById(x.RoleId)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New(notis.UndefinedRoleWarnMsg)
	}
	if x.RoleName != "" {
		res.RoleName = strings.TrimSpace(x.RoleName)
	}
	return tr.roleRepo.UpdateRole(*res)
}

func (tr *roleService) RemoveRole(id string) error {
	return tr.roleRepo.RemoveRole(id)
}

func (tr *roleService) ActivateRole(id string) error {
	return tr.roleRepo.ActivateRole(id)
}

func generateRoleId(repo *irepositories.IRoleRepo) (string, error) {
	list, err := (*repo).GetAllRoles()
	if err != nil {
		return "", err
	}
	//-----------------------------------
	return "R" + fmt.Sprintf("%03d", len(*list)+1), nil
}

func isRoleExisted(role string, repo *irepositories.IRoleRepo) (bool, error) {
	list, err := (*repo).GetAllRoles()
	if err != nil {
		return false, err
	}
	//-----------------------------------
	for _, v := range *list {
		if role == v.RoleId {
			return true, nil
		}
	}
	//-----------------------------------
	return false, nil
}
