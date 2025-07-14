package auth

import "github.com/golang-jwt/jwt/v4"

// Claims extends the standard JWT claims with a custom user ID field.
type Claims struct {
	jwt.RegisteredClaims       // Embedding the standard JWT claims.
	UserID               int64 `json:"userId"` // Custom claim carrying the user ID.

}
