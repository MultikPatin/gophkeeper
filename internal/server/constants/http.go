package constants

import "time"

// Content types and other common constants.
const (
	// ServerShutdownTime specifies the grace period for shutting down the server.
	ServerShutdownTime = 5 * time.Second

	// TextContentType is the MIME type for plain text content encoded in UTF-8.
	TextContentType = "text/plain; charset=utf-8"

	// JSONContentType is the MIME type for JSON-formatted data.
	JSONContentType = "application/json"

	// CertFile is the name of the certificate file used for TLS/SSL.
	CertFile = "cert.pem"

	// KeyFile is the name of the private key file used for TLS/SSL.
	KeyFile = "key.pem"
)

// IgnoreURLs is a list of URLs that should be ignored by the middleware.
var IgnoreURLs = []string{"/register", "/login"}
