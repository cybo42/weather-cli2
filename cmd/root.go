package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather-cli",
	Short: "A simple CLI application",
	Long:  `A simple CLI application that prints "hello world".`,
	// Run is the function that will be called when the command is executed
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}