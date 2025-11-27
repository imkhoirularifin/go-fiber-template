package xjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	TokenType string `json:"token_type"`
	UserName  string `json:"username"`
	UserEmail string `json:"email"`
}

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type GenerateTokenRequest struct {
	User UserInfo `json:"user"`
}
type UserInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type GenerateTokenResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type ValidateTokenRequest struct {
	Token string `json:"token"`
}

type ValidateTokenResponse struct {
	TokenClaims *TokenClaims `json:"token_claims"`
}
