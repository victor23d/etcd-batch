/*
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/

package cmd

import (
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/clientv3"
	// "go.etcd.io/etcd/etcdctl/ctlv3/command"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "batch put keys",
	Long:  `Example: etcd-batch apply -f foo.json --prefix ""`,
	Run: func(cmd *cobra.Command, args []string) {
		putCommandFunc(cmd, args)
	},
}

func putCommandFunc(cmd *cobra.Command, args []string) {
	/* opts may contains
	leaseStr       string
	putPrevKV      bool
	putIgnoreVal   bool
	putIgnoreLease bool
	which this tool doesn't not contain for simplify
	*/

	opts := []clientv3.OpOption{}

	ctx, cancel := commandCtx(cmd)

	// resp, err := mustClientFromCmd(cmd).Put(ctx, "foo", "bar", opts...)
	_, err := mustClientFromCmd(cmd).Put(ctx, "foo", "bar", opts...)

	cancel()
	if err != nil {
		ExitWithError(ExitError, err)
	}
	// Too many dependencies
	// display.Put(*resp)

	log.Println("OK")
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
