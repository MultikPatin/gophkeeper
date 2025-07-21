package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
)

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

func addBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add binary data",
		Long:  `Add binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func getBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get binary data",
		Long:  `Get binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func updateBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update binary data",
		Long:  `Update binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func removeBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete binary data",
		Long:  `Delete binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}
