package model

import (
	"studygroup/db/postgres"
	"studygroup/form"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Account struct {
	Id       string `json:"id" gorm:"column:id;primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"cloumn:password"`
}

func (Account) TableName() string {
	return "account"
}

func (a Account) Create(accountPayload form.AccountSignup) (*gorm.DB, error) {
	db := postgres.GetDb()
	id := uuid.NewV1()

	account := Account{
		Id:       id.String(),
		Name:     accountPayload.Name,
		Password: accountPayload.Password,
	}

	result := db.Create(&account)
	if err := result.Error; err != nil {
		return nil, err
	}
	return result, nil
}
