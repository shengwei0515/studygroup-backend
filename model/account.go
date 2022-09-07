package model

import (
	"studygroup/db/postgres"
	"studygroup/form"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Id       string    `json:"id" gorm:"column:id;primaryKey"`
	Name     string    `json:"name" gorm:"column:name"`
	Password string    `json:"password" gorm:"column:password"`
	CreateOn time.Time `json:"password" gorm:"column:CreateOn"`
}

func (Account) TableName() string {
	return "account"
}

func (a Account) Create(accountPayload form.AccountSignup) (*gorm.DB, error) {
	db := postgres.GetDb()
	id := uuid.NewV1()

	hashPwd, err := HashPassword(accountPayload.Password)
	if err != nil {
		return nil, err
	}

	account := Account{
		Id:       id.String(),
		Name:     accountPayload.Name,
		Password: hashPwd,
		CreateOn: time.Now(),
	}

	result := db.Create(&account)
	if err = result.Error; err != nil {
		return nil, err
	}
	return result, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
