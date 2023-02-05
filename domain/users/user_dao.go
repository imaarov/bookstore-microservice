package users

import (
	"github.com/imaarov/bookstore_microservice/datasources/mysql/users_db"
	"github.com/imaarov/bookstore_microservice/utils/date_utils"
	"github.com/imaarov/bookstore_microservice/utils/errors"
	"github.com/imaarov/bookstore_microservice/utils/mysql_utils"
)

const (
	queryUserInsert = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryUserGet    = "SELECT * FROM users WHERE id=?;"
	queryUpdate     = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDelete     = "DELETE FROM users WHERE id=?"
	errorNoRows     = "no rows in result set"
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
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
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
	insertRes, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		mysql_utils.ParseError(saveErr)
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

	_, statementErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
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
