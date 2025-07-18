// Package config provides functionality for parsing and managing the application configuration.
// It loads settings from environment variables and applies default values when necessary.
//
// The main configuration structure is:
//
//	type Config struct {
//	    DatabaseDSN   *url.URL      // Parsed DSN for the database connection.
//	    DatabaseType  string        // Database type, e.g., "postgres".
//	    JWTSecret     string        // Secret key used for signing JWT tokens.
//	    JWTExpiration time.Duration // Token expiration time in hours.
//	    CryptoSecret  string        // Secret used for encryption.
//	    GRPCPort      string        // Port for the gRPC server.
//	}
package config
