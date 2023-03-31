package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.POST("/create-diagram", CreateDiagramWithRawData)
	r.POST("/")
}
