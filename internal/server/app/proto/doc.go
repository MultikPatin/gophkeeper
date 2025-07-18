// Package proto contains gRPC server setup and protocol buffer service handlers for the application.
// It is responsible for binding business logic services to gRPC endpoints and applying interceptors.
//
// Key features:
// - Integration with JWT-based authentication via the auth package.
// - Registration of service handlers for Users, Passwords, Cards, and Binaries.
// - Configuration of logging and authentication interceptors using the interceptors subpackage.
// - Dependency on the "Services" struct from the main application to access business logic.
package proto
