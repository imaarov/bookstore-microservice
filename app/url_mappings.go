package app

import (
	"github.com/imaarov/bookstore_microservice/controllers/ping"
	"github.com/imaarov/bookstore_microservice/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internal/users/search", users.Search)
}
