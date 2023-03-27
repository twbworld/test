package main

import (
	"context"
	"fmt"
	"log"
	pd "test/src/grpc-s/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type cre struct{
}


func (cre)GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error){
    return map[string]string{
        "appid": "twba",
        "key": "123",
    }, nil
}

func (cre)RequireTransportSecurity() bool{
    return false
}


func main() {

    var  arr []grpc.DialOption

    arr = append(arr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    arr = append(arr, grpc.WithPerRPCCredentials(new(cre)))


	listen, ok := grpc.Dial("127.0.0.1:8081", arr...)

    if ok != nil {
        log.Fatalf("出现错误: %s", ok)
    }
	defer listen.Close()

	pdc := pd.NewHiClient(listen)

	rep, ok := pdc.Hi(context.Background(), &pd.HiRequest{RequestName: "server陛下"})

    if ok != nil {
        log.Fatalf("出现错误: %s", ok)
    }

	fmt.Println(rep.GetJsonMsg())

}
