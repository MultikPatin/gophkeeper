package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gothkeeper",
	Short: "It is a client-server to store private information.",
	Long: `It is a client-server system that allows users to securely store 
	login, passwords, binary data, and other private information.`,
}

func Execute(client *proto.GothKeeperClient) {
	rootCmd.AddCommand(SetupBinaryCommand(client))
	rootCmd.AddCommand(SetupCardCommand(client))
	rootCmd.AddCommand(SetupPasswordCommand(client))
	rootCmd.AddCommand(SetupUserCommand(client))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
