package constants

import iservices "v3/interfaces/iServices"

var Roles = map[string]string{
	"Admin":    "Admin",
	"Staff":    "Staff",
	"Customer": "Customer",
}

func FetchRoles(service *iservices.IRoleService) map[string]string {
	list, err := (*service).GetAllRoles()
	if err != nil || list == nil {
		return nil
	}
	//-----------------------------------------
	var res map[string]string
	for _, v := range *list {
		res[v.RoleName] = v.RoleId
	}
	//-----------------------------------------
	return res
}
