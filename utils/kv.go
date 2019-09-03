package utils

import (
	"context"
	"github.com/sirupsen/logrus"

	"go.etcd.io/etcd/clientv3"
	// "go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

var (
	log = logrus.New()
)

func BatchStringFlatedMap(ctx context.Context, cli *clientv3.Client, sfp map[string]string, prefix string) error {
	for k, v := range sfp {
		_, err := cli.Put(ctx, prefix+k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func KV_getAnddelete(ctx context.Context, cli *clientv3.Client, key string, val string, log *logrus.Logger) (*clientv3.DeleteResponse, error) {

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
