package users

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User authorization and registration",
	Long: `The section contains methods for user 
	authorization and registration based on JWT tokens`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(userCmd)
}
