package main

import (
	"studygroup/db/postgres"
	"studygroup/model/pgmigrate"
	"studygroup/server"
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

	postgres.InitWithRetry()
	defer postgres.CloseDb()

	pgmigrate.AutoMigratePostgres()

	server.Init()
}
