package proto

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "main/proto"
)

// GothKeeperClient represents a client wrapper for interacting with GRPC services.
// Provides access to different servers through one unified interface.
type GothKeeperClient struct {
	conn      *grpc.ClientConn   // Connection to GRPC server
	Token     string             // Authorization token
	Users     pb.UsersClient     // Client for users operations
	Passwords pb.PasswordsClient // Client for passwords operations
	Cards     pb.CardsClient     // Client for cards operations
	Binaries  pb.BinariesClient  // Client for binaries operations
}

// NewGothKeeperClient creates a new connection to a GRPC server and initializes corresponding clients.
// Parameters:
// - GRPCAddr: Address of the GRPC server.
// - token: Authentication token.
// Returns created client and possible error.
// Potential errors:
// - FailedConnection: Unable to establish a connection to the GRPC server.
// - InvalidArguments: Incorrect arguments passed to the constructor.
func NewGothKeeperClient(GRPCAddr string) (*GothKeeperClient, error) {
	conn, err := grpc.NewClient(GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GothKeeperClient{
		conn:      conn,
		Users:     pb.NewUsersClient(conn),
		Passwords: pb.NewPasswordsClient(conn),
		Cards:     pb.NewCardsClient(conn),
		Binaries:  pb.NewBinariesClient(conn),
	}, nil
}

// Close closes the active connection to the GRPC server.
// Used for resource cleanup.
// Potential errors:
// - ConnectionClosedError: An error occurred while closing the connection.
func (c *GothKeeperClient) Close() error {
	return c.conn.Close()
}
