package models

import (
	"time"
)

type User struct {
	UserId                     int64     `json:"user_id"`
	Username                   string    `json:"username"`
	Email                      string    `json:"email"`
	FullName                   string    `json:"full_name"`
	Password                   string    `json:"password"`
	RoleShortName              string    `json:"role_short_name"`
	RoleId                     int16     `json:"role_id"`
	IsActive                   bool      `json:"is_active"`
	IsBlocked                  bool      `json:"is_blocked"`
	IsDeleted                  bool      `json:"is_deleted"`
	TryLogin                   int16     `json:"try_login"`
	LastSucceedLoginDate       time.Time `json:"last_succeed_login_date"`
	RefreshTokenExpirationDate time.Time `json:"refresh_token_expiration_date"`
	RefreshToken               string    `json:"refresh_token"`
	TryRefresh                 int16     `json:"try_refresh"`
	LastFailedLoginDate        time.Time `json:"last_failed_login_date"`
	LastFailedRefreshDate      time.Time `json:"last_failed_refresh_date"`
	CreateBy                   int64     `json:"create_by"`
	CreateDate                 time.Time `json:"create_date"`
	UpdateBy                   int64     `json:"update_by"`
	UpdateDate                 time.Time `json:"update_date"`
	VerificationDate           time.Time `json:"verification_date"`
	IsVerified                 bool      `json:"is_verified"`
	Code                       string    `json:"code"`
}
