package main

import (
	"context"
	"time"

	"github.com/victor23d/etcd-batch/common"
	"github.com/victor23d/etcd-batch/utils"
	"go.etcd.io/etcd/clientv3"
)

func main() {
	log := common.SetLog()
	// ExampleKV_putErrorHandling()

	var m map[string]interface{}
	m, err := common.ReadJSONFromFile("foo.json", log)
	if err != nil {
		log.Fatal(err)
	}

	// use var instead of make will cause panic: assignment to entry in nil map
	// var fp map[string]interface{}
	fp := make(map[string]interface{})
	utils.FlatMap(m, fp, "/", "", log)
	log.Println(fp)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = utils.KV_putErrorHandling(ctx, cli, "foo", "bar", log)
	if err != nil {
		log.Fatal(err)
	}

}

// TODO
// unknown file type
// dry run
