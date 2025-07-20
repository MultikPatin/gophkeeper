package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gothkeeper",
	Short: "A brief description of your CLI application",
	Long: `A longer description that explains your CLI application in detail, 
    including available commands and their usage.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("Welcome to gothkeeper! Use --help for usage.")
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
