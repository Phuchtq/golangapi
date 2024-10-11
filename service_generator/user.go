package servicegenerator

import (
	"v3/db"
	iservices "v3/interfaces/iServices"
	"v3/repositories"
	"v3/services"
)

func ConstructUserService() (iservices.IUserService, error) {
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	//-----------------------------------
	userRepo := repositories.InitializeUserRepo(db)
	roleRepo := repositories.InitializeRoleRepo(db)
	//-----------------------------------
	return services.InitializeUserService(userRepo, roleRepo), nil
}
