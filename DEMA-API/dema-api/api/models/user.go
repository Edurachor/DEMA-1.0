package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	HASHSEED = 10
)

type User struct {
	*gorm.Model
	Name     string `gorm:"type:varchar(20); not null" json:"name"`
	Email    string `gorm:"type:varchar(40); not null; uniqueIndex; index" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func (user *User) HashPassword(password string) error {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), HASHSEED)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
