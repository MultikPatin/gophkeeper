package binaries

import (
	"github.com/spf13/cobra"
)

func init() {
	binaryCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add binary data",
	Long:  `Add binary data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
