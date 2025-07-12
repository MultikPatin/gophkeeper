package server

import (
	"errors"
	"go.uber.org/zap"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/config"
	"main/internal/server/interfaces"
	"main/internal/server/services"
)

type Services struct {
	binaries interfaces.BinariesService
	cards    interfaces.CardsService
	notes    interfaces.NotesService
	users    interfaces.UsersService
	r        *Repositories
}

func NewServices(c *config.Config, l *zap.SugaredLogger) (*Services, error) {
	r, err := NewRepositories(c, l)
	if err != nil {
		return nil, err
	}

	return &Services{
		binaries: services.NewBinariesService(r.binaries),
		cards:    services.NewCardsService(r.cards),
		notes:    services.NewNotesService(r.notes),
		users:    services.NewUsersService(r.users),
		r:        r,
	}, nil
}

func (s *Services) Close() error {
	err := s.r.Close()
	if err != nil {
		return err
	}
	return nil
}

type Repositories struct {
	binaries interfaces.BinariesRepository
	cards    interfaces.CardsRepository
	notes    interfaces.NotesRepository
	users    interfaces.UsersRepository
	db       interfaces.DB
}

func NewRepositories(c *config.Config, l *zap.SugaredLogger) (*Repositories, error) {
	switch c.DatabaseType {
	case "postgres":
		repositories, err := postgresRepositories(c, l)
		if err != nil {
			return nil, err
		}
		return repositories, nil
	default:
		return nil, errors.New("unsupported database type")
	}

}

func (r *Repositories) Close() error {
	err := r.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func postgresRepositories(c *config.Config, l *zap.SugaredLogger) (*Repositories, error) {
	db, err := psql.NewDB(c.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(c.MigrationDir)
	if err != nil {
		return nil, err
	}

	return &Repositories{
		binaries: psql.NewBinariesRepository(db),
		cards:    psql.NewCardsRepository(db),
		notes:    psql.NewNotesRepository(db),
		users:    psql.NewUsersRepository(db),
		db:       db,
	}, nil
}
