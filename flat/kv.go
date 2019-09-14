package flat

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
)



func BatchMap(ctx context.Context, cli *clientv3.Client, m map[string]interface{}, prefix string) error {
	defer cli.Close() // make sure to close the client
	// Flat
	fp := make(map[string]interface{})
	FlatMap(m, fp, "/", "")
	sfp := StringFlatedMap(fp)
	// Batch
	err := BatchStringFlatedMap(context.TODO(), cli, sfp, prefix)
	if err != nil {
		return err
	}
	return nil
}

func BatchStringFlatedMap(ctx context.Context, cli *clientv3.Client, sfp map[string]string, prefix string) error {
	defer cli.Close() // make sure to close the client
	for k, v := range sfp {
		_, err := cli.Put(ctx, prefix+k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func KV_getAnddelete(ctx context.Context, cli *clientv3.Client, key string, val string) (*clientv3.DeleteResponse, error) {
	defer cli.Close() // make sure to close the client

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
