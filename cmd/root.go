package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Version is set during build using ldflags
var Version = "dev"

// Flag to determine if output should be reversed
var reverseOutput bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "weather-cli",
	Version: Version,
	Short:   "A simple CLI application",
	Long:    `A simple CLI application that prints "hello world".`,
	// Run is the function that will be called when the command is executed
	Run: func(cmd *cobra.Command, args []string) {
		if reverseOutput {
			// Reverse "hello world" and print it
			message := "hello world"
			runes := []rune(message)
			for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
				runes[i], runes[j] = runes[j], runes[i]
			}
			fmt.Println(string(runes))
		} else {
			fmt.Println("hello world")
		}
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

func init() {
	// Add the reverse flag to the root command
	rootCmd.Flags().BoolVarP(&reverseOutput, "reverse", "r", false, "Print hello world backwards")
}