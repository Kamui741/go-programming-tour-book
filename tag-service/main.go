/*
 * @Author: ChZheng
 * @Date: 2022-01-22 00:45:23
 * @LastEditTime: 2022-02-03 17:35:01
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /tag-service/main.go
 */
package main

import (
	"context"
	"encoding/json"
	"flag"
	"path"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"go-programming-tour-book/tag-service/pkg/swagger"

	"net/http"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2/h2c"

	"golang.org/x/net/http2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "go-programming-tour-book/tag-service/proto"
	"go-programming-tour-book/tag-service/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port string

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func init() {
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()
}
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
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	// endpoint := "0.0.0.0" + port
	// runtime.HTTPError = grpcGatewayError
	// gwmux := runtime.NewServeMux()

	gatewayMux := runGrpcGatewayServer()
	httpMux.Handle("/", gatewayMux)
	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)
		http.ServeFile(w, r, p)
	})
	return serveMux
}
func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			HelloInterceptor,
			WorldInterceptor,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}
func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)
	return gwmux
}
func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}
	httpError := httpError{Code: int32(s.Code()), Message: s.Message()}
	datails := s.Details()
	for _, detail := range datails {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}
	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-Type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}

func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Hello")
	resp, err := handler(ctx, req)
	log.Println("bye")
	return resp, err
}
func WorldInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("Hello")
	resp, err := handler(ctx, req)
	log.Println("bye")
	return resp, err
}
