package binaries

import (
	"github.com/spf13/cobra"
)

func init() {
	binaryCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete binary data",
	Long:  `Delete binary data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
