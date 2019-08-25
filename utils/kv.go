package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"strconv"

	"go.etcd.io/etcd/clientv3"
	// "go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
)

var (
	log = logrus.New()
)

func StringFlatedMap(fp map[string]interface{}) map[string]string {
	sfp := make(map[string]string)
	for k, v := range fp {
		switch vv := v.(type) {
		case string:
			val := v.(string)
			sfp[k] = val
		case bool:
			val := strconv.FormatBool(v.(bool))
			sfp[k] = val
		case float64:
			val := strconv.FormatFloat(v.(float64), 'f', -1, 64)
			sfp[k] = val
		default:
			log.Fatalf(k, "type known", vv)
		}
		// TODO
	}
	return sfp
}

func BatchStringFlatedMap(ctx context.Context, cli *clientv3.Client, sfp map[string]string, prefix string, log *logrus.Logger) error {
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

func Example() {}
