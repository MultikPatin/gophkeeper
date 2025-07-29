package psql

const createTables = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		login VARCHAR(64) UNIQUE NOT NULL,
		password TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS passwords (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) UNIQUE NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		login BYTEA NOT NULL,
		password BYTEA NOT NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS passwords_user_id_title_idx 
	ON passwords (user_id, title);

	CREATE TABLE IF NOT EXISTS cards (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) UNIQUE NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		bank BYTEA NOT NULL,
		number BYTEA NOT NULL,
		data_end BYTEA NOT NULL,
		secret_code BYTEA NOT NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS cards_user_id_title_idx 
	ON cards (user_id, title);

	CREATE TABLE IF NOT EXISTS binaries (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) UNIQUE NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		data BYTEA NOT NULL
	);
	CREATE UNIQUE INDEX IF NOT EXISTS binaries_user_id_title_idx 
	ON binaries (user_id, title);
`
