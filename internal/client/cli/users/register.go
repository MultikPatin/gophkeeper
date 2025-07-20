package users

import (
	"github.com/spf13/cobra"
)

func init() {
	userCmd.AddCommand(registerCmd)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register User",
	Long:  `Register User.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
