package users

import (
	"github.com/spf13/cobra"
)

func init() {
	userCmd.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login User",
	Long:  `Login User.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
