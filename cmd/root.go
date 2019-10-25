package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grok",
	Short: "parse log files with grok",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(version string) {
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}
`)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
