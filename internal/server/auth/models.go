package auth

import "github.com/golang-jwt/jwt/v4"

// Claims defines extended JWT claims including a custom 'UserID' field.
// This type embeds the standard RegisteredClaims and adds a custom 'UserID'.
type Claims struct {
	jwt.RegisteredClaims       // Standard JWT registered claims embedded here.
	UserID               int64 `json:"userId"` // Custom claim representing the authenticated user's unique identifier.
}
