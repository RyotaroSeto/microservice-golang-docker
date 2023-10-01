package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	fmt.Println("Hello etcd")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{os.Getenv("ETCD_CLIENT_URL")},
		DialTimeout: 5 * time.Second,
	})
	fmt.Println("Hello etcd")
	if err != nil {
		fmt.Printf("Error etcd: %v\n", err)
		return
	}
	defer cli.Close()

	fmt.Println("Hello etcd")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = cli.Put(ctx, "/my-key/kv", "my-value")
	if err != nil {
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}

	resp, err := cli.Get(ctx, "/my-key/kv")
	if err != nil {
		log.Fatal(err)
	}
	if resp.Count == 0 {
		log.Fatal("/my-key/kv not found")
	}
	fmt.Println(string(resp.Kvs[0].Value))

	_, err = cli.Delete(context.TODO(), "/my-key/kv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted")
}
