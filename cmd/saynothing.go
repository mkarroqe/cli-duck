/*
Copyright © 2022 Mary Karroqe <mkarroqe@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// saynothingCmd represents the saynothing command
var saynothingCmd = &cobra.Command{
	Use:   "nothing",
	Short: "The duck says nothing.",

	Run: func(cmd *cobra.Command, args []string) {
		fullDuck := ""
		message := "....."
		cool, _ := cmd.Flags().GetBool("cool")
		size, _ := cmd.Flags().GetString("size")

		duckie := Duck{
			Size:    size,
			Cool:    cool,
			Speaks:  true,
			Message: message,
		}

		if size == "little" {
			fullDuck = duckie.CreateLittleDuck()
		} else if size == "large" {
			fullDuck = duckie.CreateLargeDuck()
		} else {
			fullDuck = "A " + size + " duck is too much to handle."
		}

		fmt.Println(fullDuck)
	},
}

func init() {
	sayCmd.AddCommand(saynothingCmd)
}
