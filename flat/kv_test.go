package flat

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
	dialTimeout    = 3 * time.Second
	RequestTimeout = 6 * time.Second
	// Doesn't work with docker etcd
	// endpoints      = []string{"localhost:2379"}
	endpoints = []string{"127.0.0.1:2379"}
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

	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	_, err = cli.Put(ctx, "foo", "bbb")
	if err != nil {
		t.Fatal(err)
	}
	cancel()
}

func TestBatchFlatMap(t *testing.T) {
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))
	t.Log("import")

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
	m, err = common.ReadJSONFromFile("../foo.json")
	if err != nil {
		t.Fatal(err)
	}
	fp := make(map[string]interface{})
	FlatMap(m, fp, "/", "")
	sfp := StringFlatedMap(fp)
	t.Log(sfp)
	// Batch
	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	err = BatchStringFlatedMap(ctx, cli, sfp, "/PREFIX/")
	if err != nil {
		t.Fatal(err)
		t.Errorf("BatchStringFlatedMap failed")
	}
	cancel()
}
