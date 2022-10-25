/*
Copyright © 2022 Mary Karroqe <mkarroqe@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sayhelloCmd represents the sayhello command
var sayhelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "The duck says hello.",
	Long:  "The duck will say hello. Default language is duckspeak, but it will speak english if you specify with --lang.",

	Run: func(cmd *cobra.Command, args []string) {
		fullDuck := ""

		// --lang
		translations := map[string]string{
			"albanian":  "mak",
			"duckspeak": "quack",
			"english":   "..hello",
			"french":    "coin",
		}
		lang, _ := cmd.Flags().GetString("lang")
		message := translations[lang]

		// --cool
		cool, _ := cmd.Flags().GetBool("cool")

		// --size
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
	sayCmd.AddCommand(sayhelloCmd)
	sayhelloCmd.PersistentFlags().String("lang", "duckspeak", "Set language to duckspeak or english.")
}
