package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}

// func LoadEnvPath(name string) {
// 	err := godotenv.Load(name)
// 	if err != nil {
// 		panic(err)
// 	}
// }
