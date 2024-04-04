package routes

import (
	"github.com/gin-gonic/gin"
	"hh_test_autho/internal/api"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	car := api.NewCarApi()

	r.GET("/api/car", car.Get)
	r.GET("/api/car/:id", car.GetID)
	r.POST("/api/car", car.Post)
	r.DELETE("/api/car:id", car.Delete)
	r.PUT("/api/car/:id", car.Update)

	return r
}
