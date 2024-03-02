package rice

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdClient(endpoints []string) (*clientv3.Client, error) {
	if len(endpoints) == 0 {
		endpoints = []string{"localhost:2379"}
	}
	return clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
}
