package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"

	"bufio"
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"

	"github.com/victor23d/etcd-batch/utils"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func main() {
	log := setLog()
	ReadFromFile(log)
	// ExampleKV_putErrorHandling()
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

func ReadFromCommand() {
	Prefix := os.Args[1]
	log.Println(Prefix)

	scanner := bufio.NewScanner(os.Stdin)
	var JsonString strings.Builder
	for scanner.Scan() {
		JsonString.WriteString(scanner.Text() + "\n")
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	log.Println(JsonString.String())
}

func ReadFromFile(log *logrus.Logger) {
	b, err := ioutil.ReadFile("foo.json")
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(string(b))
	var mf interface{}

	err = json.Unmarshal(b, &mf)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	m := mf.(map[string]interface{})

	log.Println(m)

	fp := make(map[string]interface{})
	utils.FlatMap(m, fp, "/", "", log)
	log.Println(fp)

}

func setLog() *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}
	return log
}

// TODO
// unknown file type
// dry run
