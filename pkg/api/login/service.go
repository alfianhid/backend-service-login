package login

import (
	"backend-service/pkg/application_service/data_service/user"
	models "backend-service/pkg/utl/models"

	"github.com/Nerzal/gocloak/v8"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Securer represents security interface
type Securer interface {
	Hash(string) string
	Password(string, ...string) bool
}

type DBClientInterface interface {
	FindByUsername(*gorm.DB, string) (*models.Users, error)
	FindByEmail(*gorm.DB, string) (*models.Users, error)
	CheckPassword(*gorm.DB, string, string) (*models.Users, error)
	Update(*gorm.DB, *models.Users) error
}

type Keycloak interface {
	Login(string, string) (*gocloak.JWT, error)
	GetTokenAdmin() (*gocloak.JWT, error)
	ResetPassword(*gocloak.JWT, string) error
}

// Service represents user application interface
type Service interface {
	Login(echo.Context, models.User) (*models.LoginResponse, error)
	ResetPassword(echo.Context, string) (*models.ResetPasswordResponse, error)
}

// RequestHandler represents user application service
type RequestHandler struct {
	sec   Securer
	cloak Keycloak
	db    *gorm.DB
	udb   DBClientInterface
}

// New creates new user RequestHandler application service
func New(db *gorm.DB, udb DBClientInterface, sec Securer, cloak Keycloak) *RequestHandler {
	return &RequestHandler{db: db, udb: udb, sec: sec, cloak: cloak}
}

// Initialize initalizes User RequestHandler application service with defaults
func Initialize(db *gorm.DB, sec Securer, cloak Keycloak) *RequestHandler {
	return New(db, user.NewUserDBClient(), sec, cloak)
}
