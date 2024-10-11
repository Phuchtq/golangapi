package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"v3/constants/notis"
	"v3/dbo"
	irepositories "v3/interfaces/iRepositories"
)

type userRepo struct {
	db *sql.DB
}

func InitializeUserRepo(db *sql.DB) irepositories.IUserRepo {
	return &userRepo{
		db: db,
	}
}

func (tr *userRepo) GetAllUsers() (*[]dbo.User, error) {
	errLogMsg := notis.UserRepoMsg + "GetAllUsers - "
	query := "Select id, email, password, roleId, activeStatus from Users"

	rows, err := tr.db.Query(query)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()

	var res []dbo.User
	for rows.Next() {
		var x dbo.User
		if err := rows.Scan(&x.UserId, &x.Email, &x.Pasword, &x.RoleId, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}
		res = append(res, x)
	}
	return &res, nil
}

func (tr *userRepo) GetUsersByRole(id string) (*[]dbo.User, error) {
	errLogMsg := notis.UserRepoMsg + "GetUsersByRole - "
	query := "Select id, email, password, roleId, activeStatus from Users where roleId = ?"

	rows, err := tr.db.Query(query, id)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()

	var res []dbo.User
	for rows.Next() {
		var x dbo.User
		if err := rows.Scan(&x.UserId, &x.Email, &x.Pasword, &x.RoleId, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}
		res = append(res, x)
	}
	return &res, nil
}

// func (tr *userRepo) GetUsersFromFilterFailCase() (*[]dbo.User, error) {
// 	errLogMsg := notis.UserRepoMsg + "GetUsersFromFailCase - "

// 	query := "Select email where actionToken is null or actionPeriod is null or actionPeriod < (now() - interval '3 hours') or accessFail > 0 or activeStatus = true"
// 	rows, err := tr.db.Query(query)
// 	if err != nil {
// 		tr.db.Close()
// 		log.Print(errLogMsg, err)
// 		return nil, errors.New(notis.InternalErr)
// 	}
// 	defer rows.Close()

// 	var res []dbo.User
// 	for rows.Next() {
// 		var x dbo.User
// 		if err := rows.Scan(&x.Email); err != nil {
// 			tr.db.Close()
// 			log.Print(errLogMsg, err)
// 			return nil, errors.New(notis.InternalErr)
// 		}
// 		res = append(res, x)
// 	}
// 	return &res, nil
// }

// func (tr *userRepo) GetUsersExistWithInputtedEmail(email string) (*[]dbo.User, error) {
// 	errLogMsg := notis.UserRepoMsg + "GetUsersExistWithInputtedEmail - "

// 	query := "Select userId where actionToken is null or actionPeriod is null or actionPeriod < (now() - interval '3 hours') or accessFail > 0 or activeStatus = true and email = ?"
// 	rows, err := tr.db.Query(query, email)
// 	if err != nil {
// 		tr.db.Close()
// 		log.Print(errLogMsg, err)
// 		return nil, errors.New(notis.InternalErr)
// 	}
// 	defer rows.Close()

// 	var res []dbo.User
// 	for rows.Next() {
// 		var x dbo.User
// 		if err := rows.Scan(&x.UserId); err != nil {
// 			tr.db.Close()
// 			log.Print(errLogMsg, err)
// 			return nil, errors.New(notis.InternalErr)
// 		}
// 		res = append(res, x)
// 	}
// 	return &res, nil
// }

func (tr *userRepo) GetUsersByStatus(status bool) (*[]dbo.User, error) {
	errLogMsg := notis.UserRepoMsg + "GetUsersByStatus - "
	query := "Select id, email, password, roleId, activeStatus from Users where activeStatus = ?"

	rows, err := tr.db.Query(query, status)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	defer rows.Close()

	var res []dbo.User
	for rows.Next() {
		var x dbo.User
		if err := rows.Scan(&x.UserId, &x.Email, &x.Pasword, &x.RoleId, &x.ActiveStatus); err != nil {
			tr.db.Close()
			log.Print(errLogMsg, err)
			return nil, errors.New(notis.InternalErr)
		}
		res = append(res, x)
	}
	return &res, nil
}

func (tr userRepo) GetUserById(id string) (*dbo.User, error) {
	errLogMsg := notis.UserRepoMsg + "GetUserById - "
	query := "Select id, email, password, roleId, activeStatus from Users where id = ?"
	var res dbo.User

	if err := tr.db.QueryRow(query, id).Scan(&res.UserId, &res.Email, &res.Pasword, &res.RoleId, &res.ActiveStatus); err != nil && err == sql.ErrNoRows {
		tr.db.Close()
		return nil, nil // No data found with incoming ID parameter - actually not considered as an error -> no data and error returned
	} else if err != nil && err != sql.ErrNoRows {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}

	tr.db.Close()
	return &res, nil
}

// func (tr *userRepo) GetUserByLogin(email string, password string) (*dbo.User, error) {
// 	errLogMsg := notis.UserRepoMsg + "GetUserByLogin - "
// 	query := "Select id, roleId, activeStatus from Users where email = ? and password = ?"
// 	var res dbo.User
// 	if err := tr.db.QueryRow(query, email, password).Scan(&res.UserId, &res.RoleId, &res.ActiveStatus, &res.LastFail); err != nil && err == sql.ErrNoRows {
// 		tr.db.Close()
// 		return nil, nil
// 	} else if err != nil && err != sql.ErrNoRows {
// 		tr.db.Close()
// 		log.Print(errLogMsg, err)
// 		return nil, errors.New(notis.InternalErr)
// 	}
// 	tr.db.Close()
// 	res.Email = email
// 	return &res, nil
// }

func (tr *userRepo) GetUserByEmail(email string) (*dbo.User, error) {
	errLogMsg := notis.UserRepoMsg + "GetUserByEmail - "
	query := "Select id, roleId, activeStatus, lastFail, failAccess, password from Users where lower(email) = lower($1)"
	var res dbo.User
	if err := tr.db.QueryRow(query, email).Scan(&res.UserId, &res.RoleId, &res.ActiveStatus, &res.LastFail, &res.FailAccess, &res.Pasword); err != nil && err == sql.ErrNoRows {
		tr.db.Close()
		return nil, nil
	} else if err != nil && err != sql.ErrNoRows {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return nil, errors.New(notis.InternalErr)
	}
	tr.db.Close()
	res.Email = email
	return &res, nil
}

func (tr *userRepo) AddUser(u dbo.User) error {
	errLogMsg := notis.UserRepoMsg + "AddUser - "
	query := "Insert into Users(id, email, password, roleId, activeStatus, failAccess, lastFail) values (?, ?, ?, ?, ?, ?, ?)"
	if _, err := tr.db.Exec(query, u.UserId, u.Email, u.Pasword, u.RoleId, u.ActiveStatus, u.FailAccess, u.LastFail); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	tr.db.Close()
	return nil
}

func (tr *userRepo) UpdateUser(u dbo.User) error {
	errLogMsg := notis.UserRepoMsg + "UpdateUser - "
	query := "Update Users set email = ?, password = ?, roleId = ?, accessToken = ?, refreshToken = ?, activeStatus = ?, failAccess = ?, lastFail = ? where id = ?"
	res, err := tr.db.Exec(query, u.Email, u.Pasword, u.RoleId, u.AccessToken, u.RefreshToken, u.ActiveStatus, u.FailAccess, u.LastFail, u.UserId)
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	if rowsAffected == 0 {
		tr.db.Close()
		return errors.New("User not found")
	}
	tr.db.Close()
	return nil
}

func (tr *userRepo) ChangeUserStatus(status bool, id string) error {
	errLogMsg := notis.UserRepoMsg + "ChangeUserStatus - "
	lastFailValueQuery := "NULL"
	if status {
		lastFailValueQuery = fmt.Sprint(time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC))
	}

	query := "Update Users set activeStatus = " + fmt.Sprint(status) + ", failAccess = 0, lastFail = " + lastFailValueQuery + ", accessToken = NULL, refreshToken = NULL, actionPeriod = NULL, actionToken = NULL, isHaveToResetPw = NULL where id = ?"
	if _, err := tr.db.Exec(query, id); err != nil {
		tr.db.Close()
		log.Print(errLogMsg, err)
		return errors.New(notis.InternalErr)
	}
	tr.db.Close()
	return nil
}
