/*
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unflatCmd represents the unflat command
var unflatCmd = &cobra.Command{
	Use:   "unflat",
	Short: "unflat a plain text file to a hierarchical file",
	Long:  `Example: etcd-batch unflat -f foo.json -o json`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unflat called")
	},
}

func init() {
	rootCmd.AddCommand(unflatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unflatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unflatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
