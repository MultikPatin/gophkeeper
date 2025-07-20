package passwords

import (
	"github.com/spf13/cobra"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Processing of login password pairs",
	Long: `Processing of login and password pairs. 
	Includes methods for saving, retrieving, modifying, and deleting.`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(passwordCmd)
}
