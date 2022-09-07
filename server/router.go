package server

import (
	"net/http"
	"studygroup/controller"
	_ "studygroup/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		helloWorldGroup := v1.Group("/hello-world")
		{
			helloWorldController := new(controller.HelloWroldController)
			helloWorldGroup.GET("/", helloWorldController.GetHelloWorld)
		}
		accountGroup := v1.Group("/account")
		{
			accountController := new(controller.AccountController)
			accountGroup.POST("/", accountController.CreateAccount)
		}
	}
	return router
}
