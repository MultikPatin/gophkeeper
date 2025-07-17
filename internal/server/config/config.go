package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"net/url"
	"strconv"
	"time"
)

type DatabaseType string

const (
	DefaultJWTExpiration = time.Hour * 3
	DefaultDatabaseType  = "postgres"
	DefaultPostgresDNS   = "postgres://postgres:postgres@localhost:5432/gophkeeper"
	DefaultGRPCPort      = "5050"

	PostgresSQL DatabaseType = "postgres"
)

// Config stores all the necessary configurations from both environment variables and command line inputs.
type Config struct {
	DatabaseDSN   *url.URL      // Database connection details (Data Source Name).
	DatabaseType  string        // Database type (e.g. "postgres", "mysql", etc.).
	JWTSecret     string        // Secret key for JWT authentication.
	JWTExpiration time.Duration // JWT expiration time.
	CryptoSecret  string        // Secret key for cryptographic operations.
	GRPCPort      string        // gRPC server port.
}

// envConfig holds configuration settings retrieved from environment variables.
type envConfig struct {
	DatabaseDSN   string `env:"DATABASE_DSN"`   // PostgresSQL Data Source Name received from an environment variable.
	DatabaseType  string `env:"DATABASE_TYPE"`  // Database type (e.g. "postgres", "mysql", etc.).
	JWTSecret     string `env:"JWT_SECRET"`     // Secret key for JWT authentication.
	JWTExpiration string `env:"JWT_EXPIRATION"` // JWT expiration time in hours (1..24).
	CryptoSecret  string `env:"CRYPTO_SECRET"`  // Secret key for cryptographic operations.
	GRPCPort      string `env:"GRPC_PORT"`      // gRPC server port.
}

// Parse merges environment variables and command-line options into a single configuration object.
func Parse(logger *zap.SugaredLogger) *Config {
	cfg := &Config{}

	envCfg, err := parseEnv()
	if err != nil {
		logger.Infow("Error while parsing environment variables", "error", err.Error())
	}
	fmt.Printf("envCfg: %+v\n", envCfg)

	dbType, err := parseDatabaseType(envCfg.DatabaseType)
	if err != nil {
		logger.Infow("Invalid database type", "error", err.Error())
		logger.Infow("Using default database:", "type", DefaultDatabaseType)
		cfg.DatabaseType = string(PostgresSQL)
	}
	cfg.DatabaseType = string(dbType)

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

	if err := ValidatePort(envCfg.GRPCPort); err != nil {
		logger.Infow("Invalid GRPC port", "error", err.Error())
		logger.Infow("Using default GRPC:", "port", DefaultGRPCPort)
		cfg.GRPCPort = DefaultGRPCPort
	} else {
		cfg.GRPCPort = envCfg.GRPCPort
	}

	if envCfg.CryptoSecret == "" {
		logger.Infow("Crypto secret is empty")
		logger.Infow("Using default crypto secret:", "secret", "secret")
		cfg.CryptoSecret = "secret"
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

// parseEnv extracts configuration from environment variables.
func parseEnv() (*envConfig, error) {
	cfg := &envConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// parseDSN converts a raw Data Source Name (DSN) string into a structured URL object.
func parseDSN(dsn string) (*url.URL, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DSN: %w", err)
	}
	return u, nil
}

func parseDatabaseType(s string) (DatabaseType, error) {
	switch s {
	case string(PostgresSQL):
		return PostgresSQL, nil
	default:
		return "", fmt.Errorf("unsupported database type: %s", s)
	}
}

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

func IsNumberInRange(s string, min int, max int) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil || num < min || num > max {
		return -1, fmt.Errorf("invalid number: %s", s)
	}
	return num, nil
}
