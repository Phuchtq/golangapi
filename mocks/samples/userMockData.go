package samples

import (
	"strings"
	"time"
	"v3/dbo"
	"v3/mocks"
	"v3/spModels"
	"v3/utils"
)

func GetUsersMockData() *[]dbo.User {
	primitiveTime := utils.GetPrimitiveTime()
	curTime := time.Now().UTC()
	token := mocks.Token
	return &[]dbo.User{
		{
			UserId:       "1",
			RoleId:       "1",
			Email:        "1",
			Pasword:      "1",
			ActiveStatus: true,
			FailAccess:   1,
			LastFail:     &time.Time{},
		},
		//----------------------------------
		{
			UserId:       "2",
			RoleId:       "2",
			Email:        "2",
			Pasword:      "2",
			ActiveStatus: false,
			FailAccess:   6,
			LastFail:     &time.Time{},
		},
		//----------------------------------
		{
			UserId:       "3",
			RoleId:       "3",
			Email:        "3",
			Pasword:      "3",
			ActiveStatus: true,
			FailAccess:   0,
			AccessToken:  &token,
			RefreshToken: &token,
			ActionToken:  &token,
			ActionPeriod: &curTime,
			LastFail:     &primitiveTime,
		},
	}
}

func GetExistedMockUser() *dbo.User {
	for _, v := range *GetUsersMockData() {
		if mocks.Existed == v.UserId {
			return &v
		}
	}
	//--------------------------------------
	return &dbo.User{ // In case test data has some problems
		UserId: mocks.Existed,
	}
}

func GetExistedMockUserByExistedEmail() *dbo.User {
	for _, v := range *GetUsersMockData() {
		if mocks.Existed == v.Email {
			return &v
		}
	}
	//--------------------------------------
	return &dbo.User{ // In case test data has some problems
		UserId:       mocks.Existed,
		Email:        mocks.Existed,
		ActiveStatus: true,
	}
}

func GetExistedMockUserByStatus(status bool) *[]dbo.User {
	var res []dbo.User
	for _, v := range *GetUsersMockData() {
		if v.ActiveStatus == status {
			res = append(res, v)
		}
	}
	//--------------------------------------
	if len(res) > 0 {
		return &res
	}
	//--------------------------------------
	return &[]dbo.User{ // In case test data has some problems
		{
			RoleId:       "1",
			UserId:       "1",
			ActiveStatus: status,
		},
	}
}

func GetMockUsersByExistedRole() *[]dbo.User {
	var res []dbo.User
	for _, v := range *GetUsersMockData() {
		if v.RoleId == mocks.Existed {
			res = append(res, v)
		}
	}
	//--------------------------------------
	if len(res) > 0 {
		return &res
	}
	//--------------------------------------
	return &[]dbo.User{ // In case test data has some problems
		{
			UserId:       mocks.Existed,
			RoleId:       mocks.Existed,
			ActiveStatus: true,
		},
	}
}

func GetMockStandardUser(role string) *dbo.User {
	return &dbo.User{
		UserId:       mocks.Existed,
		RoleId:       role,
		ActiveStatus: true,
	}
}

func supportGenerateStandardAccount(userState string) (fails int, lastFail *time.Time, activeStatus bool) {
	activeStatus = false
	tmpCurtime := time.Now()
	lastFail = &tmpCurtime
	//--------------------------------
	switch userState {
	case mocks.AccountStates[0]:
		fails = 6
	//--------------------------------
	case mocks.AccountStates[1]:
		fails = 4
	//--------------------------------
	case mocks.AccountStates[2]:
		fails = 3
	//--------------------------------
	case mocks.AccountStates[3]:
		fails = 0
		lastFail = nil
	//--------------------------------
	case mocks.AccountStates[4]:
		fails = 0
		tmpPrimitive := utils.GetPrimitiveTime()
		lastFail = &tmpPrimitive
	//--------------------------------
	default:
		fails = 0
		activeStatus = true
		tmpPrimitive := utils.GetPrimitiveTime()
		lastFail = &tmpPrimitive
	}
	//--------------------------------
	return fails, lastFail, activeStatus
}

func GetStandardMockAccountForLogin(role, userState string) *dbo.User {
	password, _ := utils.ToHashString(mocks.SecurePassword)
	//---------------------------------------------
	var isResetPw *bool = nil
	if strings.Contains(userState, "Reset") {
		positiveStatus := true
		isResetPw = &positiveStatus
	}
	fails, lastFail, activeStatus := supportGenerateStandardAccount(userState)
	return &dbo.User{
		UserId:          mocks.Existed,
		RoleId:          role,
		Email:           mocks.ValidEmail,
		Pasword:         password,
		FailAccess:      fails,
		LastFail:        lastFail,
		ActiveStatus:    activeStatus,
		IsHaveToResetPw: isResetPw,
	}
}

func GetSignUpModelBasedOnCase(caseType int) spModels.SignUpModel {
	var res spModels.SignUpModel
	//--------------------------------------
	cases := mocks.GetSignUpModelCases()
	keys := mocks.SignUpModelKeyCases
	//--------------------------------------
	switch caseType {
	case cases[keys[0]]:
		res.Password = mocks.SecurePassword
		//--------------------------------------
	case cases[keys[1]]:
		res.Email = mocks.ValidEmail
		//--------------------------------------
	case cases[keys[2]]:
		res.Email = mocks.Existed
		res.Password = mocks.SecurePassword
		//--------------------------------------
	case cases[keys[3]]:
		res.Email = mocks.ValidEmail
		res.Password = mocks.InvalidPassword
		//--------------------------------------
	case cases[keys[4]]:
		res.Email = mocks.ValidEmail
		res.Password = mocks.SecurePassword
		res.RoleId = mocks.Roles["Admin"]
		//--------------------------------------
	case cases[keys[5]]:
		res.Email = mocks.ValidEmail
		res.Password = mocks.SecurePassword
		res.RoleId = mocks.NotExisted
		//--------------------------------------
	case cases[keys[6]]:
		res.Email = mocks.ValidEmail
		res.Password = mocks.SecurePassword
		//--------------------------------------
	case cases[keys[7]]:
		res.Email = mocks.ValidEmail
		res.Password = mocks.SecurePassword
		res.RoleId = mocks.Roles["Customer"]
	}
	//--------------------------------------
	return res
}
