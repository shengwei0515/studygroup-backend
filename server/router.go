package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"net/http"
	"studygroup/controller"
	_ "studygroup/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const KEY = "sessionkey"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	store := cookie.NewStore([]byte(KEY))
	router.Use(sessions.Sessions("studygroup", store))

	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		helloWorldGroup := v1.Group("/hello-world")
		{
			helloWorldController := new(controller.HelloWorldController)
			helloWorldGroup.GET("/", helloWorldController.GetHelloWorld)
		}

		accountController := new(controller.AccountController)
		v1.POST("/register", accountController.CreateAccount)

		authGroup := v1.Group("/auth")
		{
			authController := new(controller.AuthController)
			authGroup.POST("/login", authController.Login)
			authGroup.GET("/logout", authController.Logout)
		}
		authedGroup := v1.Group("/authed")
		{
			accountGroup := authedGroup.Group("/account", AuthSessionMiddle())
			{
				accountGroup.POST("/", accountController.CreateAccount)
				accountGroup.GET("/", accountController.GetAccountInfo)
				accountGroup.DELETE("/", accountController.DeleteAccount)
				accountGroup.PUT("/resetpwd", accountController.ResetPassword)
				accountGroup.GET("/all", accountController.GetAllAccounts)
			}
		}

	}
	return router
}

func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("userId")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("userId", sessionValue.(string))

		c.Next()
		return
	}
}
