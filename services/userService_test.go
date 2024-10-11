package services

import (
	"os"
	"strings"
	"testing"
	"v3/constants/notis"
	"v3/dbo"
	"v3/mocks"
	"v3/mocks/repos"
	"v3/mocks/samples"
	"v3/mocks/services"
	"v3/spModels"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsersServiceResultOfLength3WithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetAllUsers").Return(samples.GetUsersMockData(), nil)
	//----------------------------------
	res, err := testTrg.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetUsersMockData()), len(*res))
}

func TestGetUsersByRoleServiceInputExistedRoleResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUsersByRole", mock.AnythingOfType("string")).Return(samples.GetMockUsersByExistedRole(), nil)
	//----------------------------------
	res, err := testTrg.GetUsersByRole(mocks.Existed)
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetMockUsersByExistedRole()), len(*res))
}

func TestGetUsersByRoleServiceInputNotExistedRoleResultOfEmptyWithErrorMessageAsRoleNotFoundReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var errorResponse error
	var data *[]dbo.User
	userRepoMock.On("GetUsersByRole", mock.AnythingOfType("string")).Return(data, errorResponse)
	//----------------------------------
	res, err := testTrg.GetUsersByRole(mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.UndefinedRoleWarnMsg)
	assert.Nil(t, res)
}

func TestGetUsersByRoleServiceInputEmptyResultOfDataWithFullUsersOfAllRolesWithLengthAs3WithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetAllUsers").Return(samples.GetUsersMockData(), nil)
	//----------------------------------
	userRepoMock.On("GetUsersByRole", mock.AnythingOfType("string")).Return(samples.GetUsersMockData(), nil)
	//----------------------------------
	res, err := testTrg.GetUsersByRole(mocks.EmptyString)
	assert.NoError(t, err)
	assert.Equal(t, len(*res), 3)
}

func TestGetUsersByStatusInputValidPositiveStatusResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUsersByStatus", mock.AnythingOfType("bool")).Return(samples.GetExistedMockUserByStatus(mocks.PositiveStatus), nil)
	//----------------------------------
	res, err := testTrg.GetUsersByStatus(mocks.RawPositiveStatus)
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetExistedMockUserByStatus(mocks.PositiveStatus)), len(*res))
}

func TestGetUsersByStatusInputValidNegativeStatusResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUsersByStatus", mock.AnythingOfType("bool")).Return(samples.GetExistedMockUserByStatus(mocks.NegativeStatus), nil)
	//----------------------------------
	res, err := testTrg.GetUsersByStatus(mocks.RawNegativeStatus)
	assert.NoError(t, err)
	assert.Equal(t, len(*samples.GetExistedMockUserByStatus(mocks.NegativeStatus)), len(*res))
}

func TestGetUsersByStatusInputInvalidStatusResultOfEmptyWithErrorAsInvalidStatusReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *[]dbo.User
	var errorResponse error
	userRepoMock.On("GetUsersByStatus", mock.AnythingOfType("bool")).Return(data, errorResponse)
	//----------------------------------
	res, err := testTrg.GetUsersByStatus(mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.InvalidStatusWarnMsg)
	assert.Nil(t, res)
}

func TestGetUserByIdServiceInputExistedUserIdResultOfDataWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(samples.GetExistedMockUser(), nil)
	//----------------------------------
	res, err := testTrg.GetUserById(mocks.Existed)
	assert.NoError(t, err)
	assert.Equal(t, res.UserId, mocks.Existed)
}

func TestGetUserByIdServiceInputEmptyResultOfEmptyWithErrorAsInvalidDataReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	var errorResponse error
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(data, errorResponse)
	//----------------------------------
	res, err := testTrg.GetUserById(mocks.EmptyString)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.GenericsErrorWarnMsg)
	assert.Nil(t, res)
}

func TestGetUserByIdServiceInputNotExistedUserIdResultOfEmptyWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	res, err := testTrg.GetUserById(mocks.NotExisted)
	assert.NoError(t, err)
	assert.Nil(t, res)
}

func TestAddUserCaseGuestDoInputEmptyEmailResultOfErrorAsEmailEmptyReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Empty Email"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	err, _ := testTrg.AddUser(model, mocks.EmptyString)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.EmailEmptyWarnMsg)
}

func TestAddUserCaseGuestDoInputRegisteredEmailResultOfErrorAsEmailRegisteredReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Email registered"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	userRepoMock.On("GetUserByEmail", model.Email).Return(samples.GetExistedMockUser(), nil)
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err, _ := testTrg.AddUser(model, mocks.EmptyString)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.EmailRegisteredWarnMsg)
}

func TestAddUserCaseGuestDoInputEmptyPasswordResultOfErrorAsPasswordEmptyReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Empty password"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	err, _ := testTrg.AddUser(model, mocks.EmptyString)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.PasswordEmptyWarnMsg)
}

func TestAddUserCaseGuestDoInputNotSecurePasswordResultOfErrorAsPasswordNotSecureReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Password not secure"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserByEmail", strings.ToLower(model.Email)).Return(data, nil)
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err, _ := testTrg.AddUser(model, mocks.EmptyString)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.PasswordNotSecureWarnMsg)
}

func TestAddUserCaseGuestDoInputValidInformationResultOfSuccessWithNoErrorAndMessageAsCheckMailToConfirmReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Valid"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserByEmail", strings.ToLower(model.Email)).Return(data, nil)
	//----------------------------------
	userRepoMock.On("GetAllUsers").Return(samples.GetUsersMockData(), nil)
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err, res := testTrg.AddUser(model, mocks.EmptyString)
	assert.NoError(t, err)
	assert.Equal(t, res, notis.RegistrationAccountMsg)
}

func TestAddUserStaffDoInputNotExistedActorResultOfErrorWithNotDefinedReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	var errorResponse error
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err, _ := testTrg.AddUser(spModels.SignUpModel{}, mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), notis.UndefinedUserWarnMsg)
}

func TestAddUserStaffDoProvideNotExistedRoleResultOfErrorAsUndefinedRoleReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
		roleRepo: &roleRepoMock,
	}
	var errorResponse error
	//----------------------------------
	var actor *dbo.User = samples.GetMockStandardUser(mocks.Roles["Staff"])
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(actor, nil)
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Invalid role"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserByEmail", strings.ToLower(model.Email)).Return(data, nil)
	//----------------------------------
	roleRepoMock.On("GetAllRoles").Return(samples.GetMockStandardRoles(), nil)
	//----------------------------------
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err, _ := testTrg.AddUser(model, (*actor).UserId)
	assert.Error(t, err)
	assert.Equal(t, notis.UndefinedRoleWarnMsg, err.Error())
}

func TestAddUserStaffDoProvideAdminAccountResultOfErrorAStaffEditAdminReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
		roleRepo: &roleRepoMock,
	}
	var errorResponse error
	//----------------------------------
	var actor *dbo.User = samples.GetMockStandardUser(mocks.Roles["Staff"])
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(actor, nil)
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Staff Provides Admin"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	var data *dbo.User = nil
	userRepoMock.On("GetUserByEmail", strings.ToLower(model.Email)).Return(data, nil)
	//----------------------------------
	roleRepoMock.On("GetAllRoles").Return(samples.GetMockStandardRoles(), nil)
	//----------------------------------
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(errorResponse)
	//----------------------------------
	err, _ := testTrg.AddUser(model, actor.UserId)
	assert.Error(t, err)
	assert.Equal(t, notis.StaffEditAdminWarnMsg, err.Error())
}

func TestAddUserStaffDoProvideValidAccountResultOfSuccessWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	roleRepoMock := repos.RoleRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
		roleRepo: &roleRepoMock,
	}
	//----------------------------------
	var actor *dbo.User = samples.GetMockStandardUser(mocks.Roles["Staff"])
	userRepoMock.On("GetUserById", mock.AnythingOfType("string")).Return(actor, nil)
	//----------------------------------
	var caseType int = mocks.GetSignUpModelCases()["Valid Provide"]
	var model spModels.SignUpModel = samples.GetSignUpModelBasedOnCase(caseType)
	//----------------------------------
	var data *dbo.User = nil
	userRepoMock.On("GetUserByEmail", strings.ToLower(model.Email)).Return(data, nil)
	//----------------------------------
	roleRepoMock.On("GetAllRoles").Return(samples.GetMockStandardRoles(), nil)
	//----------------------------------
	userRepoMock.On("GetAllUsers").Return(samples.GetUsersMockData(), nil)
	//----------------------------------
	userRepoMock.On("AddUser", mock.AnythingOfType("dbo.User")).Return(nil)
	//----------------------------------
	err, res := testTrg.AddUser(model, actor.UserId)
	assert.NoError(t, err)
	assert.Equal(t, "Success", res)
}

func TestLoginInputNotExistEmailResultOfErrorAsWrongCredentialsReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	var data *dbo.User
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(data, nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.NotExisted, mocks.SecurePassword)
	assert.Error(t, err)
	assert.Equal(t, notis.WrongCredentialsWarnMsg, err.Error())
}

func TestLoginInputWrongPasswordResultOfErrorAsWrongCredentialsReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Active"), nil)
	//----------------------------------
	userRepoMock.On("UpdateUser", mock.AnythingOfType("dbo.User")).Return(nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, notis.WrongCredentialsWarnMsg, err.Error())
}

func TestLoginInputWrongPasswordInStateOfTemporarilyLockedResultOfErrorAsWrongCredentialsReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Customer locked"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, notis.LockWarnMsg, err.Error())
}

func TestLoginStaffAccountInputWrongPasswordInStateOfTemporarilyLockedResultOfErrorAsTemporarilyLockedReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Staff"], "Staff locked"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, notis.LockWarnMsg, err.Error())
}

func TestLoginInputWrongPasswordInStateOfSelfLockedResultOfErrorAsWrongCredentialsReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Self locked"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, notis.WrongCredentialsWarnMsg, err.Error())
}

func TestLoginInputCorrectCredentialsInStateOfSelfLockedResultOfErrorAsSelfLockedReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Self locked"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.SecurePassword)
	assert.Error(t, err)
	assert.Equal(t, notis.InactiveAccountMsg, err.Error())
}

func TestLoginInputWrongPasswordInStateOfNotActivateAccountResultOfErrorAsWrongCredentialsReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Not activate"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.NotExisted)
	assert.Error(t, err)
	assert.Equal(t, notis.WrongCredentialsWarnMsg, err.Error())
}

func TestLoginInputCorrectCredentialsInStateOfNotActivateAccountResultOfMessageAsMindToActivateAccountWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Not activate"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, msg, err := testTrg.Login(mocks.ValidEmail, mocks.SecurePassword)
	assert.NoError(t, err)
	assert.Equal(t, notis.ActivateAccountMsg, msg)
}

func TestLoginInputCorrectCredentialsResultOfFlagAsResetWithUrlRedirectsToResetPasswordPageWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Reset Password"), nil)
	//----------------------------------
	userRepoMock.On("UpdateUser", mock.AnythingOfType("dbo.User")).Return(nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	res1, res2, err := testTrg.Login(mocks.ValidEmail, mocks.SecurePassword)
	assert.NoError(t, err)
	assert.Equal(t, "Reset", res1)
	assert.True(t, strings.Contains(res2, "token"))
}

func TestLoginInputCorrectCredentialsInStateOfBannedResultOfErrorAsBannedReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Banned"), nil)
	//----------------------------------
	var errorResponse error
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", errorResponse)
	//----------------------------------
	_, _, err := testTrg.Login(mocks.ValidEmail, mocks.SecurePassword)
	assert.Error(t, err)
	assert.Equal(t, notis.AccountBanWarnMsg, err.Error())
}

func TestLoginInputCorrectCredentialsInNormalStateResultOf2TokensAsReturnWithNoErrorReturnsWell(t *testing.T) {
	userRepoMock := repos.UserRepoMock{}
	//----------------------------------
	testTrg := userService{
		userRepo: &userRepoMock,
	}
	//----------------------------------
	userRepoMock.On("GetUserByEmail", mock.AnythingOfType("string")).Return(samples.GetStandardMockAccountForLogin(mocks.Roles["Customer"], "Active"), nil)
	//----------------------------------
	userRepoMock.On("UpdateUser", mock.AnythingOfType("dbo.User")).Return(nil)
	//----------------------------------
	userServiceMock := services.UserServiceMock{}
	userServiceMock.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", "", nil)
	//----------------------------------
	res1, res2, err := testTrg.Login(mocks.ValidEmail, mocks.SecurePassword)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(res1, os.Getenv("SECRET_KEY")))
	assert.True(t, strings.Contains(res2, os.Getenv("SECRET_KEY")))
}
