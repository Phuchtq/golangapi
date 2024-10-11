package repos

import (
	"v3/dbo"

	"github.com/stretchr/testify/mock"
)

type RoleRepoMock struct {
	mock.Mock
}

func (tr *RoleRepoMock) GetAllRoles() (*[]dbo.Role, error) {
	mockData := tr.Called()
	//----------------------------------
	var res1 *[]dbo.Role
	if mockFunc, ok := mockData.Get(0).(func() *[]dbo.Role); ok {
		res1 = mockFunc()
	} else {
		res1 = mockData.Get(0).(*[]dbo.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(1).(func() error); ok {
		res2 = mockFunc()
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *RoleRepoMock) GetRolesByName(name string) (*[]dbo.Role, error) {
	mockData := tr.Called(name)
	//----------------------------------
	var res1 *[]dbo.Role
	if mockFunc, ok := mockData.Get(0).(func(string) *[]dbo.Role); ok {
		res1 = mockFunc(name)
	} else {
		res1 = mockData.Get(0).(*[]dbo.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		res2 = mockFunc(name)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *RoleRepoMock) GetRolesByStatus(status bool) (*[]dbo.Role, error) {
	mockData := tr.Called(status)
	//----------------------------------
	var res1 *[]dbo.Role
	if mockFunc, ok := mockData.Get(0).(func(bool) *[]dbo.Role); ok {
		res1 = mockFunc(status)
	} else {
		res1 = mockData.Get(0).(*[]dbo.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(bool) error); ok {
		res2 = mockFunc(status)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *RoleRepoMock) GetRoleById(id string) (*dbo.Role, error) {
	mockData := tr.Called(id)
	//----------------------------------
	var res1 *dbo.Role
	if mockFunc, ok := mockData.Get(0).(func(string) *dbo.Role); ok {
		res1 = mockFunc(id)
	} else {
		res1 = mockData.Get(0).(*dbo.Role)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(1).(func(string) error); ok {
		res2 = mockFunc(id)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *RoleRepoMock) CreateRole(r dbo.Role) error {
	mockData := tr.Called(r)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(dbo.Role) error); ok {
		return mockFunc(r)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *RoleRepoMock) UpdateRole(r dbo.Role) error {
	mockData := tr.Called(r)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(dbo.Role) error); ok {
		return mockFunc(r)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *RoleRepoMock) RemoveRole(id string) error {
	mockData := tr.Called(id)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		return mockFunc(id)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *RoleRepoMock) ActivateRole(id string) error {
	mockData := tr.Called(id)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		return mockFunc(id)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}
