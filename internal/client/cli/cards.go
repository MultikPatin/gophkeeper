package cli

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"main/internal/client/app/proto"
	pb "main/proto"
)

// SetupCardCommand configures the top-level command for managing bank cards.
// It provides functionality for creating, reading, updating, and deleting bank card entries.
// No error handling needed at this level as it merely organizes sub-commands.
func SetupCardCommand(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Processing of bank card data",
		Long: `Processing of bank card data. 
		Includes methods for saving, retrieving, modifying, and deleting.`,
	}
	cmd.AddCommand(addCard(client))
	cmd.AddCommand(getCard(client))
	cmd.AddCommand(updateCard(client))
	cmd.AddCommand(removeCard(client))
	return cmd
}

// addCard manages the addition of a new bank card record.
// It requires multiple parameters such as title, bank name, card number, expiration date, and security code.
// The process involves sending these details to the backend via gRPC.
// Potential errors include duplicated card records (`AlreadyExists`) and invalid/unauthorized tokens (`Unauthenticated`).
func addCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add bank card data",
		Long:  `Add bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			bank, err := cmd.Flags().GetString("bank")
			if err != nil {
				cmd.PrintErr(err)
			}
			number, err := cmd.Flags().GetString("number")
			if err != nil {
				cmd.PrintErr(err)
			}
			dataEnd, err := cmd.Flags().GetString("dataEnd")
			if err != nil {
				cmd.PrintErr(err)
			}
			secretCode, err := cmd.Flags().GetString("secretCode")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardCreateRequest{
				Title:      title,
				Bank:       bank,
				Number:     number,
				DataEnd:    dataEnd,
				SecretCode: secretCode,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Cards.Add(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.AlreadyExists:
						cmd.Print("Card already exists")
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
	cmd.Flags().StringP("bank", "b", "", "Bank name")
	cmd.Flags().StringP("number", "n", "", "Card number")
	cmd.Flags().StringP("dataEnd", "d", "", "Date end")
	cmd.Flags().StringP("secretCode", "s", "", "Secret code")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("bank")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("number")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("dataEnd")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("secretCode")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// getCard retrieves details about a bank card given its title.
// It communicates with the gRPC server to fetch the requested card’s attributes.
// Potential issues include an invalid token (`Unauthenticated`) or a missing card record (`NotFound`).
func getCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get bank card data",
		Long:  `Get bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardRequest{
				Title: title,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Cards.Get(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Card not found")
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
				cmd.Print("Bank: ", result.Bank)
				cmd.Print("Card number: ", result.Number)
				cmd.Print("Date end: ", result.DataEnd)
				cmd.Print("Secret code: ", result.SecretCode)
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

// updateCard modifies an existing bank card record by its title.
// It expects several inputs (like bank name, card number, expiration date, and security code), which are then sent to the gRPC server.
// Common errors include a non-existent card (`NotFound`) or failed authentication (`Unauthenticated`).
func updateCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update bank card data",
		Long:  `Update bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			bank, err := cmd.Flags().GetString("bank")
			if err != nil {
				cmd.PrintErr(err)
			}
			number, err := cmd.Flags().GetString("number")
			if err != nil {
				cmd.PrintErr(err)
			}
			dataEnd, err := cmd.Flags().GetString("dataEnd")
			if err != nil {
				cmd.PrintErr(err)
			}
			secretCode, err := cmd.Flags().GetString("secretCode")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardUpdateRequest{
				Title:      title,
				Bank:       bank,
				Number:     number,
				DataEnd:    dataEnd,
				SecretCode: secretCode,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			result, err := client.Cards.Update(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Card not found")
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
	cmd.Flags().StringP("bank", "b", "", "Bank name")
	cmd.Flags().StringP("number", "n", "", "Card number")
	cmd.Flags().StringP("dataEnd", "d", "", "Date end")
	cmd.Flags().StringP("secretCode", "s", "", "Secret code")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

// removeCard removes a bank card record specified by its title.
// It makes use of gRPC to perform the deletion action.
// Possible problems include incorrect authentication (`Unauthenticated`) or absence of the target card (`NotFound`).
func removeCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete bank card data",
		Long:  `Delete bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardRequest{
				Title: title,
			}

			ctx := cmd.Context()
			newCtx := metadata.NewOutgoingContext(ctx, metadata.Pairs("token", client.Token))

			_, err = client.Cards.Delete(newCtx, &cond)
			if err != nil {
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.NotFound:
						cmd.Print("Card not found")
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
