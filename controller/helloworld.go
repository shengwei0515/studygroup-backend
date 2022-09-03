package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloWorld struct {
	Message string `json:"message"`
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} HelloWorld
// @Router /hello-world/ [get]
func GetHelloWorld(c *gin.Context) {

	response := new(HelloWorld)
	response.Message = "Hello World !"

	c.JSON(http.StatusOK, response)
}
