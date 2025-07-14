package psql

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

type statements struct {
	user     user
	binary   binaries
	card     cards
	password passwords
}

type user struct {
	register string
	login    string
}

type binaries struct {
	add    string
	get    string
	delete string
	update string
}

type cards struct {
	add    string
	get    string
	delete string
	update string
}

type passwords struct {
	add    string
	get    string
	delete string
	update string
}

const (
	// Users
	registerUser = `
		INSERT INTO users (name, password)
		VALUES ($1, $2) 
		RETURNING id;`
	loginUser = `
		SELECT id, name 
		FROM users 
		WHERE name = $1;`
	// Binaries
	addBinary = `
			INSERT INTO binaries (name, data) 
			VALUES ($1, $2) 
			RETURNING id`
	getBinary = `
			SELECT id, name, data 
			FROM binaries 
			WHERE id = $1`
	deleteBinary = `
			DELETE 
			FROM binaries 
			WHERE id = $1`
	updateBinary = `
			UPDATE binaries 
			SET name = $1, data = $2 
			WHERE id = $3`
	// Cards
	addCard = `
			INSERT INTO cards (name, data) 
			VALUES ($1, $2) 
			RETURNING id`
	getCard = `
			SELECT id, name, data 
			FROM cards 
			WHERE id = $1`
	deleteCard = `
			DELETE 
			FROM cards 
			WHERE id = $1`
	updateCard = `
			UPDATE cards 
			SET name = $1, data = $2 
			WHERE id = $3`
	// Passwords
	addPassword = `
			INSERT INTO passwords (name, data) 
			VALUES ($1, $2) 
			RETURNING id`
	getPassword = `
			SELECT id, name, data 
			FROM passwords 
			WHERE id = $1`
	deletePassword = `
			DELETE 
			FROM passwords 
			WHERE id = $1`
	updatePassword = `
			UPDATE passwords 
			SET name = $1, data = $2 
			WHERE id = $3`
)
