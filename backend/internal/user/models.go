package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User represents someone with access to our system.
type User struct {
	ID           string    `db:"user_id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash []byte    `db:"password_hash" json:"-"`
	DateCreated  time.Time `db:"date_created" json:"date_created"`
	DateUpdated  time.Time `db:"date_updated" json:"date_updated"`
}

// LoginUser is used when a user logs into the system
type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// WebUser is used when a user is returned to the web
type WebUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"admin"`
}

// JwtUserToken provides the structure to return a JWT token for the user
type JwtUserToken struct {
	WebUser
	jwt.StandardClaims
}

// JwtUserTokenResponse delivers the response to the UI, containing the JWT
type JwtUserTokenResponse struct {
	Token string `json:"token"`
}

// NewUser contains information needed to create a new User.
type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
