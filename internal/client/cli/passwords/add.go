package passwords

import "github.com/spf13/cobra"

func init() {
	passwordCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add login password pair",
	Long:  `Add login password pair.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
	},
}
