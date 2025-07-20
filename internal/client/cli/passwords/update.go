package passwords

import (
	"github.com/spf13/cobra"
)

func init() {
	passwordCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update login password pair",
	Long:  `Update login password pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
