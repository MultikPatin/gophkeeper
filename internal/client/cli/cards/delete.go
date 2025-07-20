package cards

import (
	"github.com/spf13/cobra"
)

func init() {
	cardCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete bank card data",
	Long:  `Delete bank card data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
