package cli

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/client/app/proto"
	pb "main/proto"
)

// SetupPasswordCommand initializes the main command group for password management.
// It includes various subcommands to manage login-password pairs stored securely.
// No direct error handling happens here; instead, errors are handled within individual subcommands.
func SetupPasswordCommand(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Processing of login password pairs",
		Long: `Processing of login and password pairs. 
		Includes methods for saving, retrieving, modifying, and deleting.`,
	}
	cmd.AddCommand(addPassword(client))
	cmd.AddCommand(getPassword(client))
	cmd.AddCommand(updatePassword(client))
	cmd.AddCommand(removePassword(client))
	return cmd
}

// addPassword creates a new login-password pair.
// It accepts three required parameters—title, login, and password—and uses them to send a gRPC request.
// Errors include conflicts (`AlreadyExists`) and authentication issues (`Unauthenticated`).
func addPassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add login password pair",
		Long:  `Add login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			login, err := cmd.Flags().GetString("login")
			if err != nil {
				cmd.PrintErr(err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.PasswordCreateRequest{
				Title:    title,
				Login:    login,
				Password: password,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Passwords.Add(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.AlreadyExists:
						cmd.Print("Password already exists")
					case codes.Unauthenticated:
						cmd.Print("invalid token")
					default:
						cmd.Println("Error:", st.Message())
					}
				} else {
					cmd.PrintErrf("Error: %v", err)
				}
			} else {
				cmd.Print("Save object with title: ", result.Title)
			}
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	cmd.Flags().StringP("login", "l", "", "Login")
	cmd.Flags().StringP("password", "p", "", "Password")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("login")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("password")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// getPassword retrieves a previously added login-password pair by its title.
// It connects to the gRPC server to obtain the necessary data.
// Possible errors include an absent record (`NotFound`) or an invalid token (`Unauthenticated`).
func getPassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get login password pair",
		Long:  `Get login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.PasswordRequest{
				Title: title,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Passwords.Get(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Password not found")
					case codes.Unauthenticated:
						cmd.Print("invalid token")
					default:
						cmd.Println("Error:", st.Message())
					}
				} else {
					cmd.PrintErrf("Error: %v", err)
				}
			} else {
				cmd.Print("Get object with title: ", result.Title)
				cmd.Print("Login: ", result.Login)
				cmd.Print("Password: ", result.Password)
			}
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// updatePassword alters an existing login-password pair.
// It accepts the same parameters as addPassword but focuses on modifying rather than creating a new record.
// Errors might arise due to a missing record (`NotFound`) or an improper token (`Unauthenticated`).
func updatePassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update login password pair",
		Long:  `Update login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			login, err := cmd.Flags().GetString("login")
			if err != nil {
				cmd.PrintErr(err)
			}
			password, err := cmd.Flags().GetString("password")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.PasswordUpdateRequest{
				Title:    title,
				Login:    login,
				Password: password,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Passwords.Update(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Password not found")
					case codes.Unauthenticated:
						cmd.Print("invalid token")
					default:
						cmd.Println("Error:", st.Message())
					}
				} else {
					cmd.PrintErrf("Error: %v", err)
				}
			} else {
				cmd.Print("Update object with title: ", result.Title)
			}
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	cmd.Flags().StringP("login", "l", "", "Login")
	cmd.Flags().StringP("password", "p", "", "Password")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// removePassword eliminates a login-password pair based on its title.
// It initiates a gRPC request to permanently delete the selected record.
// Errors could stem from the absence of the record (`NotFound`) or invalid token usage (`Unauthenticated`).
func removePassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete login password pair",
		Long:  `Delete login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.PasswordRequest{
				Title: title,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			_, err = client.Passwords.Delete(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Password not found")
					case codes.Unauthenticated:
						cmd.Print("invalid token")
					default:
						cmd.Println("Error:", st.Message())
					}
				} else {
					cmd.PrintErrf("Error: %v", err)
				}
			} else {
				cmd.Print("Successfully deleted")
			}
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}

	return cmd
}
