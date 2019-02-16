package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateSecretCmd represents the generateSecret command
var generateSecretCmd = &cobra.Command{
	Use:   "generateSecret",
	Short: "Generate secret key to encrypt/decrypt values",
	Long: `Generates AES secret key to use later with 
encrypt/decrypt and outputs it to stdout.
Store it somewhere!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateSecret called")
	},
}

func init() {
	rootCmd.AddCommand(generateSecretCmd)
}
