package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(255)" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(255)" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(255)" validate:"required"`
	FullName    string    `json:"full_name" gorm:"column:full_name;type:varchar(255)" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255)" validate:"required"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (user User) Validate() error {
	v := validator.New()
	return v.Struct(user)
}

type UserSession struct {
	ID                  int       `gorm:"primaryKey"`
	UserID              string    `json:"user_id" gorm:"column:user_id;type:int" validate:"required"`
	Token               string    `json:"token" gorm:"column:token;type:varchar(255)" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"column:refresh_token;type:varchar(255)" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
	CreatedAt           time.Time `json:"-"`
	UpdatedAt           time.Time `json:"-"`
}

func (*UserSession) TableName() string {
	return "user_session"
}

func (userSession UserSession) Validate() error {
	v := validator.New()
	return v.Struct(userSession)
}
