# kafka
```
kafka-console-consumer.sh --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
```

# etcdh
```
set ETCDCTL_API=3
etcdctl --endpoints=http://127.0.0.1:2379 put liang biao
etcdctl --endpoints=http://127.0.0.1:2379 get liang

# 安装etcd需要这么指定不能直接使用clientv3
go get go.etcd.io/etcd/client/v3
```