package controller

import (
	"net/http"
	"studygroup/model"

	"github.com/gin-gonic/gin"
)

type HelloWroldController struct{}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} HelloWorld
// @Router /hello-world/ [get]
func (controller HelloWroldController) GetHelloWorld(ctx *gin.Context) {

	response := new(model.HelloWorld)
	response.Message = "Hello World !"

	ctx.JSON(http.StatusOK, response)
}
