package jwtgo

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/sociosarbis/grpc/boilerplate/internal/config"
	"github.com/sociosarbis/grpc/boilerplate/internal/pkg/errgo"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type User struct {
	ID    uint32 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type UserClaims struct {
	jwt.StandardClaims
	User
}

var errUnexpectedTokenMethod = errors.New("unexpected token signing method")

func NewJWTManager(config config.AppConfig) *JWTManager {
	return &JWTManager{
		secretKey:     config.JWTSecretKey,
		tokenDuration: time.Second * time.Duration(config.JWTDuration),
	}
}

func (manager *JWTManager) Generate(user User) (string, error) {
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.tokenDuration).Unix(),
		},
		User: user,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errUnexpectedTokenMethod
		}
		return []byte(manager.secretKey), nil
	})

	if err != nil {
		return nil, errgo.Wrap(err, "invalid token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errgo.Wrap(err, "invalid token claims")
	}
	return claims, nil
}
