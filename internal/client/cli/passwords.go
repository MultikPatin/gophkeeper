package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
	pb "main/proto"
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
			result, err := client.Passwords.Add(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Save object with title: ", result.Title)
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
			result, err := client.Passwords.Get(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Get object with title: ", result.Title)
			cmd.Print("Login: ", result.Login)
			cmd.Print("Password: ", result.Password)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

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
			result, err := client.Passwords.Update(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Update object with title: ", result.Title)
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
			_, err = client.Passwords.Delete(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Delete object with title: ", title)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}

	return cmd
}
