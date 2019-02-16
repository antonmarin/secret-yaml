package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateSecretKeyCmd represents the generateSecretKey command
var generateSecretKeyCmd = &cobra.Command{
	Use:   "generateSecretKey",
	Short: "Generate secret key to encrypt/decrypt values",
	Long: `Generates AES secret key to use later with 
encrypt/decrypt and outputs it to stdout.
Store it somewhere!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateSecret called")
	},
}

func init() {
	rootCmd.AddCommand(generateSecretKeyCmd)
}
