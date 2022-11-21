package models

import (
	"time"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

type Users struct {
	Base
	UserID           string `json:"user_id"`
	Email            string `json:"email"`
	Username         string `json:"username"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
	Company          string `json:"company"`
	BusinessRelation string `json:"business_relation"`
	Password         string `json:"password"`
	Approval         bool   `json:"approval" gorm:"default:false"`
	UrlImgProfile    string `json:"url_image_profile"`

	LastLogin          time.Time `json:"last_login" gorm:"default:CURRENT_TIMESTAMP"`
	LastPasswordChange time.Time `json:"last_password_change" gorm:"default:CURRENT_TIMESTAMP"`
}

func (u *Users) UpdateLastPasswordChange() {
	u.LastPasswordChange = time.Now()
}

func (u *Users) UpdateLastLogin() {
	u.LastLogin = time.Now()
}
