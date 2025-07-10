package psql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	pm "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"net/url"
	"path/filepath"
)

// PostgresDB encapsulates a SQL database connection for interacting with a PostgresSQL backend.
type PostgresDB struct {
	db *sql.DB // The active database connection.
}

// NewPostgresDB establishes a new PostgresSQL database connection using provided credentials.
func NewPostgresDB(dsn *url.URL) (*PostgresDB, error) {
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

	return &PostgresDB{
		db: db,
	}, nil
}

// Close terminates the database connection cleanly.
func (p *PostgresDB) Close() error {
	err := p.db.Close()
	if err != nil {
		return err
	}
	return nil
}

// Ping verifies connectivity to the database by issuing a ping request.
func (p *PostgresDB) Ping() error {
	err := p.db.Ping()
	return err
}

// Migrate applies database schema migrations using the provided context and connection.
func (p *PostgresDB) Migrate(migrationsPath string) error {
	databaseName := "postgres"
	sourceURL := "file://"

	if migrationsPath == "" {
		sourceURL += "db/migrations"
	} else {
		path, err := filepath.Abs(migrationsPath)
		if err != nil {
			return err
		}
		sourceURL += path
	}

	driver, err := pm.WithInstance(p.db, &pm.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
