package utils

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/victor23d/etcd-batch/common"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/grpclog"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"localhost:2379"}
)

func TestExample(t *testing.T) {
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close() // make sure to close the client

	_, err = cli.Put(context.TODO(), "foo", "bbb")
	if err != nil {
		t.Fatal(err)
	}
}

func TestBatchFlatMap(t *testing.T) {
	t.Log("import")

	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer cli.Close() // make sure to close the client
	// Flat
	var m map[string]interface{}
	m, err = common.ReadJSONFromFile("../foo.json", log)
	if err != nil {
		t.Fatal(err)
	}
	fp := make(map[string]interface{})
	FlatMap(m, fp, "/", "")
	sfp := StringFlatedMap(fp)
	t.Log(sfp)
	// Batch
	err = BatchStringFlatedMap(context.TODO(), cli, sfp, "/PREFIX/")
	if err != nil {
		t.Fatal(err)
		t.Errorf("BatchStringFlatedMap failed")
	}
}
