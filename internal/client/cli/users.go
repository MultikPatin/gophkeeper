package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
	pb "main/proto"
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
				cmd.PrintErr(err)
			}
			client.Token = result.Token

			cmd.Print("Successfully registered")
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
				cmd.PrintErr(err)
			}
			client.Token = result.Token

			cmd.Print("Successfully logged in")
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
