package controllers

import (
	servicegenerator "v3/service_generator"
)

func fetchRoles() map[string]string {
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		return map[string]string{
			"Admin":    "R001",
			"Staff":    "R002",
			"Customer": "R003",
		}
	}
	//-----------------------------------------
	list, err := service.GetAllRoles()
	if err != nil || list == nil {
		return nil
	}
	//-----------------------------------------
	res := make(map[string]string)
	for _, v := range *list {
		res[v.RoleName] = v.RoleId
	}
	//-----------------------------------------
	return res
}

func isAdminAccess(role string) bool { // This func mostly used for admin case
	if role != fetchRoles()["Admin"] { // Actor is not admin
		return false
	}
	//---------------------------------------------
	return true
}
