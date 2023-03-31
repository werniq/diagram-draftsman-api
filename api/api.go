package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/werniq/diagram-creating-api/Logger"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		Logger.Logger().Printf("error loading .env file: %v\n", err)
		return
	}

	r := gin.Default()

	SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		Logger.Logger().Println("error running server: ", err)
	}
}