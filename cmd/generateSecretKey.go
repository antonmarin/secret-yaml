package cmd

import (
	"fmt"
	"github.com/antonmarin/secret-yaml/useCases"

	"github.com/spf13/cobra"
)

// generateSecretKeyCmd represents the generateSecretKey command
var generateSecretKeyCmd = &cobra.Command{
	Use:   "generateSecretKey",
	Short: "Generate secret key to encrypt/decrypt values",
	Long: `Generates AES secret key to use later with 
encrypt/decrypt and outputs it to stdout.
Store it somewhere!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := useCases.GenerateSecretKeyUseCase.Execute()
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", result)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateSecretKeyCmd)
}

type GenerateSecretKeyCommandUseCase interface {
	Execute() ([]byte, error)
}
