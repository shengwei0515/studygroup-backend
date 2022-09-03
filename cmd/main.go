package main

import (
	"github.com/gin-gonic/gin"

	"studygroup/db/postgres"
	_ "studygroup/docs"
	"studygroup/router"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title          	Studygroup
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   API Support
// @contact.email  shengwei199505@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /api/v1
func main() {

	db := postgres.GetConnection()
	defer db.Close()

	r := gin.Default()

	router.SetupRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run("0.0.0.0:8080")
}
