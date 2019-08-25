package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/victor23d/etcd-batch/common"
	"github.com/victor23d/etcd-batch/utils"
	"testing"
)

func TestFlatMap(t *testing.T) {
	log := logrus.New()
	log = common.SetLog(log)
	log.Println("flat")
	var m map[string]interface{}
	m, err := common.ReadJSONFromFile("../foo.json", log)
	if err != nil {
		log.Fatal(err)
	}

	// use var instead of make will cause panic: assignment to entry in nil map
	// var fp map[string]interface{}
	fp := make(map[string]interface{})
	utils.FlatMap(m, fp, "/", "", log)
	log.Println(fp)

}
