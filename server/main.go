package main

import (
	"log"
)

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
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
)

func main() {
	log.Printf("rpc server on %s:%d start...\r\n", rpcHost, rpcPort)
	RpcServer(rpcHost, rpcPort)
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rpc服务器
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func RpcServer(host string, port int) {
	//ssl证书
	creds, _ := grpclib.GetTLSCredentials("", "")

	//rpc server new
	rpcServer := grpclib.NewRpcServer(host, port, creds)

	//register rpc service
	rpcServer.RegisterService(func(rpcSrv *grpc.Server) {
		protobuf.RegisterActionServer(rpcSrv, &serverImpl{})
	})

	//authorize enabled
	rpcServer.Authorize(authorize)

	//logger enabled
	rpcServer.Logger()

	rpcServer.Serve()
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Rpc服务验证
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func authorize(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}
	grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))

	newCtx := context.WithValue(ctx, "token.info", tokenInfo)
	return newCtx, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Token解析
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Token的用户标识
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func userClaimFromToken(struct{}) string {
	return "test123"
}
