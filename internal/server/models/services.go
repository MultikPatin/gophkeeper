package models

// User represents a user entity with unique identification, login, and password attributes.
type User struct {
	ID       int64  // Unique identifier for the user.
	Login    string // Username or email address for logging in.
	Password string // Hashed password for authentication.
}

// Password stores password details associated with a particular user.
// Fields such as login and password are stored in encrypted format.
type Password struct {
	ID       int64  // Unique identifier for this password entry.
	Title    string // Title or label describing the password usage.
	UserID   int64  // Foreign key linking to the owning user.
	Login    []byte // Encrypted login credential.
	Password []byte // Encrypted password itself.
}

// Card encapsulates credit/debit card information, ensuring sensitive data remains encrypted.
type Card struct {
	ID         int64  // Unique identifier for this card entry.
	Title      string // Descriptive title for identifying the card.
	UserID     int64  // Foreign key pointing to the associated user.
	Bank       []byte // Encrypted bank name.
	Number     []byte // Encrypted card number.
	DataEnd    []byte // Encrypted expiry date.
	SecretCode []byte // Encrypted CVV code.
}

// BinaryData represents generic binary blobs attached to users.
// Useful for storing files, images, or other forms of binary data.
type BinaryData struct {
	ID     int64  // Unique identifier for this binary data entry.
	Title  string // Label or description for the binary data.
	UserID int64  // Foreign key referencing the owning user.
	Data   []byte // Raw binary content.
}
