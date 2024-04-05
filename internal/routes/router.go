package routes

import (
	"github.com/gin-gonic/gin"
	"hh_test_autho/internal/api"
	"hh_test_autho/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Timer())

	car := api.NewCarApi()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/api/car", car.Get)
	r.GET("/api/car/:id", car.GetID)
	r.POST("/api/car", car.Post)
	r.DELETE("/api/car/:id", car.Delete)
	r.PUT("/api/car/:id", car.Update)

	return r
}
