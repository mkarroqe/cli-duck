/*
Copyright © 2022 Mary Karroqe <mkarroqe@gmail.com>
*/
package cmd

import (
	// "errors"

	"fmt"

	"github.com/spf13/cobra"
)

// hiCmd represents the hi command
var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "A duck summons",
	Long:  `Rubber duck will appear. It's ready to listen.`,

	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetString("size")
		cool, _ := cmd.Flags().GetBool("cool")
		count, _ := cmd.Flags().GetInt("count")
		fullDuck := ""

		// we only need to create the duck once, since right now, multiple counts will always be of the same duck
		duckie := Duck{
			Size:   size,
			Cool:   cool,
			Speaks: false,
		}

		for i := 0; i < count; i++ {
			if size == "little" {
				fullDuck = duckie.CreateLittleDuck()
			} else if size == "large" {
				fullDuck = duckie.CreateLargeDuck()
			} else {
				fullDuck = "A " + size + " duck is too much to handle."
			}

			fmt.Println(fullDuck)
		}
	},
}

func init() {
	rootCmd.AddCommand(hiCmd)
	hiCmd.PersistentFlags().Bool("cool", false, "Rubber duck will be wearing sunglasses.")
	hiCmd.PersistentFlags().Int("count", 1, "Generates a specific number of ducks.")
	hiCmd.PersistentFlags().String("size", "small", "Generates a ducky of size small (default) or large.")
}
