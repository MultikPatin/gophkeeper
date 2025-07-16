package psql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"
)

// DB encapsulates a SQL database connection for interacting with a PostgresSQL backend.
type DB struct {
	Conn *sql.DB // The active database connection.
}

// NewDB establishes a new PostgresSQL database connection using provided credentials.
func NewDB(dsn *url.URL) (*DB, error) {
	host := dsn.Hostname()
	port := dsn.Port()
	user := dsn.User.Username()
	password, _ := dsn.User.Password()
	dbname := dsn.Path[1:]

	ps := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("pgx", ps)
	if err != nil {
		return nil, err
	}

	return &DB{
		Conn: db,
	}, nil
}

// Close terminates the database connection cleanly.
func (p *DB) Close() error {
	err := p.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}

// Ping verifies connectivity to the database by issuing a ping request.
func (p *DB) Ping() error {
	err := p.Conn.Ping()
	return err
}

// Migrate applies database schema migrations using the provided context and connection.
func (p *DB) Migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := p.Conn.ExecContext(ctx, createTables)
	if err != nil {
		return err
	}
	return nil
}
