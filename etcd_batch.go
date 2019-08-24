package main

import (
	"context"
	"log"
	"time"

	"github.com/victor23d/etcd-batch/common"
	"github.com/victor23d/etcd-batch/utils"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
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
}

func ExampleKV_putErrorHandling() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = cli.Put(ctx, "foo", "bar")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			log.Printf("ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			log.Printf("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			log.Printf("client-side error: %v\n", err)
		default:
			log.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
	}
	log.Println("OK")
}

// TODO
// unknown file type
// dry run
