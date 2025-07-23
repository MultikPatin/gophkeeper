// Package config provides functionality for parsing and managing the application configuration.
// It loads settings from environment variables and applies default values when necessary.
//
// The main configuration structure is:
//
//	type Config struct {
//	    GRPCAddr      string        // Port for the gRPC server.
//	}
package config
