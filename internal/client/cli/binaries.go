package cli

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/client/app/proto"
	pb "main/proto"
)

// SetupBinaryCommand initializes the main binary processing commands.
// This function creates a parent command that groups all operations dealing with binary data management.
// It adds four child commands: adding, getting, updating, and removing binaries.
//
// The function does not return an error since it's only setting up the structure.
func SetupBinaryCommand(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Processing of binary data",
		Long: `Processing of binary data. 
		Includes methods for saving, retrieving, modifying, and deleting.`,
	}
	cmd.AddCommand(addBinary(client))
	cmd.AddCommand(getBinary(client))
	cmd.AddCommand(updateBinary(client))
	cmd.AddCommand(removeBinary(client))
	return cmd
}

// addBinary implements the logic for adding new binary records via the API.
// It takes a record's title and its associated binary content, sends them over gRPC to create a new entry.
// Potential errors include duplicate titles (`AlreadyExists`), unauthorized access (`Unauthenticated`), etc.
// On success, it outputs the saved object's title.
func addBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add binary data",
		Long:  `Add binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			binary, err := cmd.Flags().GetBytesHex("binary")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesCreateRequest{
				Title: title,
				Data:  binary,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Binaries.Add(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.AlreadyExists:
						cmd.Print("Binary already exists")
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
	cmd.Flags().StringP("binary", "b", "", "Binary data")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("binary")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// getBinary fetches binary data by providing its unique title.
// Errors can occur due to lack of authorization (`Unauthenticated`) or if no matching record is found (`NotFound`).
// Upon success, it displays the retrieved binary record's title and contents.
func getBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get binary data",
		Long:  `Get binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesRequest{
				Title: title,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Binaries.Get(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Binary not found")
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
				cmd.Print("Binary data: ", result.Data)
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

// updateBinary updates an existing binary record using its title and updated binary content.
// Possible errors arise from either unauthenticated requests (`Unauthenticated`) or trying to modify a nonexistent record (`NotFound`).
// A successful operation results in printing the updated record's title.
func updateBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update binary data",
		Long:  `Update binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			binary, err := cmd.Flags().GetBytesHex("binary")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesUpdateRequest{
				Title: title,
				Data:  binary,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Binaries.Update(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Binary not found")
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
	cmd.Flags().StringP("binary", "b", "", "Binary data")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("binary")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// removeBinary deletes a binary record identified by its title.
// Deletion may fail because of insufficient authentication (`Unauthenticated`) or attempting to delete a non-existent record (`NotFound`).
// In case of success, it confirms deletion through a print statement.
func removeBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete binary data",
		Long:  `Delete binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesRequest{
				Title: title,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			_, err = client.Binaries.Delete(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Binary not found")
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
