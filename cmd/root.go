package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-example",
	Short: "go-cli-example is a cli tool for performing basic mathematical operations",
	Long:  "go-cli-example is a cli tool for performing basic mathematical operations - addition, multiplication, division and subtraction.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing go-cli-example '%s'\n", err)
		os.Exit(1)
	}
}
