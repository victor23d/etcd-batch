/*
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/

package cmd

import (
	"errors"

	"github.com/prometheus/common/log"
	"github.com/spf13/cobra"
	"github.com/victor23d/etcd-batch/common"
	"github.com/victor23d/etcd-batch/flat"
	"go.etcd.io/etcd/clientv3"
	"os"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "batch put keys",
	Long:  `Example: etcd-batch apply -f foo.json --prefix "" -d "/"`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("filename=%s, prefix=%s, delimiter=%s \n", filename, prefix, sep)
		if filename == "" {
			log.Fatal(errors.New("must specify -f"))
		}
		m, err := common.ReadJSONFromFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		fp := make(map[string]interface{})
		flat.FlatMap(m, fp, sep, prefix)

		sfp := flat.StringFlatedMap(fp)
		log.Info(sfp)

		// Suppress message: pkg/flags: unrecognized environment variable ETCDCTL_API
		os.Unsetenv("ETCDCTL_API")

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
		cli := mustClientFromCmd(cmd)

		i := 0
		for key, value := range sfp {
			i++
			_, err := cli.Put(ctx, key, value, opts...)
			if err != nil {
				ExitWithError(ExitError, err)
			}
		}
		cancel()
		log.Infof("OK, number of keys put: %d", i)
		defer cli.Close()
		// Too many dependencies
		// display.Put(*resp)

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	applyCmd.Flags().StringVarP(&filename, "filename", "f", "", "the file to apply")
	applyCmd.Flags().StringVarP(&sep, "delimiter", "d", "/", "keys are delimited by")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
