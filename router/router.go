package router

import (
	"github.com/gin-gonic/gin"

	"studygroup/controller"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	setupHelloWorldGeroup(v1)
}

func setupHelloWorldGeroup(g *gin.RouterGroup) {
	rg := g.Group("/hello-world")

	rg.GET("/", controller.GetHelloWorld)
}
