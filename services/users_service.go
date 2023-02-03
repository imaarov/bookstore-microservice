package services

import (
	"github.com/imaarov/bookstore_microservice/domain/users"
	"github.com/imaarov/bookstore_microservice/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
