package utils

import (
	"github.com/victor23d/etcd-batch/common"
	"testing"
)

func TestFlatMap(t *testing.T) {
	t.Log("flat")
	var m map[string]interface{}
	m, err := common.ReadJSONFromFile("../foo.json", log)
	if err != nil {
		t.Fatal(err)
	}

	// use var instead of make will cause panic: assignment to entry in nil map
	// var fp map[string]interface{}
	fp := make(map[string]interface{})
	FlatMap(m, fp, "/", "")
	t.Log(fp)

}

func TestStringFlatedMap(t *testing.T) {
	// sfp := StringFlatedMap(fp)

}
