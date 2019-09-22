package cmd

import (
	"fmt"
	"github.com/antonmarin/secret-yaml/random"
	"github.com/antonmarin/secret-yaml/useCases/generateSecretKey"

	"github.com/spf13/cobra"
)

// generateSecretKeyCmd represents the generateSecretKey command
var generateSecretKeyCmd = &cobra.Command{
	Use:     "generateSecretKey",
	Aliases: []string{"gensec"},
	Short:   "Generate secret key to encrypt/decrypt values",
	Long: `Generates AES secret key to use later with
encrypt/decrypt and outputs it to stdout.
Store it somewhere!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		useCase := generateSecretKey.NewGenerateSecretKey(new(random.CryptoGeneratorService))
		result, err := useCase.Execute()
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
