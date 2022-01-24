/*
 * @Author: ChZheng
 * @Date: 2022-01-24 23:36:12
 * @LastEditTime: 2022-01-24 23:50:56
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /tag-service/client/client.go
 */
package main

import (
	"context"
	"log"

	pb "go-programming-tour-book/tag-service/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:8004", nil)
	defer clientConn.Close()

	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "GO"})

	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
