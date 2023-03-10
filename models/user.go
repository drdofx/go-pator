package models

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"pator-api/database"
	"pator-api/utils/token"
	"strings"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null;"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255);not null;" json:"-"`
	NIM      string `gorm:"type:varchar(15);not null;unique"`
	Major    string `gorm:"type:varchar(150);not null;"`
	Year     int    `gorm:"type:int;not null;"`
}

func FindUserByID(id uint) (User, error) {
	var user User

	err := database.DB.First(&user, id).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) Login(email string, password string) (string, error) {

	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return "", err
	}

	// split encoded hash into parts
	hashParts := strings.Split(user.Password, "$")

	// base64 decode salt and hashed password
	salt, err := base64.RawStdEncoding.DecodeString(hashParts[4])
	if err != nil {
		return "", err
	}

	hashedPassword, err := base64.RawStdEncoding.DecodeString(hashParts[5])
	if err != nil {
		return "", err
	}

	// hash user password with argon2
	key := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// compare hashed password with hashed password from database
	if subtle.ConstantTimeCompare(key, hashedPassword) == 1 {
		token, err := token.GenerateJWT(user.ID)

		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", errors.New("invalid login credentials")
}

func (user *User) SaveUser() (*User, error) {
	err := database.DB.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	// generate random salt
	salt := make([]byte, 16)
	_, err := rand.Read(salt)

	if err != nil {
		return err
	}

	// hash user password with argon2
	key := argon2.IDKey([]byte(user.Password), salt, 1, 64*1024, 4, 32)

	// base64 encode salt and hashed password
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Key := base64.RawStdEncoding.EncodeToString(key)

	// set standard encoded salt and hashed password
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, 64*1024, 1, 4, base64Salt, base64Key)

	// set user password to encoded hash
	user.Password = encodedHash

	return nil
}
