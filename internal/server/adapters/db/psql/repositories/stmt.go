package repositories

// Statements is a collection of structured SQL queries for various database operations.
//
// Each nested structure represents a group of related SQL queries for specific entity types (e.g., users, passwords).
// This structure allows convenient organization of queries and easy access by names.
var stmt = statements{
	user: user{
		register: registerUser,
		login:    loginUser,
	},
	binary: binaries{
		add:    addBinary,
		get:    getBinary,
		delete: deleteBinary,
		update: updateBinary,
	},
	card: cards{
		add:    addCard,
		get:    getCard,
		delete: deleteCard,
		update: updateCard,
	},
	password: passwords{
		add:    addPassword,
		get:    getPassword,
		delete: deletePassword,
		update: updatePassword,
	},
}

// statements describes the storage structure of SQL queries.
type statements struct {
	user     user      // Queries for managing users
	binary   binaries  // Queries for working with binary files
	card     cards     // Queries for working with credit cards
	password passwords // Queries for working with stored passwords
}

// user holds SQL queries for CRUD operations on users.
type user struct {
	register string // Register new user
	login    string // Authenticate user
}

// binaries stores SQL queries for working with binary objects.
type binaries struct {
	add    string // Add new binary file
	get    string // Retrieve binary file
	delete string // Delete binary file
	update string // Update binary file content
}

// cards contains SQL queries for working with user's credit cards.
type cards struct {
	add    string // Add new credit card
	get    string // Get credit card details
	delete string // Remove credit card record
	update string // Update credit card information
}

// passwords stores SQL queries for working with saved passwords.
type passwords struct {
	add    string // Save new password entry
	get    string // Fetch existing password entry
	delete string // Delete password entry
	update string // Modify password entry
}

// Constants containing predefined SQL queries.
const (
	// Users
	registerUser = `
        INSERT INTO users (login, password)
        VALUES ($1, $2) 
        RETURNING id;` // Insert new user and return its ID

	loginUser = `
        SELECT id, login, password
        FROM users 
        WHERE login = $1;` // Verify user credentials by username

	// Passwords
	addPassword = `
            INSERT INTO passwords (title, user_id, login, password)
            VALUES ($1, $2, $3, $4) 
            RETURNING title` // Store new password entry and return its title

	getPassword = `
            SELECT id, title, user_id, login, password
            FROM passwords 
            WHERE title = $1 AND user_id = $2` // Find password entry by title and user ID

	deletePassword = `
            DELETE 
            FROM passwords 
            WHERE title = $1 AND user_id = $2` // Remove password entry by title and user ID

	updatePassword = `
            UPDATE passwords 
            SET login = $1, password = $2 
            WHERE title = $3 AND user_id = $4
            RETURNING title` // Update login/password fields in an existing entry

	// Binary Files
	addBinary = `
            INSERT INTO binaries (title, user_id, data) 
            VALUES ($1, $2, $3) 
            RETURNING title` // Create new binary object with associated owner

	getBinary = `
            SELECT id, title, user_id, data
            FROM binaries 
            WHERE title = $1 AND user_id = $2` // Fetch binary object by title and owner

	deleteBinary = `
            DELETE 
            FROM binaries 
            WHERE title = $1 AND user_id = $2` // Remove binary object by title and owner

	updateBinary = `
            UPDATE binaries 
            SET data = $1 
            WHERE title = $2 AND user_id = $3
            RETURNING title` // Update binary object's content

	// Credit Cards
	addCard = `
            INSERT INTO cards (title, user_id, bank, number, data_end, secret_code) 
            VALUES ($1, $2, $3, $4, $5, $6) 
            RETURNING title` // Store new credit card details

	getCard = `
            SELECT id, title, user_id, bank, number, data_end, secret_code 
            FROM cards 
            WHERE title = $1 AND user_id = $2` // Retrieve credit card info by title and user ID

	deleteCard = `
            DELETE 
            FROM cards 
            WHERE title = $1 AND user_id = $2` // Delete credit card record by title and user ID

	updateCard = `
            UPDATE cards 
            SET bank = $1, number = $2, data_end = $3, secret_code = $4  
            WHERE title = $5 AND user_id = $6
            RETURNING title` // Update credit card details by title and user ID
)
