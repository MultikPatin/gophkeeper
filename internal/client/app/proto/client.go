package proto

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "main/proto"
)

type GothKeeperClient struct {
	conn      *grpc.ClientConn
	Tokens    map[string]string
	Users     pb.UsersClient
	Passwords pb.PasswordsClient
	Cards     pb.CardsClient
	Binaries  pb.BinariesClient
}

func NewGothKeeperClient(GRPCAddr string) (*GothKeeperClient, error) {
	conn, err := grpc.NewClient(GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	userTokens := make(map[string]string, 1)

	return &GothKeeperClient{
		conn:      conn,
		Tokens:    userTokens,
		Users:     pb.NewUsersClient(conn),
		Passwords: pb.NewPasswordsClient(conn),
		Cards:     pb.NewCardsClient(conn),
		Binaries:  pb.NewBinariesClient(conn),
	}, nil
}

func (c *GothKeeperClient) Close() error {
	return c.conn.Close()
}
