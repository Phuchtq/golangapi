package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"v3/constants/notis"
	"v3/dbo"
	irepositories "v3/interfaces/iRepositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	staffLockPeriod  int = 4
	norAccLockPeriod int = 3
)

var customerFailRange = []int{3, 4, 5}

func getCustomerLockPeriods() map[int]time.Duration {
	res := make(map[int]time.Duration)
	//-----------------------------------
	var primitiveDuration time.Duration = 15 * time.Minute
	//-----------------------------------
	for _, v := range customerFailRange {
		res[v] = primitiveDuration
		primitiveDuration *= 2
	}
	//-----------------------------------
	return res
}

func GenerateTokens(email string, userId string, role string) (string, string, error) {
	var bytes = []byte(os.Getenv("SECRET_KEY"))
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"role":   role,
		"expire": time.Now().Add(normalActionDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		log.Print(notis.SecurityMsg + "GenerateTokens - " + fmt.Sprint(err))
		return "", "", errors.New(notis.InternalErr)
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"email":  email,
		"role":   role,
		"expire": time.Now().Add(refreshTokenDuration).Unix(),
	}).SignedString(bytes)
	if err != nil {
		log.Print(notis.SecurityMsg + "GenerateTokens - " + fmt.Sprint(err))
		return "", "", errors.New(notis.InternalErr)
	}

	return accessToken, refreshToken, nil
}

func ToHashString(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 10)
	if err != nil {
		log.Print(notis.SecurityMsg + "ToHashString - " + fmt.Sprint(err))
		return "", errors.New(notis.InternalErr)
	}
	return string(bytes), nil
}

func isLoginPasswordMatched(password string, inputPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(inputPass))
	if err != nil {
		return false
	}
	return true
}

func VerifyToken(s string) (string, string, time.Time, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		log.Print("Error at VerifyToken - ", err)
		return "", "", GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	}

	var userId string = ""
	var role string = ""
	var exp time.Time
	if claims, ok := token.Claims.(jwt.MapClaims); !ok || token.Valid {
		return "", "", GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
	} else {
		// if email, ok := claims["email"].(string); email == "" || !ok {
		// 	return "", "", GetPrimitiveTime(), errors.New(errorMsg)
		// }

		if rawRole, ok := claims["role"].(string); rawRole == "" || !ok {
			return "", "", GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
		} else {
			role = rawRole
		}

		if id, ok := claims["userId"].(string); id == "" || !ok {
			return "", "", GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
		} else {
			userId = id
		}

		if expPeriod, ok := claims["exp"].(time.Time); expPeriod == GetPrimitiveTime() || !ok {
			return "", "", GetPrimitiveTime(), errors.New(notis.GenericsErrorWarnMsg)
		} else {
			exp = expPeriod
		}

	}

	return userId, role, exp, nil
}

func VerifyLogin(password string, acc *dbo.User, repo *irepositories.IUserRepo) error {
	// acc: account that has been found with the inputted email
	isCorrectCredentials := isLoginPasswordMatched(acc.Pasword, password)

	isActivated := isAccountActivated(acc)
	if !isActivated {
		if !isCorrectCredentials { // Incorrect password
			return errors.New(notis.WrongCredentialsWarnMsg) // Return message, no counts for fail attempts
		}

		if acc.LastFail != nil && *acc.LastFail == GetPrimitiveTime() { // Still not activate account when register
			return errors.New("Activate")
		}

		return errors.New(notis.InactiveAccountMsg) // Lock account as personal purpose
	}
	//-----------------------------------
	if acc.RoleId != FetchRoles(nil)["Customer"] {
		if !isStaffLockExpired(acc) {
			return errors.New(notis.LockWarnMsg)
		}
	} else {
		if err := caseBodyForLogin(acc.FailAccess, *acc.LastFail, isCorrectCredentials); err != nil {
			return err
		}
	}
	//-----------------------------------
	if !isCorrectCredentials {
		if err := updateLoginFailCase(*acc, repo); err != nil {
			return err
		}
		//-----------------------------------
		return errors.New(notis.WrongCredentialsWarnMsg)
	}
	//-----------------------------------
	if acc.IsHaveToResetPw != nil && *acc.IsHaveToResetPw {
		if acc.FailAccess > 0 {
			acc.FailAccess = 0
		}
		//-----------------------------------
		if !acc.ActiveStatus {
			acc.ActiveStatus = true
		}
		//-----------------------------------
		primitiveTime := GetPrimitiveTime()
		if acc.LastFail != &primitiveTime {
			acc.LastFail = &primitiveTime
		}
		// Reset these previous status back to its normal state if user pass correct credentials before send reset flag and redirect them to reset-password page
		//-----------------------------------
		return errors.New("Reset") // As a flag to check if an account is forced to change password as previous action such as:
	} // Recover account by user or staff, ...
	//-----------------------------------
	if acc.FailAccess > 0 {
		return recoverUserStateLogin(*acc, repo)
	}
	//-----------------------------------
	return nil // If previous attempt to this account had correct credentials -> No need to to reset the state back to its origin
}

func isStaffLockExpired(acc *dbo.User) bool {
	var lockDuration time.Duration = staffLockDuration
	if acc.RoleId == FetchRoles(nil)["Admin"] {
		lockDuration = adminLockDuration
	}
	//-----------------------------------
	if acc.FailAccess%staffLockPeriod == 0 && acc.FailAccess > 0 {
		if !isActionExpired(*acc.LastFail, lockDuration) {
			return false
		}
	}
	//-----------------------------------
	return true
}

func caseBodyForLogin(fail int, last time.Time, flag bool) error {
	// flag: a boolean variable representing the state of login with 2 cases:
	// correct username and password with True value and False with just only username
	// last: the last time which user accessed to this account with wrong credentials
	// fail: number of fail attempts to this account

	if flag {
		if fail == 6 { // banned period
			return errors.New(notis.AccountBanWarnMsg) // Display message ban for correct credentials
		}
	}
	//-----------------------------------
	if fail == 6 { // As have gone through case flag, if fail equals to 6 -> displays message of login with wrong credentials
		return errors.New(notis.WrongCredentialsWarnMsg)
	}
	//-----------------------------------
	for failAttempts, lockDuration := range getCustomerLockPeriods() {
		// Generate a case body based on keys-values of lockPeriods by iterating through it
		if fail == failAttempts { // Number of actual fails equals to one of those in lockPeriods
			if !isActionExpired(last, lockDuration) { // If from the fail period, it still not passed the set lock duration for this account
				return errors.New(notis.LockWarnMsg) // Return lock message
			}
		}
	}
	//-----------------------------------
	return nil
}

func updateLoginFailCase(acc dbo.User, repo *irepositories.IUserRepo) error {
	//if acc.FailAccess == 0 && acc.LastFail
	acc.FailAccess += 1
	tmpCur := time.Now().UTC()
	acc.LastFail = &tmpCur
	//-----------------------------------
	if acc.RoleId == "R003" {
		if acc.FailAccess >= norAccLockPeriod { // Temporarily set false for normal accounts with fail attempts greater than 2
			acc.ActiveStatus = false
		}
	} else {
		if acc.FailAccess > 0 && acc.FailAccess%staffLockPeriod == 0 { // Admin case
			acc.ActiveStatus = false
		}
	}
	//-----------------------------------
	return (*repo).UpdateUser(acc)
}

func recoverUserStateLogin(u dbo.User, repo *irepositories.IUserRepo) error {
	tmpTime := GetPrimitiveTime()
	u.LastFail = &tmpTime
	u.FailAccess = 0
	u.ActiveStatus = true
	if u.ActionToken != nil { // Check if a request is out dated such as: reset password request, update new email, ...
		if isActionExpired(*u.ActionPeriod, normalActionDuration) { // Expired
			u.ActionPeriod = nil
			u.ActionToken = nil
		}
	}
	return (*repo).UpdateUser(u)
}

func VerifyActionToken(token string, user *dbo.User, repo *irepositories.IUserRepo) error {
	userId, roleId, _, err := VerifyToken(token)
	if err != nil {
		return err
	}
	//-----------------------------------
	user, err = (*repo).GetUserById(userId)
	if err != nil {
		return err
	}
	//-----------------------------------
	if user == nil {
		return errors.New(notis.GenericsErrorWarnMsg)
	}
	//-----------------------------------
	if user.RoleId != roleId {
		return errors.New(notis.GenericsErrorWarnMsg)
	}
	//-----------------------------------
	if user.ActionToken == nil || *user.ActionToken != token { // Fake url with a token not exist in database
		return errors.New(notis.GenericsErrorWarnMsg)
	}
	//-----------------------------------
	if user.ActionPeriod == nil || isActionExpired(*user.ActionPeriod, normalActionDuration) { // Action expired
		return errors.New(notis.ExpirationWarnMsg)
	}
	//-----------------------------------
	if !user.ActiveStatus {
		return errors.New(notis.LockWarnMsg)
	}
	//-----------------------------------
	return nil
}

func VerifyActorAndObject(actorId, userId string, actor, origin *dbo.User, repo *irepositories.IUserRepo) error {
	if actorId == "" || userId == "" {
		return errors.New(notis.AnonymousWarnMsg)
	}
	//-----------------------------------
	actor, err := (*repo).GetUserById(actorId)
	if err != nil {
		return err
	}
	//-----------------------------------
	if actor == nil {
		return errors.New(notis.AnonymousWarnMsg)
	}
	//-----------------------------------
	origin, err = (*repo).GetUserById(userId)
	if err != nil {
		return err
	}
	//-----------------------------------
	if origin == nil {
		return errors.New(notis.AnonymousWarnMsg)
	}
	//-----------------------------------
	return nil
}
