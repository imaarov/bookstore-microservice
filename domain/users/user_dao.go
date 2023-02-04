package users

import (
	"fmt"

	"github.com/imaarov/bookstore_microservice/datasources/mysql/users_db"
	"github.com/imaarov/bookstore_microservice/utils/date_utils"
	"github.com/imaarov/bookstore_microservice/utils/errors"
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
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s Already Exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d Already Exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	userDB[user.Id] = user

	return nil
}
