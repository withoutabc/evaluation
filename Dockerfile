FROM quay.io/coreos/etcd:latest
MAINTAINER ETCD
EXPOSE 2379 2380 4001 7001
ENTRYPOINT ["/usr/local/bin/etcd"]
CMD ["--name", "ETCD", "--data-dir", "/etcd-data", "--advertise-client-urls", "http://0.0.0.0:2379,http://0.0.0.0:4001", "--listen-client-urls", "http://0.0.0.0:2379,http://0.0.0.0:4001", "--initial-advertise-peer-urls", "http://0.0.0.0:2380", "--listen-peer-urls", "http://0.0.0.0:2380", "--initial-cluster", "my-etcd-1=http://0.0.0.0:2380", "--initial-cluster-token", "my-etcd-token", "--initial-cluster-state", "new"]