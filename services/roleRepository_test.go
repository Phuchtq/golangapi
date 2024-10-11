package services

import (
	"testing"
	"v3/dbo"
	"v3/mocks"
	"v3/mocks/repos"
	"v3/mocks/samples"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllRolesResultOfLength3WithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//----------------------------------
	roleRepoMock.On("GetAllRoles").Return(samples.GetRolesMockData(), nil)
	//----------------------------------
	res, err := testTrg.GetAllRoles()
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetRolesMockData()), len(*res))
}

func TestGetRoleByIdInputExistedRoleResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//----------------------------------
	roleRepoMock.On("GetRoleById", mock.AnythingOfType("string")).Return(samples.GetExistedMockRole(), nil)
	res, err := testTrg.GetRoleById(mocks.Existed)
	expectedVal := mocks.Existed
	assert.NoError(t, err)
	assert.Equal(t, expectedVal, res.RoleId)
}

func TestGetRoleByIdInputNotExistedRoleResultOfEmptyWithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//----------------------------------
	var data *dbo.Role
	//----------------------------------
	roleRepoMock.On("GetRoleById", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	res, err := testTrg.roleRepo.GetRoleById(mocks.NotExisted)
	assert.NoError(t, err)
	assert.Nil(t, res)
}

func TestCreateRoleResultOfSuccessWithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//----------------------------------
	var err error
	//----------------------------------
	roleRepoMock.On("CreateRole", mock.AnythingOfType("dbo.Role")).Return(err)
	//----------------------------------
	res := testTrg.roleRepo.CreateRole(samples.GetExistedMockRole())
	assert.NoError(t, res)
}

func TestGetRolesByNameInputExistedNameResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//----------------------------------
	roleRepoMock.On("GetRolesByName", mock.AnythingOfType("string")).Return(samples.GetExistedMockRoleByExistedName(), nil)
	//----------------------------------
	res, err := testTrg.roleRepo.GetRolesByName(mocks.Existed)
	assert.NoError(t, err)
	assert.Equal(t, len(*res), len(*samples.GetExistedMockRoleByExistedName()))
}

func TestGetRolesByStatusInputValidStatusResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//-----------------------------------
	roleRepoMock.On("GetRolesByStatus", mocks.PositiveStatus).Return(samples.GetExistedMockRoleByStatus(mocks.PositiveStatus), nil)
	//-----------------------------------
	res, err := testTrg.roleRepo.GetRolesByStatus(mocks.PositiveStatus)
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetExistedMockRoleByStatus(mocks.PositiveStatus)), len(*res))
	//-----------------------------------
	roleRepoMock.On("GetRolesByStatus", mocks.NegativeStatus).Return(samples.GetExistedMockRoleByStatus(mocks.NegativeStatus), nil)
	//-----------------------------------
	res2, err2 := testTrg.roleRepo.GetRolesByStatus(mocks.NegativeStatus)
	assert.NoError(t, err2)
	assert.Equal(t, len(*samples.GetExistedMockRoleByStatus(mocks.NegativeStatus)), len(*res2))
}

func TestUpdateRoleWithValidInputResultSuccessWithNoErrorReturnsWell(t *testing.T) {
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := roleService{
		roleRepo: &roleRepoMock,
	}
	//-----------------------------------
	var errorResponse error
	roleRepoMock.On("UpdateRole", mock.AnythingOfType("dbo.Role")).Return(errorResponse)
	err := testTrg.roleRepo.UpdateRole(samples.GetUpdatedMockRole())
	assert.NoError(t, err)
}
