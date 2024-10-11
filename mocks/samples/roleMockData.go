package samples

import (
	"v3/dbo"
	"v3/mocks"
)

func GetRolesMockData() *[]dbo.Role {
	return &[]dbo.Role{
		{
			RoleId:       "1",
			RoleName:     "1",
			ActiveStatus: true,
		},
		//----------------------------------
		{
			RoleId:       "2",
			RoleName:     "2",
			ActiveStatus: false,
		},
		//----------------------------------
		{
			RoleId:       "3",
			RoleName:     "3",
			ActiveStatus: true,
		},
	}
}

func GetMockStandardRoles() *[]dbo.Role {
	return &[]dbo.Role{
		{
			RoleId: mocks.Roles["Admin"],
		},
		//----------------------------------
		{
			RoleId: mocks.Roles["Staff"],
		},
		//----------------------------------
		{
			RoleId: mocks.Roles["Customer"],
		},
	}
}

func GetExistedMockRole() dbo.Role {
	for _, v := range *GetRolesMockData() {
		if mocks.Existed == v.RoleId {
			return v
		}
	}
	//--------------------------------------
	return dbo.Role{ // In case test data has some problems
		RoleId: mocks.Existed,
	}
}

func GetExistedMockRoleByExistedName() *[]dbo.Role {
	var res []dbo.Role
	for _, v := range *GetRolesMockData() {
		if mocks.Existed == v.RoleName {
			res = append(res, v)
		}
	}
	//--------------------------------------
	if len(res) > 0 {
		return &res
	}
	//--------------------------------------
	return &[]dbo.Role{ // In case test data has some problems
		{
			RoleId:       mocks.Existed,
			RoleName:     mocks.Existed,
			ActiveStatus: true,
		},
	}
}

func GetExistedMockRoleByStatus(status bool) *[]dbo.Role {
	var res []dbo.Role
	for _, v := range *GetRolesMockData() {
		if v.ActiveStatus == status {
			res = append(res, v)
		}
	}
	//--------------------------------------
	if len(res) > 1 {
		return &res
	}
	//--------------------------------------
	return &[]dbo.Role{ // In case test data has some problems
		{
			RoleId:       "1",
			RoleName:     "1",
			ActiveStatus: status,
		},
	}
}

func GetUpdatedMockRole() dbo.Role {
	var res dbo.Role
	for _, v := range *GetRolesMockData() {
		if v.RoleId == mocks.Existed {
			res = v
			break
		}
	}
	//--------------------------------------
	if res.RoleId == mocks.Existed {
		res.RoleName = mocks.UpdatedName
		return res
	}
	//--------------------------------------
	return dbo.Role{ // In case test data has some problems
		RoleId:       mocks.Existed,
		RoleName:     mocks.UpdatedName,
		ActiveStatus: false,
	}
}
