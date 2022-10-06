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
var stdDuckL = "\n		      ██████                                    \n                    ██      ██                                  \n                  ██          ██                                \n                  ██      ██  ██                                \n                  ██        ░░░░██                              \n                    ██      ████                                \n      ██              ██  ██                                    \n    ██  ██        ████    ██                                    \n    ██    ████████          ██                                  \n    ██                        ██                                \n      ██              ██      ██                                \n      ██    ██      ██        ██                                \n        ██    ████████      ██                                  \n        ██                  ██                                  \n          ████          ████                                    \n              ██████████        \n"
var coolDuckL = "\n		      ██████                                    \n                    ██      ██                                  \n                  ██          ██                                \n                  ██     ⌐-██--██                                \n                  ██        ░░░░██                              \n                    ██      ████                                \n      ██              ██  ██                                    \n    ██  ██        ████    ██                                    \n    ██    ████████          ██                                  \n    ██                        ██                                \n      ██              ██      ██                                \n      ██    ██      ██        ██                                \n        ██    ████████      ██                                  \n        ██                  ██                                  \n          ████          ████                                    \n              ██████████        \n"

// hiCmd represents the hi command
var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "A duck summons",
	Long:  `Rubber duck will appear. It's ready to listen.`,

	Run: func(cmd *cobra.Command, args []string) {
		superSize, _ := cmd.Flags().GetString("size")
		coolStatus, _ := cmd.Flags().GetBool("cool")
		reqCount, _ := cmd.Flags().GetInt("count")

		for i := 0; i < reqCount; i++ {
			if superSize == "large" {
				if coolStatus == true {
					fmt.Println(coolDuckL)
				} else {
					fmt.Println(stdDuckL)
				}
			} else if superSize == "small" {
				if coolStatus == true {
					fmt.Println(coolDuck)
				} else {
					fmt.Println(stdDuck)
				}
			} else {
				fmt.Println("Sorry, we couldn't find any duckies of size " + superSize + ".")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(hiCmd)
	hiCmd.PersistentFlags().Bool("cool", false, "Rubber duck will be wearing sunglasses.")
	hiCmd.PersistentFlags().Int("count", 1, "Generates a specific number of ducks.")
	hiCmd.PersistentFlags().String("size", "small", "Big ducky?")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
