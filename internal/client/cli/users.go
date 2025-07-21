package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
)

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

func registerUser(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register User",
		Long:  `Register User.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}

func logiUser(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login User",
		Long:  `Login User.`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: implement
		},
	}
	return cmd
}
