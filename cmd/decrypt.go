package cmd

import (
	"flag"
	"fmt"
	"github.com/antonmarin/secret-yaml/documentManipulator"
	"github.com/antonmarin/secret-yaml/encryption"
	"github.com/antonmarin/secret-yaml/io"
	"github.com/antonmarin/secret-yaml/useCases"

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
		secret := flag.String("secret", "", "")
		flag.Parse()

		inputFile := io.NewFile(args[0])
		data, err := inputFile.AsString()
		if err != nil {
			return err
		}
		encryptionService, err := encryption.NewAesEncryptionService(*secret)
		if err != nil {
			return err
		}

		useCase := useCases.NewDecrypt(encryptionService, documentManipulator.NewYamlManipulator())
		result, err := useCase.Execute(data)
		if err != nil {
			return err
		}

		_, err = fmt.Println(string(result))
		if err != nil {
			return err
		}

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

type DecryptCommandUseCase interface {
	Execute(yaml string) (string, error)
}
