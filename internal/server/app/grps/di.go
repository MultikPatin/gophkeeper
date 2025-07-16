package grps

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"main/internal/server"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/config"
	"main/internal/server/interfaces"
	"main/internal/server/services"
	"net"
	"sync"
)

// App encapsulates the core application state and dependencies.
type App struct {
	services *Services          // Business logic and service instances.
	srv      *grpc.Server       // gRPC servers
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

	srv, err := NewServer(s)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	app := &App{
		services: s,
		log:      l,
		conf:     c,
		srv:      srv,
		ctx:      ctx,
		cancel:   cancel,
	}
	return app, nil
}

// StartServer launches a secondary server dedicated to performance profiling tools.
func (a *App) StartServer() error {
	defer a.wg.Done()

	a.log.Infow("Starting gRPC server", "addr", a.conf.GRPCPort)

	listen, err := net.Listen("tcp", a.conf.GRPCPort)
	if err != nil {
		return err
	}

	errCh := make(chan error)
	go func() {
		if err := a.srv.Serve(listen); err != nil {
			errCh <- fmt.Errorf("ListenAndServe gRPC failed: %w", err)
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		log.Println("Error in gRPC server:", err)
	case <-a.ctx.Done():
		_, cancelShutdown := context.WithTimeout(context.Background(), server.ServerShutdownTime)
		defer cancelShutdown()
		a.srv.GracefulStop()
		if err := listen.Close(); err != nil {
			a.log.Fatalw(err.Error(), "event", "TCP listen shutdown")
		}
	}
	return nil
}

// Close gracefully cleans up running services and dependencies.
func (a *App) Close() error {
	a.cancel()
	a.wg.Wait()
	a.srv.Stop()
	err := a.services.Close()
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
