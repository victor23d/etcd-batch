/*
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// flatCmd represents the flat command
var flatCmd = &cobra.Command{
	Use:   "flat",
	Short: "flat a hierarchical file",
	Long:  `Example: etcd_batch flat -f foo.json`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("flat called")
	},
}

func init() {
	rootCmd.AddCommand(flatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
