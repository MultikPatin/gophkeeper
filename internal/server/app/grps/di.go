package grps

import (
	"errors"
	"go.uber.org/zap"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/app/grps/handlers"
	"main/internal/server/config"
	"main/internal/server/interfaces"
	"main/internal/server/services"
)

// Handlers organizes gRPC handlers into a coherent structure.
type Handlers struct {
	binaries  interfaces.BinariesHandlers
	passwords interfaces.PasswordsHandlers
	cards     interfaces.CardsHandlers
	users     interfaces.UsersHandlers
	s         *Services
}

// NewHandlers builds a set of gRPC handlers from the provided services.
func NewHandlers(s *Services) *Handlers {
	return &Handlers{
		binaries:  handlers.NewBinariesHandler(s.binaries),
		passwords: handlers.NewPasswordsHandler(s.passwords),
		cards:     handlers.NewCardsHandler(s.cards),
		users:     handlers.NewUsersHandler(s.users),
	}
}

func (h *Handlers) Close() error {
	err := h.s.Close()
	if err != nil {
		return err
	}
	return nil
}

type Services struct {
	binaries  interfaces.BinariesService
	passwords interfaces.PasswordsService
	cards     interfaces.CardsService
	users     interfaces.UsersService
	r         *Repositories
}

func NewServices(c *config.Config, l *zap.SugaredLogger) (*Services, error) {
	r, err := NewRepositories(c, l)
	if err != nil {
		return nil, err
	}

	return &Services{
		binaries:  services.NewBinariesService(r.binaries),
		passwords: services.NewPasswordsService(r.passwords),
		cards:     services.NewCardsService(r.cards),
		users:     services.NewUsersService(r.users),
		r:         r,
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
	binaries  interfaces.BinariesRepository
	passwords interfaces.PasswordsRepository
	cards     interfaces.CardsRepository
	users     interfaces.UsersRepository
	db        interfaces.DB
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
	err = db.Migrate()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		binaries:  psql.NewBinariesRepository(db),
		passwords: psql.NewPasswordsRepository(db),
		cards:     psql.NewCardsRepository(db),
		users:     psql.NewUsersRepository(db),
		db:        db,
	}, nil
}
