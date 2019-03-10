package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "syml",
	Short: "Tool to secret yaml values",
	Long: `Tool to secret yaml values.
Made to be quickly installed right in your ci pipeline.

Base usage scenario:
  1. Generate secret key with 'syml generateSecretKey'
  2. Use secret to encrypt file with 'syml encrypt'
  3. Install tool in your pipeline
  4. Use same secret to decrypt file with 'syml decrypt'`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

type Input interface {
	AsString() (string, error)
}
