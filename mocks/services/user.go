package services

import (
	"v3/dbo"
	"v3/spModels"

	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (tr *UserServiceMock) GetAllUsers() (*[]dbo.User, error) {
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

func (tr *UserServiceMock) GetUsersByRole(role string) (*[]dbo.User, error) {
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

func (tr *UserServiceMock) GetUserById(id string) (*dbo.User, error) {
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

func (tr *UserServiceMock) GetUsersByStatus(rawStatus string) (*[]dbo.User, error) {
	mockData := tr.Called(rawStatus)
	//----------------------------------
	var res1 *[]dbo.User
	if mockFunc, ok := mockData.Get(0).(func(string) *[]dbo.User); ok {
		res1 = mockFunc(rawStatus)
	} else {
		res1 = mockData.Get(0).(*[]dbo.User)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		res2 = mockFunc(rawStatus)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *UserServiceMock) AddUser(u spModels.SignUpModel, actorId string) error {
	mockData := tr.Called(u, actorId)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(spModels.SignUpModel, string) error); ok {
		return mockFunc(u, actorId)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *UserServiceMock) UpdateUser(user spModels.UserNormalModel, actorId string) error {
	mockData := tr.Called(user, actorId)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(spModels.UserNormalModel, string) error); ok {
		return mockFunc(user, actorId)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *UserServiceMock) ChangeUserStatus(rawStatus, userId, actorId string) error {
	mockData := tr.Called(rawStatus, userId, actorId)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string, string, string) error); ok {
		return mockFunc(rawStatus, userId, actorId)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *UserServiceMock) Login(email, password string) (string, string, error) {
	mockData := tr.Called(email, password)
	//----------------------------------
	var res1 string
	if mockFunc, ok := mockData.Get(0).(func(string, string) string); ok {
		res1 = mockFunc(email, password)
	} else {
		res1 = mockData.Get(0).(string)
	}
	//----------------------------------
	var res2 string
	if mockFunc, ok := mockData.Get(1).(func(string, string) string); ok {
		res2 = mockFunc(email, password)
	} else {
		res2 = mockData.Get(1).(string)
	}
	//----------------------------------
	var res3 error
	if mockFunc, ok := mockData.Get(2).(func(string, string) error); ok {
		res3 = mockFunc(email, password)
	} else {
		res3 = mockData.Error(2)
	}
	//----------------------------------
	return res1, res2, res3
}

func (tr *UserServiceMock) LogOut(userId string) error {
	mockData := tr.Called(userId)
	//----------------------------------
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		return mockFunc(userId)
	}
	//----------------------------------
	if err, ok := mockData.Error(0).(error); ok {
		return err
	}
	//----------------------------------
	return nil
}

func (tr *UserServiceMock) VerifyAction(rawToken string) (error, string) {
	mockData := tr.Called(rawToken)
	//----------------------------------
	var res1 error
	if mockFunc, ok := mockData.Get(0).(func(string) error); ok {
		res1 = mockFunc(rawToken)
	} else {
		res1 = mockData.Error(0)
	}
	//----------------------------------
	var res2 string
	if mockFunc, ok := mockData.Get(1).(func(string) string); ok {
		res2 = mockFunc(rawToken)
	} else {
		res2 = mockData.String(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *UserServiceMock) VerifyResetPassword(newPass, re_newPass, token string) (string, error) {
	mockData := tr.Called(newPass, re_newPass, token)
	//----------------------------------
	var res1 string
	if mockFunc, ok := mockData.Get(0).(func(string, string, string) string); ok {
		res1 = mockFunc(newPass, re_newPass, token)
	} else {
		res1 = mockData.String(0)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(1).(func(string, string, string) error); ok {
		res2 = mockFunc(newPass, re_newPass, token)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}

func (tr *UserServiceMock) RecoverAccountByCustomer(email string) (string, error) {
	mockData := tr.Called(email)
	//----------------------------------
	var res1 string
	if mockFunc, ok := mockData.Get(0).(func(string) string); ok {
		res1 = mockFunc(email)
	} else {
		res1 = mockData.String(0)
	}
	//----------------------------------
	var res2 error
	if mockFunc, ok := mockData.Get(1).(func(string) error); ok {
		res2 = mockFunc(email)
	} else {
		res2 = mockData.Error(1)
	}
	//----------------------------------
	return res1, res2
}
