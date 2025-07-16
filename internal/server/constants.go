package server

import "time"

// Type for storing the key of a user's ID in request contexts
type userIDKey string

// Content types and other common constants.
const (
	// ShutdownTime specifies the grace period for shutting down the server.
	ShutdownTime = 5 * time.Second

	// TokenExp defines the expiration duration of a JWT token (3 hours).
	// Used to set the expiration period for authentication tokens.
	TokenExp = time.Hour * 3

	// CookieMaxAge sets the maximum lifetime of cookies (3600 seconds), which equals one hour.
	CookieMaxAge = 3600

	// JwtSecret holds the secret key used for signing JWT tokens.
	// It is important to keep this key secure and avoid exposing it.
	JwtSecret = "your_secret_key"
)
