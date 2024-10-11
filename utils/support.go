package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"v3/constants/notis"
	"v3/dbo"
	irepositories "v3/interfaces/iRepositories"
	"v3/spModels"
)

func IsStatusValid(rawStatus string) (bool, error) {
	status, err := strconv.ParseBool(rawStatus)
	if err != nil {
		return false, errors.New(notis.InvalidStatusWarnMsg)
	}
	return status, nil
}

func IsPasswordSecure(password string) bool {
	upperCase := "(?i)[A-Z]"                // At least one uppercase letter
	lowerCase := "[a-z]"                    // At least one lowercase letter
	digit := "[0-9]"                        // At least one digit
	specialChar := `[!@#$%^&*()_+{}|:"<>?]` // At least one special character
	minLength := ".{8,}"                    // Minimum length of 8 characters

	// Compile regular expressions
	upRgx, _ := regexp.Compile(upperCase)
	lowRgx, _ := regexp.Compile(lowerCase)
	digRgx, _ := regexp.Compile(digit)
	speRgx, _ := regexp.Compile(specialChar)
	lenRgx, _ := regexp.Compile(minLength)

	return lenRgx.MatchString(password) &&
		upRgx.MatchString(password) &&
		lowRgx.MatchString(password) &&
		digRgx.MatchString(password) &&
		speRgx.MatchString(password)
}

func VerifyUpdateAuth(user spModels.UserNormalModel, actor, origin dbo.User) error {
	// 2 different accounts
	if user.UserId != actor.UserId {
		// origin, err := repositories.GetUserById(user.UserId) // primitive account before updated
		// if err != nil {
		// 	return nil
		// }
		//-----------------------------------------------
		var roles = FetchRoles(nil)
		if actor.RoleId == roles["Admin"] { // Actor is admin
			if origin.RoleId == roles["Admin"] { // Account update is an admin account
				return errors.New("Can't edit other admins") // Admin can't edit other admins
			}
			//-----------------------------------------------
		} else if actor.RoleId == roles["Customer"] { // Actor is user
			return errors.New(notis.UserEditOtherWarnMsg)
		} else if actor.RoleId == roles["Staff"] { // Actor is staff
			if user.RoleId != "" && user.RoleId != origin.RoleId { // Account's role changed
				return errors.New(notis.StaffEditRoleWarnMsg)
			}
			//-----------------------------------------------
			if origin.RoleId != roles["Customer"] { // Account changed is admin/staff
				return errors.New(notis.EditAdminStaffDataWarnMsg) // Staff can't edit other staffs, admins
			} else { // Account changed is user
				if user.Email != "" && origin.Email != strings.TrimSpace(strings.ToLower(user.Email)) { // Changed user email
					return errors.New(notis.LowerRoleEditOtherEmailWarnMsg)
				}
			}
		} // Other roles in future
	} else { // same account
		if user.RoleId != "" && user.RoleId != origin.RoleId { // Change their own role -> Deny
			return errors.New(notis.EditOwnRoleWarnMsg)
		}
	}
	//-----------------------------------------------
	return nil
}

func IsEmailExisted(email string, repo *irepositories.IUserRepo) (bool, error) {
	user, err := (*repo).GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}
	return false, nil // Email not registered -> Approve
}

func isAccountActivated(acc *dbo.User) bool {
	primitiveTime := GetPrimitiveTime()
	// if !acc.ActiveStatus && acc.FailAccess == 0 && acc.ActionPeriod == nil && acc.ActionToken == nil && acc.LastFail == nil || *acc.LastFail == primitiveTime {
	// 	return false
	// }
	if !acc.ActiveStatus {
		if acc.FailAccess == 0 {
			if acc.ActionPeriod == nil {
				if acc.ActionToken == nil {
					if acc.LastFail == nil || *acc.LastFail == primitiveTime {
						return false
					}
				}
			}
		}
	}

	return true
}

func GenerateCallBackUrl(sq []string) string {
	if len(sq) < 1 {
		log.Print(notis.SupportMsg + "GenerateCallBackUrl - " + "Missing data or empty string-slice parameter.")
		return ""
	}
	var res string = ""
	for i, v := range sq {
		res += v
		if i < len(sq)-1 && i > 0 {
			res += ":"
		}
	}
	return res
}

func CaseBodyForVerifyActionType(u *dbo.User, email string, actionType string, cases []string, repo *irepositories.IUserRepo) error {
	// email: represents the new email which user requests to update his/her account
	// if case belongs to Update-new-email type, parameter s will be that new email to activate to user account
	switch actionType {
	case cases[0]: // Activate account type
		u.ActiveStatus = true
		if u.FailAccess == 6 { // Recover case
			u.FailAccess = 0
		}
		if u.LastFail == nil { // Recover case
			tmpTime := GetPrimitiveTime()
			u.LastFail = &tmpTime
		}
	case cases[1]: // Reset password -> Redirect user to reset password page after confirming mail (No need to execute update user data)
	case cases[2]: // Update new email to account
		if email == "" {
			return errors.New(notis.GenericsErrorWarnMsg)
		}
		u.Email = email
		// Other cases in future
	default:
	}
	u.ActionPeriod = nil
	u.ActionToken = nil
	if err := (*repo).UpdateUser(*u); err != nil {
		return err
	}
	//-----------------------------------------------
	return nil
}

func FetchRoles(service irepositories.IRoleRepo) map[string]string {
	list, err := service.GetAllRoles()
	if err != nil || service == nil {
		return map[string]string{
			"Admin":    "R001",
			"Staff":    "R002",
			"Customer": "R003",
		}
	}
	//-----------------------------------------
	res := make(map[string]string)
	for _, v := range *list {
		res[v.RoleName] = v.RoleId
	}
	//-----------------------------------------
	return res
}
