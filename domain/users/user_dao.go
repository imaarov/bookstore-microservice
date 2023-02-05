package users

import (
	"fmt"

	"github.com/imaarov/bookstore_microservice/datasources/mysql/users_db"
	"github.com/imaarov/bookstore_microservice/utils/date_utils"
	"github.com/imaarov/bookstore_microservice/utils/errors"
)

const (
	queryUserInsert = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	err := users_db.Client.Ping()
	if err != nil {
		panic(err)
	}
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	statement, err := users_db.Client.Prepare(queryUserInsert)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer statement.Close()

	user.DateCreated = date_utils.GetNowString()
	insertRes, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error while trying to save user: %s", err.Error()),
		)
	}
	userId, err := insertRes.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(
			fmt.Sprintf("Error while trying  fetch last id: %s", err.Error()),
		)
	}
	user.Id = userId

	return nil
}
