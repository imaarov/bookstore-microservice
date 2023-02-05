package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/imaarov/bookstore_microservice/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewBadRequestError("no record matching with given id")
		}
		return errors.NewInternalServerError("error parsing db response" + err.Error())
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Invalid data %s", err.Error()))
	}
	return errors.NewInternalServerError("error processing the request ")
}
