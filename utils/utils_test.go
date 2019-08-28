package utils

import (
	"github.com/victor23d/etcd-batch/common"
	"testing"
)

func TestFlatMap(t *testing.T) {
	log.Println("flat")
	var m map[string]interface{}
	m, err := common.ReadJSONFromFile("../foo.json", log)
	if err != nil {
		log.Fatal(err)
	}

	// use var instead of make will cause panic: assignment to entry in nil map
	// var fp map[string]interface{}
	fp := make(map[string]interface{})
	FlatMap(m, fp, "/", "", log)
	log.Println(fp)

}
