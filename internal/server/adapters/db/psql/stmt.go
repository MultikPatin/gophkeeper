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
		INSERT INTO users (login, password)
		VALUES ($1, $2) 
		RETURNING id;`
	loginUser = `
		SELECT id, login, password
		FROM users 
		WHERE login = $1;`
	// Passwords
	addPassword = `
			INSERT INTO passwords (title, user_id, login, password)
			VALUES ($1, $2, $3, $4) 
			RETURNING id`
	getPassword = `
			SELECT id, title, user_id, login, password
			FROM passwords 
			WHERE title = $1 AND user_id = $2`
	deletePassword = `
			DELETE 
			FROM passwords 
			WHERE title = $1 AND user_id = $2`
	updatePassword = `
			UPDATE passwords 
			SET login = $1, password = $2 
			WHERE title = $3 AND user_id = $4
			RETURNING id`
	// Binaries
	addBinary = `
			INSERT INTO binaries (title, user_id, data) 
			VALUES ($1, $2, $3) 
			RETURNING id`
	getBinary = `
			SELECT id, title, user_id, data
			FROM binaries 
			WHERE title = $1 AND user_id = $2`
	deleteBinary = `
			DELETE 
			FROM binaries 
			WHERE title = $1 AND user_id = $2`
	updateBinary = `
			UPDATE binaries 
			SET data = $1 
			WHERE title = $2 AND user_id = $3
			RETURNING id`
	// Cards
	addCard = `
			INSERT INTO cards (title, user_id, bank, number, data_end, secret_code) 
			VALUES ($1, $2, $3, $4, $5, $6) 
			RETURNING id`
	getCard = `
			SELECT id, title, user_id, bank, number, data_end, secret_code 
			FROM cards 
			WHERE title = $1 AND user_id = $2`
	deleteCard = `
			DELETE 
			FROM cards 
			WHERE title = $1 AND user_id = $2`
	updateCard = `
			UPDATE cards 
			SET bank = $1, number = $2, data_end = $3, secret_code = $4  
			WHERE title = $5 AND user_id = $6
			RETURNING id`
)
