package xjwt

import (
	"encoding/json"
	"errors"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/lib/config"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Type      string `json:"type"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
}

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

func GenerateToken(user *entity.User, tokenType TokenType) (string, error) {
	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(user.ID)),
			Issuer:    config.Config.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.Config.Jwt.ExpiredAt) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Type:      string(tokenType),
		UserName:  user.Name,
		UserEmail: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.Jwt.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func MapClaimsToTokenClaims(token *jwt.Token) (*TokenClaims, error) {
	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims type")
	}
	claimsBytes, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, err
	}
	var claims TokenClaims
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return nil, err
	}
	return &claims, nil
}

func ExtractTokenFromCtx(c *fiber.Ctx) *TokenClaims {
	claims, ok := c.Locals("claims").(*TokenClaims)
	if !ok {
		return nil
	}
	return claims
}
