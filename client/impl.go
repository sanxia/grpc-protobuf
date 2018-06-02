package main

import (
	//"fmt"
	"log"
	//"reflect"
	//"strings"
)

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

import (
	"github.com/sanxia/grpc-protobuf/protobuf"
	"github.com/sanxia/grpclib"
)

/* ================================================================================
 * Client Impl
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rpc Client
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type rpcClientImpl struct {
	*grpclib.RpcClient
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化Rpc Client
 * host: 主机地址
 * port: 主机端口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewRpcClientImpl(host string, port int, token string, args ...credentials.TransportCredentials) (*rpcClientImpl, error) {
	rpcClient, err := grpclib.NewRpcClient(host, port, token, args...)
	c := &rpcClientImpl{
		RpcClient: rpcClient,
	}

	return c, err
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 调用Rpc Song方法
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *rpcClientImpl) Song(title string) {
	rs, err := s.GetClient("action").(protobuf.ActionClient).Song(context.Background(), &protobuf.SongRequest{Title: title})
	if err != nil {
		log.Fatalf("could not add err: %v", err)
	}

	log.Printf("grpc: %s, %d", rs.Lyric, rs.Year)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 调用Rpc MakeAWish方法
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *rpcClientImpl) MakeAWish(from, to, content string) {
	rs, err := s.GetClient("action").(protobuf.ActionClient).MakeAWish(context.Background(), &protobuf.MakeAWishRequest{From: from, To: to, Content: content})
	if err != nil {
		log.Fatalf("could not echo err: %v", err)
	}

	resp, err := rs.Recv()
	if err != nil {
		log.Printf("rpc Echo err: %v", err)
	}

	log.Printf("grpc: %s", resp.Content)
}
