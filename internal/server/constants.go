package server

import "time"

// Type for storing the key of a user's ID in request contexts
type userIDKey string

// Content types and other common constants.
const (
	// ServerShutdownTime specifies the grace period for shutting down the server.
	ServerShutdownTime = 5 * time.Second

	// CertFile is the name of the certificate file used for TLS/SSL.
	CertFile = "cert.pem"

	// KeyFile is the name of the private key file used for TLS/SSL.
	KeyFile = "key.pem"

	// TokenExp defines the expiration duration of a JWT token (3 hours).
	// Used to set the expiration period for authentication tokens.
	TokenExp = time.Hour * 3

	// CookieMaxAge sets the maximum lifetime of cookies (3600 seconds), which equals one hour.
	CookieMaxAge = 3600

	// JwtSecret holds the secret key used for signing JWT tokens.
	// It is important to keep this key secure and avoid exposing it.
	JwtSecret = "your_secret_key"
)
