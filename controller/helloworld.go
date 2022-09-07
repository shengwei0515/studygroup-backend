package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWroldController struct{}

type HelloWorld struct {
	Message string `json:"message"`
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags hello world
// @Accept json
// @Produce json
// @Success 200 {object} HelloWorld
// @Router /hello-world/ [get]
func (controller HelloWroldController) GetHelloWorld(ctx *gin.Context) {

	response := new(HelloWorld)
	response.Message = "Hello World !"

	ctx.JSON(http.StatusOK, response)
}
