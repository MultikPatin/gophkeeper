package config

import (
	"net/url"
)

//const (
//	defaultStorageFilePath = "shorter"        // Default path for storage file if no custom path is provided.
//	defaultPProfAddr       = "localhost:6060" // Address for pprof profiling endpoint.
//	defaultConfFileName    = "conf.json"      // Name of the configuration file in json format
//)

// Config stores all the necessary configurations from both environment variables and command line inputs.
type Config struct {
	DatabaseDSN  *url.URL // Database connection details (Data Source Name).
	DatabaseType string   // Database type (e.g. "postgres", "mysql", etc.).
	MigrationDir string   // Directory where migration files are located.
	Addr         string   // Server listening address.
	JWTSecret    string   // Secret key for JWT authentication.
	//GRPCAddr         string   // gRPC server address.
	//ShortLinkPrefix  string   // Base URL for short links.
	//StorageFilePaths string   // Path where storage files are located.
	//ExecutableDir    string   // Project directory
	//HTTPSEnable      bool     // Indicates whether HTTPS is enabled for the server.
	//TrustedSubnet    string   // Trusted subnet for the server.
}
