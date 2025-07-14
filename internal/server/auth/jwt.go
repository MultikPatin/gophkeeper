package auth

import (
	"errors"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

// VerifyJWT validates a JWT token and extracts the user ID claim.
func VerifyJWT(tokenStr string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GenerateJWT issues a signed JWT token with a specified user ID and expiry.
func GenerateJWT(userID int64, secret string, tokenExp time.Duration) (string, error) {
	expirationTime := time.Now().Add(tokenExp)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
