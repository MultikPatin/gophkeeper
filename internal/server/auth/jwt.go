package auth

import (
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

var (
	ErrUnexpectedMethod = errors.New("unexpected signing method")
	ErrInvalidToken     = errors.New("invalid token")
)

type JWTService struct {
	secret   string
	tokenExp time.Duration
}

func NewJWTService(secret string, tokenExp time.Duration) (*JWTService, error) {
	return &JWTService{
		secret:   secret,
		tokenExp: tokenExp,
	}, nil
}

func (s *JWTService) Verify(tokenStr string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedMethod
		}
		return []byte(s.secret), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return -1, ErrInvalidToken
	}
	return claims.UserID, nil
}

func (s *JWTService) Generate(userID int64) (string, error) {
	expirationTime := time.Now().Add(s.tokenExp)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
