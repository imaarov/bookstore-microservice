package app

import "github.com/imaarov/bookstore_microservice/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
