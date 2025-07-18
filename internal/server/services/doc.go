// Package services implements business logic for handling user and data operations in the server.
// It provides a layer between the application's API and the underlying data storage,
// encapsulating operations such as registration, login, and secure data management.
//
// Key components:
//
//   - UsersService: Manages user registration and authentication using a UserRepository
//     and a PassCryptoService for password hashing.
//   - PasswordsService: Handles password data, encrypting and decrypting sensitive fields
//     with the help of a CryptoService.
//   - CardsService: Manages credit card data, securing sensitive fields through encryption.
//   - BinariesService: Stores and retrieves binary data, ensuring confidentiality via encryption.
//
// All services depend on repositories and crypto services defined in the interfaces package,
// which allows for easy mocking and unit testing.
package services
