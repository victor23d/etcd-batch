package utils

import (
	"context"
	"github.com/sirupsen/logrus"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

func KV_putErrorHandling(ctx context.Context, cli *clientv3.Client, key, val string, log *logrus.Logger) (*clientv3.PutResponse, error) {
	PutResponse, err := cli.Put(ctx, "", "sample_value")
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
	return PutResponse, err
	// Output: client-side error: etcdserver: key is not provided
}

func KV_getAnddelete(ctx context.Context, cli *clientv3.Client, key, val string, log *logrus.Logger) (*clientv3.DeleteResponse, error) {

	// count keys about to be deleted
	gresp, err := cli.Get(ctx, "key", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	// delete the keys
	dresp, err := cli.Delete(ctx, "key", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	log.Println("Deleted all keys:", int64(len(gresp.Kvs)) == dresp.Deleted)
	// Output:
	// Deleted all keys: true
	return dresp, nil
}
