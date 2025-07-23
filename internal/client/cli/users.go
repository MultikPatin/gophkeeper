package cli

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"main/internal/client/app/proto"
	pb "main/proto"
)

// SetupUserCommand sets up the 'user' command with subcommands for registration and login.
// No error handling is required here as it simply returns a cobra.Command pointer.
func SetupUserCommand(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "User authorization and registration",
		Long: `The section contains methods for user 
		authorization and registration based on JWT tokens`,
	}
	cmd.AddCommand(registerUser(client))
	cmd.AddCommand(logiUser(client))
	return cmd
}

// registerUser creates a new Cobra command to handle user registration.
// It retrieves flags from the CLI input, constructs a RegisterRequest protobuf message,
// sends it to the gRPC server, and handles potential errors including GRPC-specific ones like AlreadyExists.
// If successful, it prints a success message and stores the returned token in the client instance.
func registerUser(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register User",
		Long:  `Register User.`,
		Run: func(cmd *cobra.Command, args []string) {
			username, err := cmd.Flags().GetString("username")
			if err != nil {
				cmd.PrintErr(err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.RegisterRequest{
				Login:    username,
				Password: password,
			}

			result, err := client.Users.Register(cmd.Context(), &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.AlreadyExists:
						cmd.Print("User already exists")
					default:
						cmd.Println("Error:", st.Message())
					}
				} else {
					cmd.PrintErrf("Error: %v", err)
				}
			} else {
				client.Token = result.Token
				cmd.Print("Successfully registered")
			}
		},
	}
	cmd.Flags().StringP("username", "u", "", "Username")
	cmd.Flags().StringP("password", "p", "", "Password")
	err := cmd.MarkFlagRequired("username")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("password")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// logiUser creates a new Cobra command to handle user login.
// It retrieves flags from the CLI input, constructs a LoginRequest protobuf message,
// sends it to the gRPC server, and handles potential errors including GRPC-specific ones like NotFound and Unauthenticated.
// If successful, it prints a success message and stores the returned token in the client instance.
func logiUser(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login User",
		Long:  `Login User.`,
		Run: func(cmd *cobra.Command, args []string) {
			username, err := cmd.Flags().GetString("username")
			if err != nil {
				cmd.PrintErr(err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.LoginRequest{
				Login:    username,
				Password: password,
			}

			result, err := client.Users.Login(cmd.Context(), &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("User not found")
					case codes.Unauthenticated:
						cmd.Print("Incorrect password or username")
					default:
						cmd.Println("Error:", st.Message())
					}
				} else {
					cmd.PrintErrf("Error: %v", err)
				}
			} else {
				client.Token = result.Token
				cmd.Print("Successfully registered")
			}
		},
	}
	cmd.Flags().StringP("username", "u", "", "Username")
	cmd.Flags().StringP("password", "p", "", "Password")
	err := cmd.MarkFlagRequired("username")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("password")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}
