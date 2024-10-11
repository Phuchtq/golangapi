package controllers

import (
	"log"
	"net/http"
	"v3/constants/notis"
	"v3/dbo"
	servicegenerator "v3/service_generator"

	"github.com/gin-gonic/gin"
)

func getAllRoles(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	res, err := service.GetAllRoles()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"roles": res})
}

func getRolesByName(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	res, err := service.GetRolesByName(c.Param("name"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"roles": res})
}

func getRolesByStatus(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	res, err := service.GetRolesByStatus(c.Param("status"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"roles": res})
}

func getRoleById(c *gin.Context) {
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	if res, err := service.GetRoleById(c.Param("id")); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"role": res})
	}
}

func createRole(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	if err := service.CreateRole(c.Param("name")); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
}

func updateRole(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	var role dbo.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		log.Print(notis.RoleControllerMsg + "updateRole while fetching data - " + err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": notis.GenericsErrorWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	if err := service.UpdateRole(role); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
}

func removeRole(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	if err := service.RemoveRole(c.Param("id")); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
}

func activateRole(c *gin.Context) {
	if !isAdminAccess(c.GetString("role")) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": notis.GenericsRightAccessWarnMsg})
		return
	}
	//-----------------------------------------
	service, err := servicegenerator.ConstructRoleService()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	//-----------------------------------------
	if err := service.ActivateRole(c.Param("id")); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	//-----------------------------------------
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
}
