package login

import (
	"time"

	"backend-service/pkg/api/login"
	models "backend-service/pkg/utl/models"

	"github.com/labstack/echo"
)

const packageName = "login"

// LogService represents user logging service
type LogService struct {
	login.Service
	logger models.Logger
}

// New creates new user logging service
func New(svc login.Service, logger models.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// Create logging
func (ls *LogService) Login(c echo.Context, req models.User) (resp *models.LoginResponse, err error) {
	dupe := req
	dupe.Password = "xxx-xxx-xxx"
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			packageName, "Login user request", err,
			map[string]interface{}{
				"req":  dupe,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Login(c, req)
}

// Create logging
func (ls *LogService) ResetPassword(c echo.Context, email string) (resp *models.ResetPasswordResponse, err error) {
	dupe := email
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			packageName, "Reset password user request", err,
			map[string]interface{}{
				"req":  dupe,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.ResetPassword(c, email)
}
