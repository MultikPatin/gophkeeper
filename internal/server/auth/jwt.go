package auth

import (
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

// Errors that can be returned during token validation or generation.
var (
	ErrUnexpectedMethod = errors.New("unexpected signing method") // Error when unexpected signing method is detected
	ErrInvalidToken     = errors.New("invalid token")             // Error when token cannot be verified
)

// JWTService provides methods for generating and verifying JSON Web Tokens (JWT).
type JWTService struct {
	secret   string        // Secret key used for signing tokens
	tokenExp time.Duration // Token expiration duration
}

// NewJWTService initializes a new JWT service with the specified secret key and token expiration period.
func NewJWTService(secret string, tokenExp time.Duration) (*JWTService, error) {
	return &JWTService{
		secret:   secret,
		tokenExp: tokenExp,
	}, nil
}

// Verify validates a JWT token and extracts the associated user ID.
// It returns the user ID if valid, otherwise an error.
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

// Generate creates a signed JWT token for the given user ID.
// The generated token will expire after the configured token expiration duration.
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
