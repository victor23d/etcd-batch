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
)

// flatCmd represents the flat command
var flatCmd = &cobra.Command{
	Use:   "flat",
	Short: "flat a hierarchical file",
	Long:  `Example: etcd-batch flat -f foo.json`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info(filename)
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
		sb := flat.TextSFP(sfp)

		log.Info("=== Text ===")
		log.Info(sb.String())

	},
}

func init() {
	rootCmd.AddCommand(flatCmd)
	flatCmd.Flags().StringVarP(&filename, "filename", "f", "", "the file to apply")
}
