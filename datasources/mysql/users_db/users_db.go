package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imaarov/bookstore_microservice/utils/env"
)

var (
	Client *sql.DB
)

func init() {
	// Load Env Variables
	username := env.LoadEnv("DB_USER")
	password := env.LoadEnv("DB_PASSWORD")
	host := env.LoadEnv("DB_HOST")
	name := env.LoadEnv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username,
		password,
		host,
		name,
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	err = Client.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Database Successfully Configured ")
}
