package etcd

import "go.etcd.io/etcd/api/v3/mvccpb"

func KvsToCacheMap(kvs []*mvccpb.KeyValue) map[string]string {
	cacheMap := make(map[string]string)

	for _, val := range kvs {
		cacheMap[string(val.Key)] = string(val.Value)
	}

	return cacheMap
}
