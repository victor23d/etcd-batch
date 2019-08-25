package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"

	"github.com/victor23d/etcd-batch/common"
	"github.com/victor23d/etcd-batch/utils"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/grpclog"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
	endpoints      = []string{"localhost:2379"}
	log            = logrus.New()
)

func TestExample(t *testing.T) {
	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close() // make sure to close the client

	_, err = cli.Put(context.TODO(), "foo", "bbb")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
}

func TestStringFlatedMap(t *testing.T) {
	// sfp := StringFlatedMap(fp)

}

func TestBatchFlatMap(t *testing.T) {
	log.Println("import")

	clientv3.SetLogger(grpclog.NewLoggerV2(os.Stderr, os.Stderr, os.Stderr))

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close() // make sure to close the client
	// Flat
	var m map[string]interface{}
	m, err = common.ReadJSONFromFile("../foo.json", log)
	if err != nil {
		log.Fatal(err)
	}
	fp := make(map[string]interface{})
	utils.FlatMap(m, fp, "/", "", log)
	sfp := utils.StringFlatedMap(fp)
	log.Println(sfp)
	// Batch
	err = utils.BatchStringFlatedMap(context.TODO(), cli, sfp, "/PREFIX/", log)
	if err != nil {
		log.Fatal(err)
		t.Errorf("BatchStringFlatedMap failed")
	}
	log.Println("OK")
}
