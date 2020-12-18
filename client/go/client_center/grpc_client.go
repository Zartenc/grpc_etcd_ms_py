package client_center

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
type GrpcClient struct {
	Etcd3Client *clientv3.Client
	GrpcEndpoints []string
}
var mu sync.Mutex

func NewGrpcClient(EtcdIp string, EtcdPort int, EtcdPrefix string) *GrpcClient{
	keyName := EtcdPrefix+"/grpc"

	grpcClient := new(GrpcClient)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{EtcdIp + ":" + strconv.Itoa(EtcdPort)},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	res, err := cli.Get(context.Background(), keyName)
	if err != nil{
		log.Fatal(err)
	}
	for _, ev := range res.Kvs {
		endPoints := ev.Value
		err := json.Unmarshal(endPoints, &grpcClient.GrpcEndpoints)
		if err != nil{
			log.Fatal(err)
		}
		break
	}
	if len(grpcClient.GrpcEndpoints) <= 0{
		log.Fatal("No grpc services are available.Please notify the administrator to start the grpc service")
	}

	grpcClient.Etcd3Client = cli

	rch := cli.Watch(context.Background(), keyName)
	go func() {
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				mu.Lock()
				err := json.Unmarshal(ev.Kv.Value, &grpcClient.GrpcEndpoints)
				mu.Unlock()
				if err != nil{
					log.Fatal(err)
				}
				fmt.Println(grpcClient.GrpcEndpoints)
			}
		}
	}()

	return grpcClient
}

func (g *GrpcClient) GetRrpcServIp() string{
	rand.Seed(time.Now().Unix())
	n := len(g.GrpcEndpoints)
	return g.GrpcEndpoints[rand.Intn(n)]
}
