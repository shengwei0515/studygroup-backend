package model

import (
	"studygroup/db/postgres"
	"studygroup/form"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Account struct {
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`

	Id       string `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
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
	}

	result := db.Create(&account)
	if err = result.Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a Account) Get(accountPayload form.AccountSignup) (*Account, error) {
	db := postgres.GetDb()

	account := Account{}
	raw := `SELECT * FROM account where name = ?`
	result := db.Raw(raw, accountPayload.Name).Scan(&account)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (a Account) GetFromId(id string) (*Account, error) {
	db := postgres.GetDb()

	account := Account{}
	raw := `SELECT * FROM account where id = ?`
	result := db.Raw(raw, id).Scan(&account)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &account, nil
}

func (a Account) Delete(id string) error {
	db := postgres.GetDb()

	raw := `DELETE FROM account where id = ?`
	result := db.Exec(raw, id)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (a Account) UpdatePasswd(id string, password string) (*gorm.DB, error) {
	hashPwd, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	db := postgres.GetDb()

	raw := `UPDATE account SET password = ? WHERE id = ?;`
	result := db.Exec(raw, hashPwd, id)
	if err := result.Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a Account) GetAll() (*[]Account, error) {
	db := postgres.GetDb()

	var accounts []Account
	result := db.Raw("select * from account").Scan(&accounts)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &accounts, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}