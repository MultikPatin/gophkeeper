package binaries

import (
	"github.com/spf13/cobra"
)

var binaryCmd = &cobra.Command{
	Use:   "binary",
	Short: "Processing of binary data",
	Long: `Processing of binary data. 
	Includes methods for saving, retrieving, modifying, and deleting.`,
}

func Init(rootCmd *cobra.Command) {
	rootCmd.AddCommand(binaryCmd)
}
