package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
)

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

func addPassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add login password pair",
		Long:  `Add login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func getPassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get login password pair",
		Long:  `Get login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func updatePassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update login password pair",
		Long:  `Update login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func removePassword(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete login password pair",
		Long:  `Delete login password pair.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}
