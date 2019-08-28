/*
Copyright © 2019 victor23d <victor6742x@gmail.com>
This file is part of {{ .appName }}.
*/
package cmd

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/victor23d/etcd-batch/common"
	"github.com/victor23d/etcd-batch/utils"
)

var (
	prefix   string
	filename string
	log      = logrus.New()
)

const (
	sep = "/"
)

// flatCmd represents the flat command
var flatCmd = &cobra.Command{
	Use:   "flat",
	Short: "flat a hierarchical file",
	Long:  `Example: etcd-batch flat -f foo.json`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println(filename)
		if filename == "" {
			log.Fatal(errors.New("must specify -f"))
		}
		m, err := common.ReadJSONFromFile(filename, log)
		if err != nil {
			log.Fatal(err)
		}
		fp := make(map[string]interface{})
		utils.FlatMap(m, fp, sep, prefix)

		log.Println("FlatedMap")
		log.Println(fp)

		sfp := utils.StringFlatedMap(fp)
		log.Println("stringify")
		log.Println(sfp)
		sb := utils.TextSFP(sfp)
		log.Println(sb.String())

	},
}

func init() {
	rootCmd.AddCommand(flatCmd)
	flatCmd.Flags().StringVarP(&filename, "filename", "f", "", "the file to apply")
}
