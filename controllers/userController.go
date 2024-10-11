package controllers

import (
	"log"
	"net/http"
	"v3/constants/notis"
	servicegenerator "v3/service_generator"
	"v3/services"
	"v3/spModels"

	"github.com/gin-gonic/gin"
)

func loginAuth(c *gin.Context) {
	var model spModels.LoginModel
	if err := c.ShouldBindJSON(&model); err != nil {
		log.Print(notis.UserControllerMsg + "loginAuth - Error while fetching data - " + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": notis.GenericsErrorWarnMsg})
		return
	}
	//--------------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	response_1, response_2, err := service.Login(model.Email, model.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	if response_1 == "Reset" { // Forced to reset their password as previous recover-account action
		c.Redirect(http.StatusTemporaryRedirect, response_2)
		return
	} else if response_1 == "Activate" { // Case user has registered this account but have not activated it yet
		c.IndentedJSON(http.StatusContinue, gin.H{"message": response_2})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{
		"access_token":  response_1,
		"refresh_token": response_2,
	})
}

func getAllUsers(c *gin.Context) {
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------

	res, err := service.GetAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"users": res})
}

func getUsersByRole(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	res, err := service.GetUsersByRole(c.Param("role"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"users": res})
}

func getUsersByStatus(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	res, err := service.GetUsersByStatus(c.Param("status"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"users": res})
}

func getUserById(c *gin.Context) {
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	res, err := service.GetUserById(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"user": res})
}

func addUser(c *gin.Context) {
	var model spModels.SignUpModel
	if err := c.ShouldBindJSON(&model); err != nil {
		log.Print(notis.UserControllerMsg + "addUser - Error while fetching data - " + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": notis.GenericsErrorWarnMsg})
		return
	}
	//--------------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	err, res := service.AddUser(model, c.Param("userId"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": res})
}

func updateUser(c *gin.Context) {
	var model spModels.UserNormalModel
	if err := c.ShouldBindJSON(&model); err != nil {
		log.Print(notis.UserControllerMsg + "updateUser - Error while fetching data - " + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": notis.GenericsErrorWarnMsg})
		return
	}
	//--------------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	res, err := service.UpdateUser(model, c.GetString("userId")) // Get actorId set from the context as named: userId
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.IndentedJSON(http.StatusAccepted, gin.H{"message": res})
}

func changeUserStatus(c *gin.Context) {
	userId := c.Param("userId")
	rawStatus := c.Param("status")
	actorId := c.GetString("userId")
	//--------------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	err, res := service.ChangeUserStatus(rawStatus, userId, actorId)
	//--------------------------------------------
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	if res == "" {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
		return
	}
	//--------------------------------------------
	c.Redirect(http.StatusPermanentRedirect, res)
}

func verifyAction(c *gin.Context) {
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	err, res := service.VerifyAction(c.Query("rawToken"))
	//--------------------------------------------
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	if res == "" {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
		return
	}
	//--------------------------------------------
	c.Redirect(http.StatusTemporaryRedirect, res)
}

func verifyResetPassword(c *gin.Context) {
	newPass := c.Query("password")
	re_newPass := c.Query("confirmPass")
	token := c.Query("token")
	//--------------------------------------------
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	res, err := service.VerifyResetPassword(newPass, re_newPass, token)
	if err != nil {
		if res != "" {
			c.Redirect(http.StatusTemporaryRedirect, res)
			return
		}
		//--------------------------------------------
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	//--------------------------------------------
	c.Redirect(http.StatusAccepted, services.LoginPageUrl)
}

func recoverAccountByCustomer(c *gin.Context) {
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	if res, err := service.RecoverAccountByCustomer(c.Param("email")); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": res})
	}
}

func logOut(c *gin.Context) {
	service, err := servicegenerator.ConstructUserService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	if err := service.LogOut(c.GetString("userId")); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//--------------------------------------------
	c.Redirect(http.StatusPermanentRedirect, services.LoginPageUrl)
}
