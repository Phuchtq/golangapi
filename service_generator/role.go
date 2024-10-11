package servicegenerator

import (
	"v3/db"
	iservices "v3/interfaces/iServices"
	"v3/services"
)

func ConstructRoleService() (iservices.IRoleService, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	//---------------------------------------
	//roleRepo := repositories.InitializeRoleRepo(db)
	//---------------------------------------
	return services.InitializeRoleService(db), nil
}
