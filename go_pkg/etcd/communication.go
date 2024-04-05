package etcd

import (
	"context"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	CacheTimeout = 5 * time.Second
	SetupTimeout = 15 * time.Second
)

type Communication interface {
	Get(key string) (string, error)
	Put(key, value string) error
}

type communication struct {
	cache  *Cache
	client *clientv3.Client
}

func InitCommunication(addr, user, password string) (Communication, error) {
	var err error
	comm := &communication{}

	etcdCfg := clientv3.Config{Endpoints: []string{addr}, DialTimeout: CacheTimeout, Username: user, Password: password}
	comm.client, err = clientv3.New(etcdCfg)
	if err != nil {
		return nil, errors.New("cant start ETCD: " + err.Error())
	}

	comm.cache = NewCache()

	if err = comm.Setup(); err != nil {
		return nil, err
	}

	return comm, nil
}

func (comm *communication) Get(key string) (string, error) {
	return comm.cache.Get(key)
}

func (comm *communication) Put(key, value string) error {
	prev, err := comm.cache.Get(key)
	if err != nil {
		return err
	}

	if err = comm.cache.Put(key, value); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), SetupTimeout)
	defer cancel()

	if _, err = comm.client.Put(ctx, key, value); err != nil {
		if err2 := comm.cache.Put(key, prev); err2 != nil {
			return errors.Join(err, err2)
		}
		return err
	}

	return err
}

func (comm *communication) Setup() error {
	ctx, cancel := context.WithTimeout(context.Background(), SetupTimeout)
	defer cancel()

	resp2, err := comm.client.Get(ctx, "", clientv3.WithPrefix())
	if err != nil {
		return err
	}

	return comm.cache.Update(KvsToCacheMap(resp2.Kvs))
}

func (comm *communication) Watch(ctx context.Context) error {
	channel := comm.client.Watch(ctx, "", clientv3.WithPrefix())

	for resp := range channel {
		if resp.Err() != nil {
			return resp.Err()
		}
		fmt.Println(resp.Events)
	}

	return nil
}
