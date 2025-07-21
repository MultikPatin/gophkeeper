package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
)

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

func addCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add bank card data",
		Long:  `Add bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func getCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get bank card data",
		Long:  `Get bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func updateCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update bank card data",
		Long:  `Update bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func removeCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete bank card data",
		Long:  `Delete bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}
