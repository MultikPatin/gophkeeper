package cli

import (
	"github.com/spf13/cobra"
	"main/internal/client/app/proto"
	pb "main/proto"
)

func SetupCardCommand(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "card",
		Short: "Processing of bank card data",
		Long: `Processing of bank card data. 
		Includes methods for saving, retrieving, modifying, and deleting.`,
	}
	cmd.AddCommand(addCard(client))
	cmd.AddCommand(getCard(client))
	cmd.AddCommand(updateCard(client))
	cmd.AddCommand(removeCard(client))
	return cmd
}

func addCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add bank card data",
		Long:  `Add bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			bank, err := cmd.Flags().GetString("bank")
			if err != nil {
				cmd.PrintErr(err)
			}
			number, err := cmd.Flags().GetString("number")
			if err != nil {
				cmd.PrintErr(err)
			}
			dataEnd, err := cmd.Flags().GetString("dataEnd")
			if err != nil {
				cmd.PrintErr(err)
			}
			secretCode, err := cmd.Flags().GetString("secretCode")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardCreateRequest{
				Title:      title,
				Bank:       bank,
				Number:     number,
				DataEnd:    dataEnd,
				SecretCode: secretCode,
			}
			result, err := client.Cards.Add(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Save object with title: ", result.Title)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	cmd.Flags().StringP("bank", "b", "", "Bank name")
	cmd.Flags().StringP("number", "n", "", "Card number")
	cmd.Flags().StringP("dataEnd", "d", "", "Date end")
	cmd.Flags().StringP("secretCode", "s", "", "Secret code")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("bank")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("number")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("dataEnd")
	if err != nil {
		cmd.PrintErr(err)
	}
	err = cmd.MarkFlagRequired("secretCode")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

func getCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get bank card data",
		Long:  `Get bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardRequest{
				Title: title,
			}
			result, err := client.Cards.Get(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Get object with title: ", result.Title)
			cmd.Print("Bank name: ", result.Bank)
			cmd.Print("Card number: ", result.Number)
			cmd.Print("Date end: ", result.DataEnd)
			cmd.Print("Secret code: ", result.SecretCode)

		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

func updateCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update bank card data",
		Long:  `Update bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}
			bank, err := cmd.Flags().GetString("bank")
			if err != nil {
				cmd.PrintErr(err)
			}
			number, err := cmd.Flags().GetString("number")
			if err != nil {
				cmd.PrintErr(err)
			}
			dataEnd, err := cmd.Flags().GetString("dataEnd")
			if err != nil {
				cmd.PrintErr(err)
			}
			secretCode, err := cmd.Flags().GetString("secretCode")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardUpdateRequest{
				Title:      title,
				Bank:       bank,
				Number:     number,
				DataEnd:    dataEnd,
				SecretCode: secretCode,
			}
			result, err := client.Cards.Update(cmd.Context(), &cond)
			if err != nil {
				cmd.PrintErr(err)
			}

			cmd.Print("Update object with title: ", result.Title)
		},
	}
	cmd.Flags().StringP("title", "t", "", "Record title")
	cmd.Flags().StringP("bank", "b", "", "Bank name")
	cmd.Flags().StringP("number", "n", "", "Card number")
	cmd.Flags().StringP("dataEnd", "d", "", "Date end")
	cmd.Flags().StringP("secretCode", "s", "", "Secret code")
	err := cmd.MarkFlagRequired("title")
	if err != nil {
		cmd.PrintErr(err)
	}
	return cmd
}

func removeCard(client *proto.GothKeeperClient) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Delete bank card data",
		Long:  `Delete bank card data.`,
		Run: func(cmd *cobra.Command, args []string) {
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				cmd.PrintErr(err)
			}

			cond := pb.CardRequest{
				Title: title,
			}
			_, err = client.Cards.Delete(cmd.Context(), &cond)
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
