/*
Copyright © 2022 Mary Karroqe <mkarroqe@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli-duck",
	Short: "ASCII duck friend to talk to through your screen",
	Long: `									
	Rubber duck debugging for the modern era!
	🤝 A virtual rubber duck that lives in your terminal to support you when you need it. 
	🦆 When called, cli-duck presents you with an ASCII duck. 
	🤓 This project is an exercise in creating command line tools in Go using Cobra.`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
