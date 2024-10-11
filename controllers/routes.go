package controllers

import (
	"log"
	"os"
	"v3/middlewares"

	"github.com/gin-gonic/gin"
)

const (
	userContextPath string = "/users"
	roleContextPath string = "/roles"
)

func InitializeAPIRoutes() {
	server := gin.Default()
	//-----------------------------------------
	initializeUserAPI(*server)
	initializeRoleAPI(*server)
	//-----------------------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	//-----------------------------------------
	if err := server.Run(":" + port); err != nil {
		log.Fatal("InitializeAPIRoutes meet an unexpected error: ", err)
	}
}

func initializeUserAPI(server gin.Engine) {
	server.POST(userContextPath, addUser)
	server.POST("/login", loginAuth)
	server.POST("/logout", logOut)
	//-----------------------------------------
	server.GET(userContextPath+"/verifyaction", verifyAction)
	//-----------------------------------------
	server.PUT(userContextPath+"/verifyreset", verifyResetPassword)
	server.PUT(userContextPath+"/email/:email", recoverAccountByCustomer)
	//-----------------------------------------
	authGroup := server.Group("/")
	authGroup.Use(middlewares.Authorize)
	//-----------------------------------------
	authGroup.GET(userContextPath, getAllUsers)
	authGroup.GET(userContextPath+"/role/:role", getUsersByRole)
	authGroup.GET(userContextPath+"/status/:status", getUsersByStatus)
	authGroup.GET(userContextPath+"/:id", getUserById)
	//-----------------------------------------
	authGroup.PUT(userContextPath, updateUser)
	authGroup.PUT(userContextPath+"/:id/status/:status", changeUserStatus)
}

func initializeRoleAPI(server gin.Engine) {
	// These access just be available for staff/admin role
	authGroup := server.Group("/")
	authGroup.Use(middlewares.Authorize)
	//-----------------------------------------
	authGroup.GET(roleContextPath, getAllRoles)
	authGroup.GET(roleContextPath+"/name/:name", getRolesByName)
	authGroup.GET(roleContextPath+"/status/:status", getRolesByStatus)
	authGroup.GET(roleContextPath+"/:id", getRoleById)
	//-----------------------------------------
	authGroup.POST(roleContextPath+"/:name", createRole)
	//-----------------------------------------
	authGroup.PUT(roleContextPath, updateRole)
	authGroup.PUT(roleContextPath+"/:id", activateRole)
	//-----------------------------------------
	authGroup.DELETE(roleContextPath+"/:id", removeRole)
}
