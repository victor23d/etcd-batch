package main

import (
	"context"
	json "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"os"
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

func main() {
	log := logrus.New()
	log = common.SetLog(log)
	// ExampleKV_putErrorHandling()
	handleArgs(log)

}

func handleArgs(log *logrus.Logger) {
	if len(os.Args) < 3 {
		// log.Fatal(len(os.Args))
		log.Fatal("etcd_batch [ import | dump | flat ] \"prefix\" json")
	}

	if os.Args[1] == "import" || os.Args[1] == "import-file" {
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

		if os.Args[1] == "import-file" {
			m, err = common.ReadJSONFromFile(os.Args[3], log)
			if err != nil {
				log.Fatal(err)
			}
		}
		if os.Args[1] == "import" {
			s, err := common.ReadStringFromCommand()
			json.Unmarshal([]byte(s), &m)
			if err != nil {
				log.Fatal(err)
			}
		}
		fp := make(map[string]interface{})
		utils.FlatMap(m, fp, "/", "", log)
		sfp := utils.StringFlatedMap(fp)
		log.Println(sfp)
		// Batch
		prefix := os.Args[2]
		err = utils.BatchStringFlatedMap(context.TODO(), cli, sfp, prefix, log)
		if err != nil {
			log.Fatal(err)
			// t.Errorf("BatchStringFlatedMap failed")
		}
		log.Println("OK")
	}

	if os.Args[1] == "flat" {
		log.Println("flat")
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

	}
}

// TODO
// unknown file type
// dry run
