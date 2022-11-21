// Package transport contains the HTTP service for user interactions
package transport

import (
	"net/http"

	"backend-service/pkg/api/login"
	models "backend-service/pkg/utl/models"

	"github.com/labstack/echo"
)

// Custom errors
var (
	ErrUnknownPayload      = echo.NewHTTPError(http.StatusBadRequest, "payload is unknown")
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "passwords do not match")
)

// HTTP represents user http service
type HTTP struct {
	svc login.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc login.Service, er *echo.Group) {
	h := HTTP{svc}

	er.POST("/login", h.login)
	// er.POST("/reset-password", h.resetpassword)
}

// createReq is a used to serialize the request payload to a struct
type createReq struct {
	Username string `json:"username" validate:"required,min=3,alphanum"`
	Password string `json:"password" validate:"required,min=8"`
}

// responses:
//  200: userResp
//  400: errMsg
//  401: err
//  403: errMsg
//  500: err
func (h *HTTP) login(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {
		return err
	}

	usr, err := h.svc.Login(c, models.User{
		Username: r.Username,
		Password: r.Password,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}

// resetReq is a used to serialize the request payload to a struct
// type resetReq struct {
// 	Email string `json:"email" validate:"required,email"`
// }

// func (h *HTTP) resetpassword(c echo.Context) error {
// 	r := new(resetReq)

// 	if err := c.Bind(r); err != nil {
// 		return err
// 	}

// 	usr, err := h.svc.ResetPassword(c, r.Email)

// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, usr)
// }
