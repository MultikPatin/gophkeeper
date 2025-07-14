package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/app/middlewares"
	"main/internal/server/config"
	"main/internal/server/constants"
	"main/internal/server/interfaces"
	"main/internal/server/services"
	"net/http"
	"sync"
)

// App encapsulates the core application state and dependencies.
type App struct {
	Router   *chi.Mux           // Main router for handling HTTP requests.
	Services *Services          // Business logic and service instances.
	log      *zap.SugaredLogger // Configuration settings.
	conf     *config.Config     // Logger for application-wide logging.
	cancel   context.CancelFunc // Function to cancel the application context.
	ctx      context.Context    // Application context for signal propagation.
	wg       sync.WaitGroup     // Wait group for tracking background tasks.
}

// NewApp constructs a fully-configured application instance.
func NewApp(c *config.Config, l *zap.SugaredLogger) (*App, error) {
	s, err := NewServices(c, l)
	if err != nil {
		return nil, err
	}

	authConf := middlewares.AuthParams{
		IgnoreURLs: constants.IgnoreURLs,
		JWTSecret:  c.JWTSecret,
	}

	ctx, cancel := context.WithCancel(context.Background())

	app := &App{
		log:      l,
		conf:     c,
		Router:   NewRouters(NewHandlers(s), authConf),
		Services: s,
		ctx:      ctx,
		cancel:   cancel,
	}
	return app, nil
}

// StartServer boots the primary HTTP server and handles graceful shutdowns.
func (a *App) StartServer() error {
	a.wg.Add(1)

	a.log.Infow("Starting server", "addr", a.conf.Addr)

	srv := &http.Server{
		Addr:    a.conf.Addr,
		Handler: a.Router,
	}
	errCh := make(chan error)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("ListenAndServe failed: %w", err)
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		return err
	case <-a.ctx.Done():
		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), constants.ServerShutdownTime)
		defer cancelShutdown()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			a.log.Fatalw(err.Error(), "event", "server shutdown")
		}
		return nil
	}
}

// Close gracefully cleans up running services and dependencies.
func (a *App) Close() error {
	a.cancel()
	a.wg.Wait()
	err := a.Services.Close()
	if err != nil {
		return err
	}
	return nil
}

// Handlers organizes HTTP handlers into a coherent structure.
type Handlers struct {
}

// NewHandlers builds a set of HTTP handlers from the provided services.
func NewHandlers(s *Services) *Handlers {
	return &Handlers{}
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
	err = db.Migrate(c.MigrationDir)
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
