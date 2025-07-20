package cards

import (
	"github.com/spf13/cobra"
)

var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Processing of bank card data",
	Long: `Processing of bank card data. 
	Includes methods for saving, retrieving, modifying, and deleting.`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cardCmd)
}
