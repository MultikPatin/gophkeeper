package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"net/url"
)

// Config stores all the necessary configurations from both environment variables and command line inputs.
type Config struct {
	DatabaseDSN  *url.URL // Database connection details (Data Source Name).
	DatabaseType string   // Database type (e.g. "postgres", "mysql", etc.).
	JWTSecret    string   // Secret key for JWT authentication.
	GRPCPort     string   // gRPC server port.
}

// envConfig holds configuration settings retrieved from environment variables.
type envConfig struct {
	DatabaseDSN  string `env:"DATABASE_DSN"`  // PostgresSQL Data Source Name received from an environment variable.
	DatabaseType string `env:"DATABASE_TYPE"` // Database type (e.g. "postgres", "mysql", etc.).
	JWTSecret    string `env:"JWT_SECRET"`    // Secret key for JWT authentication.
	GRPCPort     string `env:"GRPC_PORT"`     // gRPC server port.
}

// Parse merges environment variables and command-line options into a single configuration object.
func Parse(logger *zap.SugaredLogger) *Config {
	cfg := &Config{}

	envCfg, err := parseEnv()
	if err != nil {
		logger.Infow("Error while parsing environment variables", "error", err.Error())
	}
	fmt.Printf("envCfg: %+v\n", envCfg)

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
