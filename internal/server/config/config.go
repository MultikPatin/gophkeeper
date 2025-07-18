package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"time"
)

// DatabaseType specifies the supported database types as strings.
type DatabaseType string

// Predefined constants for default values across different configurations.
const (
	DefaultCryptoKey     = "3a7d4e1f9c02b58e7d9a2f3e8b01c9d7"                         // Default crypto key
	DefaultJWTExpiration = time.Hour * 3                                              // Default JWT token expiration set to 3 hours.
	DefaultPostgresDNS   = "postgresql://postgres:postgres@localhost:5432/gophkeeper" // Default PostgreSQL Data Source Name (DSN).
	DefaultGRPCAddr      = "127.0.0.1:5050"                                           // Default gRPC server listening.

	PostgresSQL DatabaseType = "postgres" // Supported database type constant.
)

// Config encapsulates application-wide configuration parameters derived from environment variables and command-line arguments.
type Config struct {
	DatabaseDSN   *url.URL      // Parsed Data Source Name (DSN) for connecting to the database.
	DatabaseType  string        // Type of the database being used ("postgres", etc.).
	JWTSecret     string        // Secret key used for signing JWT tokens.
	JWTExpiration time.Duration // Expiration duration for issued JWT tokens.
	CryptoSecret  string        // Key used for encryption purposes.
	GRPCAddr      string        // Port where the gRPC server.
}

// envConfig captures configuration properties extracted directly from environment variables.
type envConfig struct {
	DatabaseDSN   string `env:"DATABASE_DSN"`        // Environment variable holding the DSN for the database.
	DatabaseType  string `env:"DATABASE_TYPE"`       // Environment variable specifying the database type.
	JWTSecret     string `env:"JWT_SECRET"`          // Environment variable storing the JWT secret key.
	JWTExpiration string `env:"JWT_EXPIRATION"`      // Environment variable controlling JWT token lifetime.
	CryptoSecret  string `env:"CRYPTO_SECRET"`       // Environment variable providing the cryptography secret.
	GRPCAddr      string `env:"GRPC_SERVER_ADDRESS"` // Environment variable defining the gRPC server.
}

// Parse consolidates configuration from multiple sources like environment variables and command-line flags into a unified Config object.
func Parse(logger *zap.SugaredLogger) *Config {
	cfg := &Config{}

	envCfg, err := parseEnv()
	if err != nil {
		logger.Infow("Error while parsing environment variables", "error", err.Error())
	}

	dbType, err := parseDatabaseType(envCfg.DatabaseType)
	if err != nil {
		logger.Infow("Invalid database type", "error", err.Error())
		logger.Infow("Using default database:", "type", string(PostgresSQL))
		cfg.DatabaseType = string(PostgresSQL)
	} else {
		cfg.DatabaseType = string(dbType)
	}

	if envCfg.DatabaseDSN == "" {
		logger.Infow("Database DSN is empty")
		logger.Infow("Using default database:", "DSN", DefaultPostgresDNS)
		cfg.DatabaseDSN, _ = parseDSN(DefaultPostgresDNS)
	} else {
		cfg.DatabaseDSN, err = parseDSN(envCfg.DatabaseDSN)
		if err != nil {
			logger.Infow("Error while parsing database DSN", "error", err.Error())
			logger.Infow("Using default database:", "DSN", DefaultPostgresDNS)
			switch cfg.DatabaseType {
			case string(PostgresSQL):
				cfg.DatabaseDSN, _ = parseDSN(DefaultPostgresDNS)
			default:
				cfg.DatabaseDSN, _ = parseDSN(DefaultPostgresDNS)
			}
		}
	}

	if envCfg.GRPCAddr == "" {
		logger.Infow("Invalid GRPC addr", "error", err.Error())
		logger.Infow("Using default GRPC:", "addr", DefaultGRPCAddr)
		cfg.GRPCAddr = DefaultGRPCAddr
	} else {
		cfg.GRPCAddr = envCfg.GRPCAddr
	}

	if envCfg.CryptoSecret == "" {
		logger.Infow("Crypto secret is empty")
		logger.Infow("Using default crypto secret:", "secret", DefaultCryptoKey)
		cfg.CryptoSecret = DefaultCryptoKey
	} else {
		cfg.CryptoSecret = envCfg.CryptoSecret
	}

	num, err := IsNumberInRange(envCfg.JWTExpiration, 1, 99)
	if err != nil {
		logger.Infow("Invalid JWT expiration", "error", err.Error())
		logger.Infow("Using default JWTExpiration:", "expiration", DefaultJWTExpiration)
		cfg.JWTExpiration = DefaultJWTExpiration
	} else {
		cfg.JWTExpiration = time.Hour * time.Duration(num)
	}

	return cfg
}

// parseEnv extracts configuration settings from environment variables.
func parseEnv() (*envConfig, error) {
	cfg := &envConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// parseDSN parses a raw Data Source Name (DSN) string into a structured URL representation.
func parseDSN(dsn string) (*url.URL, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}
	return u, nil
}

// parseDatabaseType translates a string representation of the database type into its corresponding DatabaseType value.
func parseDatabaseType(s string) (DatabaseType, error) {
	switch s {
	case string(PostgresSQL):
		return PostgresSQL, nil
	default:
		return "", fmt.Errorf("unsupported database type: %s", s)
	}
}

// ValidatePort checks whether the supplied port is valid according to basic constraints.
func ValidatePort(port string) error {
	if port == "" {
		return fmt.Errorf("GRPC port is empty")
	}
	if port[0] == ':' {
		port = port[1:]
	}
	if port == "0" {
		return fmt.Errorf("GRPC port cannot be 0")
	}
	if len(port) > 5 {
		return fmt.Errorf("GRPC port is too long")
	}
	for _, r := range port {
		if r < '0' || r > '9' {
			return fmt.Errorf("GRPC port contains non-digit characters")
		}
	}
	if p, _ := strconv.Atoi(port); p < 1 || p > 65535 {
		return fmt.Errorf("GRPC port is out of valid range (1-65535)")
	}
	return nil
}

// IsNumberInRange ensures that a numeric string falls within a defined minimum and maximum boundary.
func IsNumberInRange(s string, min int, max int) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil || num < min || num > max {
		return -1, fmt.Errorf("invalid number: %s", s)
	}
	return num, nil
}
