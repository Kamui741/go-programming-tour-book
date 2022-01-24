/*
 * @Author: ChZheng
 * @Date: 2022-01-22 00:45:23
 * @LastEditTime: 2022-01-23 01:03:26
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /tag-service/main.go
 */
package main

import (
	"google.golang.org/grpc/reflection"

	pb "go-programming-tour-book/tag-service/proto"
	"go-programming-tour-book/tag-service/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.Server err: %v", err)
	}
}
