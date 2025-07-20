package binaries

import (
	"github.com/spf13/cobra"
)

func init() {
	binaryCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update binary data",
	Long:  `Update binary data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
