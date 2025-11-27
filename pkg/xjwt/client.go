package xjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Client interface {
	GenerateToken(req GenerateTokenRequest) (*GenerateTokenResponse, error)
	ValidateToken(req ValidateTokenRequest) (*ValidateTokenResponse, error)
}

type client struct {
	cfg Config
}

func NewClient(config ...Config) Client {
	cfg := setConfig(config...)
	return &client{
		cfg: cfg,
	}
}

// GenerateToken generate access and refresh tokens for the given user
func (c *client) GenerateToken(req GenerateTokenRequest) (*GenerateTokenResponse, error) {
	accessTokenExpiresAt := time.Now().Add(time.Duration(c.cfg.AccessTokenExpiresIn) * time.Second)
	refreshTokenExpiresAt := time.Now().Add(time.Duration(c.cfg.RefreshTokenExpiresIn) * time.Second)
	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   req.User.ID,
			Issuer:    c.cfg.Issuer,
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		TokenType: TokenTypeAccess,
		UserName:  req.User.Name,
		UserEmail: req.User.Email,
	}

	// generate access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(c.cfg.SecretKey))
	if err != nil {
		return nil, err
	}

	// generate refresh token
	claims.ExpiresAt = jwt.NewNumericDate(refreshTokenExpiresAt)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshTokenString, err := refreshToken.SignedString([]byte(c.cfg.SecretKey))
	if err != nil {
		return nil, err
	}

	return &GenerateTokenResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    accessTokenExpiresAt,
	}, nil
}

func (c *client) ValidateToken(req ValidateTokenRequest) (*ValidateTokenResponse, error) {
	claims := &TokenClaims{}

	parsedToken, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(c.cfg.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return &ValidateTokenResponse{
		TokenClaims: claims,
	}, nil
}
