/*
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "put keys",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apply called")
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
