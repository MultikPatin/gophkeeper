package cards

import (
	"github.com/spf13/cobra"
)

func init() {
	cardCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update bank card data",
	Long:  `Update bank card data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
