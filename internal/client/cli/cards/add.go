package cards

import "github.com/spf13/cobra"

func init() {
	cardCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add bank card data",
	Long:  `Add bank card data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
