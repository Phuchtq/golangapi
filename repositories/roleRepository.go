package repositories

import (
	"database/sql"
	"errors"
	"log"
	"v3/constants/notis"
	"v3/dbo"
	irepositories "v3/interfaces/iRepositories"
)

type roleRepo struct {
	db *sql.DB
}

func InitializeRoleRepo(db *sql.DB) irepositories.IRoleRepo {
	return &roleRepo{
		db: db,
	}
}

func (tr *roleRepo) GetAllRoles() (*[]dbo.Role, error) {
	errLogMsg := notis.RoleRepoMsg + "GetAllRoles - "
	query := "Select * from Roles"

	rows, err := tr.db.Query(query)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()

	var res []dbo.Role
	for rows.Next() {
		var x dbo.Role
		if err := rows.Scan(&x.RoleId, &x.RoleName, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}
		res = append(res, x)
	}

	return &res, nil
}

func (tr *roleRepo) GetRolesByName(name string) (*[]dbo.Role, error) {
	errLogMsg := notis.RoleRepoMsg + "GetRolesByName - "
	query := "Select * from Roles where lower(roleName) like lower($1)"
	rows, err := tr.db.Query(query, "%"+name+"%")
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()

	var res []dbo.Role
	for rows.Next() {
		var x dbo.Role
		if err := rows.Scan(&x.RoleId, &x, x.RoleName, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}
		res = append(res, x)
	}

	return &res, nil
}

func (tr *roleRepo) GetRolesByStatus(status bool) (*[]dbo.Role, error) {
	errLogMsg := notis.RoleRepoMsg + "GetRolesByStatus - "
	query := "Select * from Roles where activeStatus = $1"

	rows, err := tr.db.Query(query)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()
	var res []dbo.Role
	for rows.Next() {
		var x dbo.Role
		if err := rows.Scan(&x.RoleId, &x, x.RoleName, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}
		res = append(res, x)
	}
	return &res, nil
}

func (tr *roleRepo) GetRoleById(id string) (*dbo.Role, error) {
	errLogMsg := notis.RoleRepoMsg + "GetRoleById - "

	query := "Select * from Roles where roleId = $1"
	var res dbo.Role
	if err := tr.db.QueryRow(query, id).Scan(&res.RoleId, &res.RoleName, &res.ActiveStatus); err != nil && err == sql.ErrNoRows {
		tr.db.Close()
		return nil, nil // No data found with incoming ID parameter - actually not considered as an error -> no data and error returned
	} else if err != nil && err != sql.ErrNoRows {
		tr.db.Close()
		log.Print(errLogMsg, err) // Error but bot caused of None-data found - Return error
		return nil, errors.New(notis.InternalErr)
	}
	tr.db.Close()
	return &res, nil
}

func (tr *roleRepo) CreateRole(r dbo.Role) error {
	errLogMsg := notis.RoleRepoMsg + "CreateRole - "
	query := "Insert into Roles(roleId, roleName, activeStatus) values (?, ?, ?)"
	if _, err := tr.db.Exec(query, r.RoleId, r.RoleName, r.ActiveStatus); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	tr.db.Close()
	return nil
}

func (tr *roleRepo) RemoveRole(id string) error {
	errLogMsg := notis.RoleRepoMsg + "RemoveRole - "
	query := "Update Roles set activeStatus = false where id = ?"
	if res, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	} else {
		if rowsAffected, err := res.RowsAffected(); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return errors.New(notis.InternalErr)
		} else if rowsAffected == 0 {
			return errors.New(notis.UndefinedRoleWarnMsg)
		}
	}
	tr.db.Close()
	return nil
}

func (tr *roleRepo) UpdateRole(r dbo.Role) error {
	errLogMsg := notis.RoleRepoMsg + "UpdateRole - "
	query := "Update Roles set roleName = ?, activeStatus = ? where id = ?"
	if res, err := tr.db.Exec(query, r.RoleName, r.ActiveStatus, r.RoleId); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	} else {
		if rowsAffected, err := res.RowsAffected(); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return errors.New(notis.InternalErr)
		} else if rowsAffected == 0 {
			return errors.New(notis.UndefinedRoleWarnMsg)
		}
	}
	tr.db.Close()
	return nil
}

func (tr *roleRepo) ActivateRole(id string) error {
	errLogMsg := notis.RoleRepoMsg + "ActivateRole - "

	query := "Update Roles set activeStatus = true where id = ?"
	if _, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	tr.db.Close()
	return nil
}
