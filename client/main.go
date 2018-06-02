package main

import (
	"log"
)

import (
	"github.com/sanxia/grpc-protobuf/protobuf"
	"github.com/sanxia/grpclib"
)

/* ================================================================================
 * Main
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

const (
	rpcHost = "127.0.0.1"
	rpcPort = 5111
	token   = "test-token-2018"
)

func main() {
	//ssl证书
	tlsCreds, _ := grpclib.GetTLSCredentials("", "")

	//实例化RpcClientImpl
	rpcClientImpl, err := NewRpcClientImpl(rpcHost, rpcPort, token, tlsCreds)
	if err != nil {
		log.Fatalf("Failed connection rpc server: %v", err)
		return
	}

	//注册Rpc客户端服务
	rpcClientImpl.RegisterClient("action", protobuf.NewActionClient)
	defer rpcClientImpl.Close()

	//调用Rpc方法
	rpcClientImpl.Song("圣诞歌")
	rpcClientImpl.MakeAWish("老刘", "亲们", "元节快乐!")
	rpcClientImpl.MakeAWish("老刘", "才子佳人", "圣诞节快乐!")
}
