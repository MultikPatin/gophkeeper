// Package auth provides authentication functionality for the server.
// It includes JWT token generation and verification, as well as user credential validation.
// The main components are:
// - Claims: Custom JWT claims structure with embedded RegisteredClaims and a UserID field
// - JWTService: Service for generating and verifying JWT tokens
// - Error handling for unexpected signing methods and invalid tokens
package auth
