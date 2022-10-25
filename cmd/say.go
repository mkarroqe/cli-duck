/*
Copyright © 2022 Mary Karroqe <mkarroqe@gmail.com>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "The duck will speak in its native tongue.",
	Long:  "The duck will speak in its native tongue if you ask it to say hello",

	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Sub command required- what do you want the duck to say?")
	},
}

func init() {
	rootCmd.AddCommand(sayCmd)
	sayCmd.PersistentFlags().Bool("cool", false, "Give the duck sunglasses.")
	sayCmd.PersistentFlags().String("size", "little", "Set duck size to little or large.")
}
