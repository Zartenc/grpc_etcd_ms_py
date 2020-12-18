package main

import (
	"context"
	"fmt"
	"go/client_center"
	go_protoc "go/grpc_proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	grpcClient := client_center.NewGrpcClient("xxxx", 2379, "/zarten")
	ip := grpcClient.GetRrpcServIp()
	fmt.Println(ip)

	conn, err := grpc.Dial(grpcClient.GetRrpcServIp(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := go_protoc.NewZartenClient(conn)
	res, _ := client.GetInfo(context.Background(), &go_protoc.ZartenRequest{ZhihuName:"zarten456"})
	fmt.Println(res.Name)
	fmt.Println(res.Homepage)

}
