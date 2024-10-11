package services

import (
	"log"
	"testing"
	"v3/dbo"
	"v3/mocks"
	"v3/mocks/repos"
	"v3/mocks/samples"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsersResultOfLength3WithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetAllUsers").Return(samples.GetUsersMockData(), nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetUsersMockData()), len(*res))
}

func TestGetUsersByRoleInputExistedRoleResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUsersByRole", mock.AnythingOfType("string")).Return(samples.GetMockUsersByExistedRole(), nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetUsersByRole(mocks.Existed)
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetMockUsersByExistedRole()), len(*res))
}

func TestGetUsersByRoleInputNotExistedRoleResultOfEmptyWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *[]dbo.User
	userRepoMock.On("GetUsersByRole", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetUsersByRole(mocks.NotExisted)
	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestGetUsersByStatusInputValidStatusResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUsersByStatus", mocks.PositiveStatus).Return(samples.GetExistedMockUserByStatus(mocks.PositiveStatus), nil)
	res, err := testTrg.userRepo.GetUsersByStatus(mocks.PositiveStatus)
	assert.NoError(t, err)
	assert.True(t, len(*res) > 0)
}

func TestGetUsersByStatusInputNegativeStatusResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUsersByStatus", mocks.NegativeStatus).Return(samples.GetExistedMockUserByStatus(mocks.NegativeStatus), nil)
	res, err := testTrg.userRepo.GetUsersByStatus(mocks.NegativeStatus)
	log.Print(res)
	assert.NoError(t, err)
	assert.True(t, len(*res) > 0)
}

func TestGetUserByIdInputExistedUserIdResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(samples.GetExistedMockUser(), nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetUserById(mocks.Existed)
	assert.NoError(t, err)
	assert.Equal(t, res.UserId, mocks.Existed)
}

func TestGetUserByIdInputNotExistedUserIdResultOfEmptyWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetUserById(mocks.NotExisted)
	assert.NoError(t, err)
	assert.Nil(t, res)
}

func TestGetUserByEmailInputExistedEmailResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetExistedMockUserByExistedEmail(), nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetUserByEmail(mocks.Existed)
	assert.NoError(t, err)
	assert.Equal(t, res.Email, mocks.Existed)
}

func TestGetUserByEmailInputNotExistedEmailResultOfEmptyWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	res, err := testTrg.userRepo.GetUserByEmail(mocks.Existed)
	assert.NoError(t, err)
	assert.Nil(t, res)
}

func TestAddUserInputValidFieldsResultOfSuccessWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err := testTrg.userRepo.AddUser(*samples.GetExistedMockUser())
	assert.NoError(t, err)
}

func TestUpdateUserInputValidFieldsResultOfSuccessWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var errorResponse error
	userRepoMock.On("UpdateUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err := testTrg.userRepo.UpdateUser(*samples.GetExistedMockUser())
	assert.NoError(t, err)
}

func TestChangeUserStatusInputOfPositiveStatusResultOfSuccessWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var errorResponse error
	userRepoMock.On("ChangeUserStatus", mock.AnythingOfType("bool"), mock.AnythingOfType("string")).Return(errorResponse)
	//----------------------------------
	err := testTrg.userRepo.ChangeUserStatus(mocks.PositiveStatus, mocks.Existed)
	assert.NoError(t, err)
}

func TestChangeUserStatusInputOfNegativeStatusResultOfSuccessWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var errorResponse error
	userRepoMock.On("ChangeUserStatus", mock.AnythingOfType("bool"), mock.AnythingOfType("string")).Return(errorResponse)
	//----------------------------------
	err := testTrg.userRepo.ChangeUserStatus(mocks.NegativeStatus, mocks.Existed)
	assert.NoError(t, err)
}
