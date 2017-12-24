package main

import (
	"fmt"
	"time"
)

import (
	"golang.org/x/net/context"
)

import (
	"github.com/sanxia/grpc-protobuf/protobuf"
)

/* ================================================================================
 * Oauth Qq
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊
 * ================================================================================ */

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实现RPC OperatorServer接口
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
type serverImpl struct{}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Add
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *serverImpl) Song(ctx context.Context, in *protobuf.SongRequest) (*protobuf.SongReply, error) {
	lyric := fmt.Sprintf("唱一首%s\x1b[%dm%s\x1b[0m", in.Title, uint8(33), "Meeey Christmas")
	var year uint32 = 2018

	time.Sleep(time.Duration(450 * time.Millisecond))

	return &protobuf.SongReply{Lyric: lyric, Year: year}, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * Echo
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *serverImpl) MakeAWish(req *protobuf.MakeAWishRequest, server protobuf.Action_MakeAWishServer) error {
	from := fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(31), req.From)
	to := fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(32), req.To)
	content := fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(33), req.Content)
	info := fmt.Sprintf("\x1b[%dm%s\x1b[0m", uint8(36), " 用彩色的蜡笔，祝福 ")

	replyContent := fmt.Sprintf("%s%s%s：%s", from, info, to, content)

	makeAWishReply := &protobuf.MakeAWishReply{
		Content: replyContent,
	}

	time.Sleep(time.Duration(550 * time.Millisecond))

	return server.Send(makeAWishReply)
}
