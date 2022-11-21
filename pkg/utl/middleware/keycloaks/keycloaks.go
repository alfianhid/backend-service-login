// Package jsonwebtoken contains logic for using JSON web tokens
package keycloaks

import (
	"context"

	"backend-service/pkg/utl/config"

	"github.com/Nerzal/gocloak/v8"
)

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	server        string
	realm         string
	client_id     string
	client_secret string
	path_key      string
	user_admin    string
	pass_admin    string
	realm_admin   string
	client        gocloak.GoCloak
	ctx           context.Context
}

// New generates new JWT service necessery for auth middleware
func New(cfg *config.Keycloaks) *Service {
	return &Service{
		server:        cfg.Server,
		realm:         cfg.Realm,
		client_id:     cfg.ClientId,
		client_secret: cfg.ClientSecret,
		path_key:      cfg.PathKey,
		user_admin:    cfg.UserAdmin,
		pass_admin:    cfg.PassAdmin,
		realm_admin:   cfg.RealmAdmin,
		client:        gocloak.NewClient(cfg.Server),
		ctx:           context.Background(),
	}
}

func (j *Service) Login(username, password string) (*gocloak.JWT, error) {
	token, err := j.client.Login(j.ctx, j.client_id, j.client_secret, j.realm, username, password)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j *Service) GetTokenAdmin() (*gocloak.JWT, error) {
	token, err := j.client.LoginAdmin(j.ctx, j.user_admin, j.pass_admin, j.realm_admin)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j *Service) ResetPassword(token *gocloak.JWT, UserID string) error {
	params := gocloak.ExecuteActionsEmail{
		ClientID: &(j.client_id),
		UserID:   &UserID,
		Actions:  &[]string{"UPDATE_PASSWORD"},
	}

	err := j.client.ExecuteActionsEmail(context.Background(), token.AccessToken, j.realm, params)

	if err != nil {
		return err
	}

	return nil
}
