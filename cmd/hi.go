/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stdDuck = "    __\n___( o)>\n\\ <_. )\n `---'"
var coolDuck = "    __\n___(⌐■)>\n\\ <_. )\n `---'"

// hiCmd represents the hi command
var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "A duck summons",
	Long:  `Rubber duck will appear. It's ready to listen.`,

	Run: func(cmd *cobra.Command, args []string) {
		coolStatus, _ := cmd.Flags().GetBool("cool")
		reqCount, _ := cmd.Flags().GetInt("count")

		for i := 0; i < reqCount; i++ {
			if coolStatus == true {
				fmt.Println(coolDuck)
			} else {
				fmt.Println(stdDuck)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(hiCmd)
	hiCmd.PersistentFlags().Bool("cool", false, "Rubber duck will be wearing sunglasses.")
	hiCmd.PersistentFlags().Int("count", 1, "Generates a specific number of ducks.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
