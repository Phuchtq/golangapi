package repos

import (
	"v3/dbo"

	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct {
	mock.Mock
}

func (tr *UserRepoMock) GetAllUsers() (*[]dbo.User, error) {
	mockData := tr.Called()
	//----------------------------------
	var res1 *[]dbo.User
	if mockFunc, ok := mockData.Get(0).(func() *[]dbo.User); ok {
		res1 = mockFunc()
	} else {
		res1 = mockData.Get(0).(*[]dbo.User)
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

func (tr *UserRepoMock) GetUsersByRole(role string) (*[]dbo.User, error) {
	mockData := tr.Called(role)
	//----------------------------------
	var res1 *[]dbo.User
	if mockFunc, ok := mockData.Get(0).(func(string) *[]dbo.User); ok {
		res1 = mockFunc(role)
	} else {
		res1 = mockData.Get(0).(*[]dbo.User)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		res2 = mockFunc(role)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *UserRepoMock) GetUserById(id string) (*dbo.User, error) {
	mockData := tr.Called(id)
	//----------------------------------
	var res1 *dbo.User
	if mockFunc, ok := mockData.Get(0).(func(string) *dbo.User); ok {
		res1 = mockFunc(id)
	} else {
		res1 = mockData.Get(0).(*dbo.User)
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

func (tr *UserRepoMock) GetUsersByStatus(status bool) (*[]dbo.User, error) {
	mockData := tr.Called(status)
	//----------------------------------
	var res1 *[]dbo.User
	if mockFunc, ok := mockData.Get(0).(func(bool) *[]dbo.User); ok {
		res1 = mockFunc(status)
	} else {
		res1 = mockData.Get(0).(*[]dbo.User)
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

func (tr *UserRepoMock) GetUserByEmail(email string) (*dbo.User, error) {
	mockData := tr.Called(email)
	//----------------------------------
	var res1 *dbo.User
	if mockFunc, ok := mockData.Get(0).(func(string) *dbo.User); ok {
		res1 = mockFunc(email)
	} else {
		res1 = mockData.Get(0).(*dbo.User)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		res2 = mockFunc(email)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *UserRepoMock) AddUser(r dbo.User) error {
	mockData := tr.Called(r)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(dbo.User) error); ok {
		return mockFunc(r)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *UserRepoMock) UpdateUser(r dbo.User) error {
	mockData := tr.Called(r)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(dbo.User) error); ok {
		return mockFunc(r)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *UserRepoMock) ChangeUserStatus(status bool, id string) error {
	mockData := tr.Called(status, id)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(bool, string) error); ok {
		return mockFunc(status, id)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}
