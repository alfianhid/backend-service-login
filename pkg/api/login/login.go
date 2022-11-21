package login

import (
	"net/http"

	models "backend-service/pkg/utl/models"

	"github.com/labstack/echo"
)

// Custom errors
var (
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

func (a *RequestHandler) Login(c echo.Context, req models.User) (*models.LoginResponse, error) {
	// check if user is existed
	u, err := a.udb.FindByUsername(a.db, req.Username)
	if err != nil {
		return nil, err
	}

	// check password match
	u, err = a.udb.CheckPassword(a.db, u.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// login processed
	token, err := a.cloak.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// update last login
	u.UpdateLastLogin()
	if err := a.udb.Update(a.db, u); err != nil {
		return nil, err
	}

	return &models.LoginResponse{Status: 200, Message: u.UserID, AccessToken: token.AccessToken, RefreshToken: token.RefreshToken, ExpiresIn: token.ExpiresIn}, nil
}

func (a *RequestHandler) ResetPassword(c echo.Context, email string) (*models.ResetPasswordResponse, error) {
	u, err := a.udb.FindByEmail(a.db, email)
	if err != nil {
		return nil, err
	}

	token, err := a.cloak.GetTokenAdmin()
	if err != nil {
		return nil, err
	}

	err = a.cloak.ResetPassword(token, u.UserID)
	if err != nil {
		return nil, err
	}

	u.UpdateLastPasswordChange()
	if err := a.udb.Update(a.db, u); err != nil {
		return nil, err
	}

	return &models.ResetPasswordResponse{Status: 200, Message: true}, nil
}
