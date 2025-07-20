package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/client/cli/binaries"
	"main/internal/client/cli/cards"
	"main/internal/client/cli/passwords"
	"main/internal/client/cli/users"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gothkeeper",
	Short: "It is a client-server to store private information.",
	Long: `It is a client-server system that allows users to securely store 
	login, passwords, binary data, and other private information.`,
}

func Execute() {
	users.Init(rootCmd)
	passwords.Init(rootCmd)
	cards.Init(rootCmd)
	binaries.Init(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
