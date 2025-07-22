package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
	pb "main/proto"
)

func SetupBinaryCommand(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binary",
		Short: "Processing of binary data",
		Long: `Processing of binary data. 
		Includes methods for saving, retrieving, modifying, and deleting.`,
	}
	cmd.AddCommand(addBinary(client))
	cmd.AddCommand(getBinary(client))
	cmd.AddCommand(updateBinary(client))
	cmd.AddCommand(removeBinary(client))
	return cmd
}

func addBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add binary data",
		Long:  `Add binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			binary, err := cmd.Flags().GetBytesHex("binary")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesCreateRequest{
				Title: title,
				Data:  binary,
			}
			result, err := client.Binaries.Add(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Save object with title: ", result.Title)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	cmd.Flags().StringP("binary", "b", "", "Binary data")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("binary")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

func getBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get binary data",
		Long:  `Get binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesRequest{
				Title: title,
			}
			result, err := client.Binaries.Get(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Get object with title: ", result.Title)
			cmd.Print("Binary data: ", result.Data)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

func updateBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update binary data",
		Long:  `Update binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			binary, err := cmd.Flags().GetBytesHex("binary")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesUpdateRequest{
				Title: title,
				Data:  binary,
			}
			result, err := client.Binaries.Update(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Update object with title: ", result.Title)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	cmd.Flags().StringP("binary", "b", "", "Binary data")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("binary")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

func removeBinary(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete binary data",
		Long:  `Delete binary data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.BinariesRequest{
				Title: title,
			}
			_, err = client.Binaries.Delete(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Delete object with title: ", title)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}
