package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Short:   "Encrypt yaml file",
	Long:    `Encrypts file with --secret and outputs to stdout`,
	Use:     "encrypt path-to-yaml-file",
	Example: "  syml encrypt ~/secrets/values.prod.decrypted.yaml > ~/chart/values.prod.yaml",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("encrypt called with args: ")
		fmt.Println(args)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	encryptCmd.Flags().String("secret", "", "Secret key to encode with")
	if err := encryptCmd.MarkFlagRequired("secret"); err != nil {
		panic(fmt.Errorf("Fatal error: %s \n", err))
	}
}
