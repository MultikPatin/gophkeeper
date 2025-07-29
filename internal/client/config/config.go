package config

import (
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
)

// DefaultGRPCAddr Predefined constants for default values across different configurations.
const (
	DefaultGRPCAddr = "127.0.0.1:5050" // Default gRPC server listening.
)

// Config encapsulates application-wide configuration parameters derived from environment variables and command-line arguments.
type Config struct {
	GRPCAddr string // Port where the gRPC server.
}

// envConfig captures configuration properties extracted directly from environment variables.
type envConfig struct {
	GRPCAddr string `env:"GRPC_SERVER_ADDRESS"` // Environment variable defining the gRPC server.
}

// Parse consolidates configuration from multiple sources like environment variables and command-line flags into a unified Config object.
func Parse(logger *zap.SugaredLogger) *Config {
	cfg := &Config{}

	envCfg, err := parseEnv()
	if err != nil {
		logger.Infow("Error while parsing environment variables", "error", err.Error())
	}

	if envCfg.GRPCAddr == "" {
		logger.Infow("Invalid GRPC addr", "error", err.Error())
		logger.Infow("Using default GRPC:", "addr", DefaultGRPCAddr)
		cfg.GRPCAddr = DefaultGRPCAddr
	} else {
		cfg.GRPCAddr = envCfg.GRPCAddr
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
