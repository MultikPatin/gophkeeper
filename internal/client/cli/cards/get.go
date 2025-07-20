package cards

import (
	"github.com/spf13/cobra"
)

func init() {
	cardCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get bank card data",
	Long:  `Get bank card data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
