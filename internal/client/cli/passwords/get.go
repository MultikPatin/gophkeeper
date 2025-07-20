package passwords

import (
	"github.com/spf13/cobra"
)

func init() {
	passwordCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get login password pair",
	Long:  `Get login password pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
