package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// decryptCmd represents the encrypt command
var decryptCmd = &cobra.Command{
	Short:   "Decrypt yaml file",
	Long:    `Decrypts file with --secret and outputs to stdout`,
	Use:     "decrypt path-to-yaml-file",
	Example: "  syml decrypt ~/chart/values.prod.yaml > ~/secrets/values.prod.decrypted.yaml",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("decrypt called with args: ")
		fmt.Println(args)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().String("secret", "", "Secret key to decode with")
	if err := decryptCmd.MarkFlagRequired("secret"); err != nil {
		panic(fmt.Errorf("Fatal error: %s \n", err))
	}
}
