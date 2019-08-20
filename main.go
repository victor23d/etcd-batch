package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func main() {
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

// TODO
// unknown file type
// dry run
