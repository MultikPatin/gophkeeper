package binaries

import (
	"github.com/spf13/cobra"
)

func init() {
	binaryCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get binary data",
	Long:  `Get binary data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
