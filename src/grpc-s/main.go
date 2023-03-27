package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	pd "test/src/grpc-s/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pd.UnimplementedHiServer
}

func (*server) Hi(ctx context.Context, req *pd.HiRequest) (*pd.HiResponse, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, errors.New("请传输凭证")
    }
    var(
        appid string
        key string
    )
    if v, err := md["appid"]; err {
        appid = v[0]
    }
    if v, err := md["key"]; err {
        key = v[0]
    }
    if appid != "twb" || key != "123" {
        return nil, errors.New("凭证错误")
    }
	return &pd.HiResponse{JsonMsg: "成功接收" + req.RequestName}, nil
}

func main() {
    listen, _ := net.Listen("tcp", ":8081")
    grpcServer := grpc.NewServer()

    pd.RegisterHiServer(grpcServer, &server{})

    fmt.Println("监听成功")


    grpcServer.Serve(listen)


}



/**
安装proto后,还需要安装核心库,编译器
go get google.golang.org/grpc
安装go语言生成器
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest




syntax = "proto3";
option go_package =".;service";
service SayHello {
    rpc SayHello(HelloRequest) returns (HelloResponse) {}
}
message HelloRequest {
    string requestName = 1;
}
message HelloResponse {
    string responseMsg = 1;
}


protoc --go_out=. he11o.proto
protoc --go-grpc_out=. he11o.proto
**/
