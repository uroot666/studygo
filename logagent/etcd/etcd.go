package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: 5 * timeout,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err: %v\n", err)
		return
	}
	return
}

func Getconf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err: %v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed, %v\n", err)
			return
		}
	}
	return
}

// 用于向etcd发送测试数据
// func PutTest() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	value := `[{"path":"/c.log","topic":"c"},{"path":"/b.log","topic":"b"}]`
// 	_, err := cli.Put(ctx, "/xxx", value)
// 	cancel()
// 	if err != nil {
// 		fmt.Printf("put to etcd failed, err: %v\n", err)
// 		return
// 	}
// }
