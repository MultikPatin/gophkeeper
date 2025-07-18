package proto

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"main/internal/server/adapters/db/psql"
	"main/internal/server/adapters/db/psql/repositories"
	"main/internal/server/config"
	"main/internal/server/crypto"
	"main/internal/server/interfaces"
	"main/internal/server/services"
	"net"
	"sync"
	"time"
)

const (
	// ShutdownTime specifies the grace period for shutting down the server.
	ShutdownTime = 5 * time.Second
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

	srv, err := NewServer(s, c, l)
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

	a.log.Infow("Starting gRPC server", "addr", a.conf.GRPCAddr)

	listen, err := net.Listen("tcp", a.conf.GRPCAddr)
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
		_, cancelShutdown := context.WithTimeout(context.Background(), ShutdownTime)
		defer cancelShutdown()
		a.srv.GracefulStop()
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

	aesCrypto, err := crypto.NewAes([]byte(c.CryptoSecret))
	if err != nil {
		return nil, err
	}
	passCrypto := crypto.NewPassCrypto()

	return &Services{
		binaries:  services.NewBinariesService(r.binaries, aesCrypto),
		passwords: services.NewPasswordsService(r.passwords, aesCrypto),
		cards:     services.NewCardsService(r.cards, aesCrypto),
		users:     services.NewUsersService(r.users, passCrypto),
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
		r, err := postgresRepositories(c, l)
		if err != nil {
			return nil, err
		}
		return r, nil
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
		binaries:  repositories.NewBinariesRepository(db),
		passwords: repositories.NewPasswordsRepository(db),
		cards:     repositories.NewCardsRepository(db),
		users:     repositories.NewUsersRepository(db),
		db:        db,
	}, nil
}
