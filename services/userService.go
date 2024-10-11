package services

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"v3/constants/notis"
	"v3/dbo"
	irepositories "v3/interfaces/iRepositories"
	iservices "v3/interfaces/iServices"
	"v3/spModels"
	"v3/templatePath"
	"v3/utils"
)

type userService struct {
	userRepo irepositories.IUserRepo
	roleRepo irepositories.IRoleRepo
}

func InitializeUserService(userRepo irepositories.IUserRepo, roleRepo irepositories.IRoleRepo) iservices.IUserService {
	return &userService{
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

const (
	activateType      string = "1"
	resetPassType     string = "2"
	updateProfileType string = "3"
	LoginPageUrl      string = "Your-login-page-ủrl"
)

func getProcessUrl() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//------------------------------------
	return "http://localhost:" + port + "/VerifyAction?rawToken="
}

func getResetPassUrl() string {
	return "Your reset-pass URL page?token="
}

func (tr *userService) GetAllUsers() (*[]dbo.User, error) {
	return tr.userRepo.GetAllUsers()
}

func (tr *userService) GetUsersByRole(role string) (*[]dbo.User, error) {
	if role == "" {
		return tr.userRepo.GetAllUsers()
	}

	role = strings.TrimSpace(role)
	if res, err := tr.userRepo.GetUsersByRole(role); res == nil && err == nil {
		return nil, errors.New(notis.UndefinedRoleWarnMsg)
	}

	return tr.userRepo.GetUsersByRole(role)
}

func (tr *userService) GetUsersByStatus(rawStatus string) (*[]dbo.User, error) {
	status, err := strconv.ParseBool(rawStatus)
	if err != nil {
		return nil, errors.New(notis.InvalidStatusWarnMsg)
	}
	return tr.userRepo.GetUsersByStatus(status)
}

func (tr *userService) GetUserById(id string) (*dbo.User, error) {
	if id = strings.TrimSpace(id); id == "" {
		return nil, errors.New(notis.GenericsErrorWarnMsg)
	}
	return tr.userRepo.GetUserById(id)
}

func (tr userService) AddUser(u spModels.SignUpModel, actorId string) (error, string) {
	// Fake userId -> stop the process
	var actor *dbo.User
	var err error
	if actorId != "" {
		if actor, err = tr.userRepo.GetUserById(actorId); err != nil {
			return err, ""
		} else if err == nil && actor == nil { // Actor not exist
			return errors.New(notis.UndefinedUserWarnMsg), ""
		}
	}

	if u.Email == "" {
		return errors.New(notis.EmailEmptyWarnMsg), ""
	}
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if isExisted, err := utils.IsEmailExisted(u.Email, &tr.userRepo); err != nil {
		return err, ""
	} else if isExisted {
		return errors.New(notis.EmailRegisteredWarnMsg), ""
	}

	if u.Password == "" {
		return errors.New(notis.PasswordEmptyWarnMsg), ""
	}

	var orgPass string = u.Password
	if isSecure := utils.IsPasswordSecure(u.Password); !isSecure {
		return errors.New(notis.PasswordNotSecureWarnMsg), ""
	}
	if hashPass, err := utils.ToHashString(u.Password); err != nil {
		return err, ""
	} else {
		u.Password = hashPass
	}

	// Guest registers an account
	var roles = utils.FetchRoles(tr.roleRepo)
	if actorId == "" {
		u.RoleId = roles["Customer"] // Customer role
	} else {
		if u.RoleId != "" {
			if isExisted, err := isRoleExisted(u.RoleId, &tr.roleRepo); err != nil {
				return err, ""
			} else if !isExisted {
				return errors.New(notis.UndefinedRoleWarnMsg), ""
			}
		}
		// Actor is a staff
		if actor.RoleId == roles["Staff"] {
			// Staff creates an Admin account
			if u.RoleId == roles["Admin"] {
				return errors.New(notis.StaffEditAdminWarnMsg), ""
			}
		}
	}

	if u.RoleId == "" {
		u.RoleId = roles["Customer"]
	}

	tmpTime := utils.GetPrimitiveTime()
	id, err := generateUserId(&tr.userRepo)
	if err != nil {
		return err, ""
	}
	//---------------------------------------------------//
	// This snippet below used for generating token for confirmation mail
	// If user confirmed and this token matched with the data saved in database, this account will be activated firmly
	tmpToken, _, _ := utils.GenerateTokens(u.Email, id, u.RoleId)
	//---------------------------------------------------
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	//---------------------------------------------------
	tmpCurTime := time.Now().UTC()
	//---------------------------------------------------
	var isHaveToResetPw *bool = nil
	if actorId != "" {
		flag := true
		isHaveToResetPw = &flag
	}
	//---------------------------------------------------
	if err := tr.userRepo.AddUser(dbo.User{
		UserId:          id,
		RoleId:          u.RoleId,
		Email:           u.Email,
		Pasword:         u.Password,
		ActiveStatus:    false,
		FailAccess:      0,
		LastFail:        &tmpTime,
		ActionToken:     &tmpToken,
		ActionPeriod:    &tmpCurTime,
		IsHaveToResetPw: isHaveToResetPw,
	}); err != nil {
		return err, ""
	}
	//---------------------------------------------------
	// Generating callback url
	var url string = utils.GenerateCallBackUrl([]string{
		getProcessUrl(),
		tmpToken,
		id,
		activateType,
	})
	if url == "" {
		return errors.New(notis.InternalErr), ""
	}
	// No problem when processing -> send confirmation mail
	if err := utils.SendMail(templatePath.AccountRegistrationTemplate, notis.RegistrationAccountSubject, spModels.MailBody{
		Email:    u.Email,
		Password: orgPass,
		Url:      url,
	}); err != nil {
		return err, ""
	}
	//---------------------------------------------------
	var msg string = "Success"
	if actorId == "" {
		msg = notis.RegistrationAccountMsg
	}
	//---------------------------------------------------
	return nil, msg
	// Account provided by staff, admin, ... withoud any problems during the process
}

func (tr *userService) Login(email string, password string) (string, string, error) {
	acc, err := tr.userRepo.GetUserByEmail(strings.TrimSpace(email))
	//-----------------------------------------------
	if err != nil { // Internal error like database connection, ...
		return "", "", err
	}
	//-----------------------------------------------
	if acc == nil { // No account found with the inputted email
		return "", "", errors.New(notis.WrongCredentialsWarnMsg)
	}
	//-----------------------------------------------
	if err := utils.VerifyLogin(password, acc, &tr.userRepo); err != nil {
		resetFlag := "Reset"
		activateFlag := "Activate"
		var url string
		var signFlag string = ""
		var res2 string = ""
		//-----------------------------------------------
		if err.Error() == resetFlag { // Have to reset pass case as admin/staff provides account/change customer's password
			if err := prepareForcedResetPass(acc, &url); err != nil { // Generate token and url to redirect user
				return "", "", err
			}
			//-----------------------------------------------
			signFlag = resetFlag
			res2 = url
			//-----------------------------------------------
		} else if err.Error() == activateFlag { // User has registered for this account but have not activated it yet -> generate callback url to process when user
			actionToken, _, err := utils.GenerateTokens(email, acc.UserId, acc.RoleId) // click to the confirmation mail to activate this account again
			if err != nil {
				return "", "", err
			}
			//-----------------------------------------------
			if actionToken == "" {
				return "", "", errors.New(notis.InternalErr)
			}
			//-----------------------------------------------
			url = utils.GenerateCallBackUrl([]string{
				getProcessUrl(),
				actionToken,
				acc.UserId,
				activateType,
			})
			//-----------------------------------------------
			if url == "" {
				return "", "", errors.New(notis.InternalErr)
			}
			//-----------------------------------------------
			if err := utils.SendMail(templatePath.AccountRegistrationTemplate, notis.RegistrationAccountSubject, spModels.MailBody{
				Email: email,
				Url:   url,
			}); err != nil {
				return "", "", err
			}
			//-----------------------------------------------
			res2 = notis.ActivateAccountMsg
			acc.ActionToken = &actionToken
			*acc.ActionPeriod = time.Now().UTC()
		}
		//-----------------------------------------------
		if signFlag != "" && res2 != "" {
			if err := tr.userRepo.UpdateUser(*acc); err != nil {
				return "", "", err
			}
			//-----------------------------------------------
			return signFlag, res2, nil
		}
		//-----------------------------------------------
		return "", "", err
	}
	//-----------------------------------------------
	token, refreshToken, err := utils.GenerateTokens(email, acc.UserId, acc.RoleId) // Generate access and refresh tokens
	if err != nil {
		return "", "", err
	}
	//-----------------------------------------------
	acc.AccessToken = &token
	acc.RefreshToken = &refreshToken
	if err := tr.userRepo.UpdateUser(*acc); err != nil {
		return "", "", err
	}
	//-----------------------------------------------
	return token, refreshToken, nil
}

func (tr *userService) UpdateUser(user spModels.UserNormalModel, actorId string) (string, error) {
	var actor dbo.User
	var origin dbo.User
	err := utils.VerifyActorAndObject(actorId, user.UserId, &actor, &origin, &tr.userRepo)
	if err != nil {
		return "", err
	}
	//--------------------------------------------------------------
	if err := utils.VerifyUpdateAuth(user, actor, origin); err != nil {
		return "", err
	}
	//--------------------------------------------------------------
	if actor.RoleId == utils.FetchRoles(tr.roleRepo)["Customer"] { // Actor is user
		if user.Pasword != "" { // New password inputed
			if !utils.IsPasswordSecure(user.Pasword) { // New password not secure enough
				return "", errors.New(notis.PasswordNotSecureWarnMsg)
			}
		}
	}
	//--------------------------------------------------------------/
	if origin.Pasword, err = utils.ToHashString(user.Pasword); err != nil {
		return "", err
	}
	//--------------------------------------------------------------/
	activeStatus, err := strconv.ParseBool(user.ActiveStatus)
	if err != nil {
		if user.ActiveStatus == "" { // Error as empty status -> take back previous status, continue the process
			activeStatus = origin.ActiveStatus
		} else { // Error as other source -> returns back error
			return "", errors.New(notis.InvalidStatusWarnMsg)
		}
	}
	//--------------------------------------------------------------/
	if activeStatus != origin.ActiveStatus {
		var tmpLastFail *time.Time = nil
		if activeStatus { // The updated status is true
			// This action can only be done by admin for other roles or staff for user
			tmpTime := utils.GetPrimitiveTime()
			tmpLastFail = &tmpTime
		}
		//--------------------------------------------------------------/
		origin.FailAccess = 0
		origin.LastFail = tmpLastFail
		origin.ActiveStatus = activeStatus
	}
	//--------------------------------------------------------------/
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//--------------------------------------------------------------/
	if user.Email != "" {
		if origin.Email != user.Email { // Change email
			isExist, err := utils.IsEmailExisted(user.Email, &tr.userRepo)
			if err != nil {
				return "", err
			}
			//-------------------------------------------------
			if !isExist {
				if user.UserId == actorId { // Actor change their own email
					actionToken, _, err := utils.GenerateTokens(user.Email, actorId, user.RoleId)
					if err != nil {
						return "", err
					}
					//-------------------------------------------------
					tmpCurTime := time.Now().UTC()
					origin.ActionPeriod = &tmpCurTime
					origin.ActionToken = &actionToken
					//-------------------------------------------------
					if err := tr.userRepo.UpdateUser(origin); err != nil {
						return "", err
					}
					//-------------------------------------------------
					// Send verification mail
					var url string = utils.GenerateCallBackUrl([]string{
						getProcessUrl(),
						actionToken,
						user.UserId,
						updateProfileType,
						user.Email,
					})
					if url == "" {
						return "", errors.New(notis.GenerateMailWarnMsg)
					}
					//-------------------------------------------------
					if err := utils.SendMail(templatePath.UpdateMailTemplate, notis.UpdateMailSubject, spModels.MailBody{
						Url:   url,
						Email: user.Email,
					}); err != nil {
						return "", err
					}
					//-------------------------------------------------
					return notis.UpdateMailMsg, nil
				}
				// Rest case: admin change email for lower role account (Only admin can edit other email except other admins)
				origin.Email = user.Email
			} else {
				return "", errors.New(notis.EmailRegisteredWarnMsg)
			}
		}
	}
	//---------------------------------------
	// Email is empty or the new one is the same as the old
	if err := tr.userRepo.UpdateUser(origin); err != nil {
		return "", err
	}
	//---------------------------------------
	return "Success", nil
}

func (tr userService) VerifyAction(rawToken string) (error, string) {
	var cmp []string = strings.Split(rawToken, ":")
	//---------------------------------------
	if len(cmp) < 3 {
		return errors.New(notis.GenericsErrorWarnMsg), ""
	}
	//---------------------------------------
	var actionToken string = cmp[0]
	var userId string = cmp[1]
	var actionType string = cmp[2]
	//---------------------------------------
	var user dbo.User
	err := utils.VerifyActionToken(actionToken, &user, &tr.userRepo)
	if err != nil {
		return err, ""
	}
	if user.UserId != userId {
		return errors.New(notis.GenericsErrorWarnMsg), ""
	}
	//---------------------------------------
	if actionType != activateType && actionType != resetPassType && actionType != updateProfileType { // Fake url with an invalid type
		return errors.New(notis.GenericsErrorWarnMsg), ""
	}
	//---------------------------------------
	if err := utils.CaseBodyForVerifyActionType(&user, cmp[3], actionType, []string{
		activateType,
		resetPassType,
		updateProfileType,
	},
		&tr.userRepo,
	); err != nil {
		return err, ""
	}
	//---------------------------------------
	var response string
	if actionType == resetPassType || *user.IsHaveToResetPw {
		if err := prepareForcedResetPass(&user, &response); err != nil {
			return err, ""
		}
	}
	//---------------------------------------
	if err := tr.userRepo.UpdateUser(user); err != nil {
		return err, ""
	}
	//---------------------------------------
	return nil, response
}

func (tr userService) VerifyResetPassword(newPass string, re_newPass string, token string) (string, error) {
	var user dbo.User
	err := utils.VerifyActionToken(token, &user, &tr.userRepo)
	if err != nil {
		return "", err
	}
	//---------------------------------------------------
	backUrl := utils.GenerateCallBackUrl([]string{
		getResetPassUrl(),
		token,
	})
	//---------------------------------------------------
	if newPass == "" || re_newPass == "" {
		return backUrl, errors.New(notis.PasswordEmptyWarnMsg)
	}
	//------------------------------------
	if newPass != re_newPass {
		return backUrl, errors.New(notis.PasswordsNotMatchWarnMsg)
	}
	//------------------------------------
	if !utils.IsPasswordSecure(newPass) {
		return backUrl, errors.New(notis.PasswordNotSecureWarnMsg)
	}
	//------------------------------------
	if user.Pasword, err = utils.ToHashString(newPass); err != nil {
		return "", err
	}
	//------------------------------------
	if *user.IsHaveToResetPw { // This case used as recovering account as forced user to execute this action
		user.IsHaveToResetPw = nil
	}
	//------------------------------------
	return "", tr.userRepo.UpdateUser(user)
}

func (tr userService) RecoverAccountByCustomer(email string) (string, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	user, err := tr.userRepo.GetUserByEmail(email)
	//------------------------------------
	if err != nil { // Internal error
		return "", err
	}
	//------------------------------------
	if user == nil { // Acount not exist
		return "", errors.New(notis.UndefinedUserWarnMsg)
	}
	//------------------------------------
	if user.ActiveStatus {
		return "", errors.New(notis.StillActiveAccountMsg)
	}
	//------------------------------------
	if !user.ActiveStatus {
		if user.FailAccess == 6 {
			if user.LastFail != nil { // Ban case -> Deny for recovering, only contact admin
				return "", errors.New(notis.AccountBanWarnMsg)
			}
		} else {
			return "", errors.New(notis.StillActiveAccountMsg)
		}
	}
	//------------------------------------
	// Pass through all these conditions -> send confirmation mail to reactivate this account
	token, _, err := utils.GenerateTokens(email, user.UserId, user.RoleId)
	if err != nil {
		return "", err
	}
	//------------------------------------
	user.ActionToken = &token
	tmpCurTime := time.Now().UTC()
	user.ActionPeriod = &tmpCurTime
	if err := tr.userRepo.UpdateUser(*user); err != nil {
		return "", err
	}
	//------------------------------------
	if err := utils.SendMail(templatePath.AccountRecoveryTemplate, notis.RecoverAccountSubject, spModels.MailBody{
		Email: email,
		Url: utils.GenerateCallBackUrl([]string{
			getProcessUrl(),
			token,
			user.UserId,
			activateType,
		}),
	}); err != nil {
		return "", err
	}
	//------------------------------------
	return notis.RecoverAccountMsg, nil
}

func (tr userService) ChangeUserStatus(rawStatus string, userId string, actorId string) (error, string) {
	var user dbo.User
	var actor dbo.User
	if err := utils.VerifyActorAndObject(actorId, userId, &actor, &user, &tr.userRepo); err != nil {
		return err, ""
	}
	//------------------------------------
	if err := utils.VerifyUpdateAuth(spModels.UserNormalModel{
		UserId: userId,
		RoleId: user.RoleId,
		Email:  user.Email,
	},
		actor,
		user,
	); err != nil {
		return err, ""
	}
	//------------------------------------
	if rawStatus == "" { // Not change
		return nil, ""
	}
	//------------------------------------
	status, err := strconv.ParseBool(rawStatus)
	if err != nil {
		return errors.New(notis.InvalidStatusWarnMsg), ""
	}
	//------------------------------------
	if status == user.ActiveStatus { // Not change
		return nil, ""
	}
	//------------------------------------
	if err := tr.userRepo.ChangeUserStatus(status, userId); err != nil {
		return err, ""
	}
	//------------------------------------
	if actorId == userId { // Self lock
		return nil, LoginPageUrl // redirect user back to login page
	}
	//------------------------------------
	return nil, ""
}

func (tr userService) LogOut(userId string) error {
	user, err := tr.userRepo.GetUserById(userId)
	if err != nil {
		return err
	}
	//------------------------------------
	if user == nil {
		return errors.New(notis.AnonymousWarnMsg)
	}
	//------------------------------------
	user.AccessToken = nil
	user.RefreshToken = nil
	//------------------------------------
	return tr.userRepo.UpdateUser(*user)
}

func prepareForcedResetPass(user *dbo.User, url *string) error {
	token, _, err := utils.GenerateTokens(user.Email, user.UserId, user.RoleId)
	if err != nil {
		return err
	}
	//---------------------------------------
	user.ActionToken = &token
	tmpCurrTime := time.Now().UTC()
	user.ActionPeriod = &tmpCurrTime
	*url = utils.GenerateCallBackUrl([]string{
		getResetPassUrl(),
		token,
	})
	//---------------------------------------
	if *url == "" {
		return errors.New(notis.InternalErr)
	}
	//---------------------------------------
	return nil
}

func generateUserId(repo *irepositories.IUserRepo) (string, error) {
	list, err := (*repo).GetAllUsers()
	if err != nil {
		return "", err
	}
	//---------------------------------------
	return "U" + fmt.Sprintf("%05d", len(*list)+1), nil
}
