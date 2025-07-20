package passwords

import (
	"github.com/spf13/cobra"
)

func init() {
	passwordCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete login password pair",
	Long:  `Delete login password pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
