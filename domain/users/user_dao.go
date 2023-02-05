package users

import (
	"fmt"

	"github.com/imaarov/bookstore_microservice/datasources/mysql/users_db"
	"github.com/imaarov/bookstore_microservice/utils/date_utils"
	"github.com/imaarov/bookstore_microservice/utils/errors"
	"github.com/imaarov/bookstore_microservice/utils/mysql_utils"
)

const (
	queryUserInsert       = "INSERT INTO users(first_name, last_name, email, date_created, password, status) VALUES(?, ?, ?, ?, ?, ?);"
	queryUserGet          = "SELECT * FROM users WHERE id=?;"
	queryUpdate           = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDelete           = "DELETE FROM users WHERE id=?"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created,status FROM users WHERE status=?"
	errorNoRows           = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {
	err := users_db.Client.Ping()
	if err != nil {
		panic(err)
	}

	statement, statementErr := users_db.Client.Prepare(queryUserGet)
	if statementErr != nil {
		return errors.NewInternalServerError(statementErr.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryUserInsert)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	user.DateCreated = date_utils.GetNowString()
	insertRes, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, idErr := insertRes.LastInsertId()
	if idErr != nil {
		return mysql_utils.ParseError(idErr)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryUpdate)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	_, statementErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
	if err != nil {
		fmt.Println("shit")
		return mysql_utils.ParseError(statementErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryDelete)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	if _, statementErr := statement.Exec(user.Id); statementErr != nil {
		return mysql_utils.ParseError(statementErr)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	statement, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	rows, rowErr := statement.Query(status)
	if rowErr != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if rowFetchErr := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); rowFetchErr != nil {
			return nil, mysql_utils.ParseError(rowFetchErr)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(
			fmt.Sprintf("no users matching given status of %s", status),
		)
	}
	return results, nil
}
