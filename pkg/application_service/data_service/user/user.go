// Package store contains the components necessary for api services
// to interact with the database
package user

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"

	"backend-service/pkg/utl/config"
	models "backend-service/pkg/utl/models"
	"backend-service/pkg/utl/secure"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var (
	ErrRecordNotFound      = echo.NewHTTPError(http.StatusNotFound, "username not found")
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "wrong password")
)

type UserDBClient struct{}

func NewUserDBClient() *UserDBClient {
	return &UserDBClient{}
}

func (u *UserDBClient) FindByUsername(db *gorm.DB, username string) (*models.Users, error) {
	var user = new(models.Users)

	if err := db.Set("gorm:auto_preload", true).Where("username = ?", username).First(user).Error; gorm.IsRecordNotFoundError(err) {
		return user, ErrRecordNotFound
	} else if err != nil {
		log.Panicln(fmt.Sprintf("db connection error %v", err))
		return user, err
	}

	return user, nil
}

func (u *UserDBClient) FindByEmail(db *gorm.DB, email string) (*models.Users, error) {
	var user = new(models.Users)

	if err := db.Set("gorm:auto_preload", true).Where("email = ?", email).First(user).Error; gorm.IsRecordNotFoundError(err) {
		return user, ErrRecordNotFound
	} else if err != nil {
		log.Panicln(fmt.Sprintf("db connection error %v", err))
		return user, err
	}

	return user, nil
}

func (u *UserDBClient) CheckPassword(db *gorm.DB, username, password string) (*models.Users, error) {
	var user = new(models.Users)

	// Store the obtained password in `user` variable
	if err := db.Set("gorm:auto_preload", true).Where("username = ?", username).First(user).Error; gorm.IsRecordNotFoundError(err) {
		return user, ErrRecordNotFound
	} else if err != nil {
		log.Panicln(fmt.Sprintf("db connection error %v", err))
		return user, err
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	var config = new(config.Application)
	sec := secure.New(config.MinPasswordStr, sha1.New())
	match := sec.HashMatchesPassword(user.Password, password)
	if !match {
		return user, ErrPasswordsNotMaching
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
	return user, nil
}

func (u *UserDBClient) Update(db *gorm.DB, user *models.Users) error {
	return db.Save(user).Error
}
