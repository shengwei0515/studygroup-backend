package pgmigrate

import (
	"studygroup/db/postgres"
	"studygroup/model"
)

func AutoMigratePostgres() {
	db := postgres.GetDb()

	db.AutoMigrate(&model.Account{})
}
