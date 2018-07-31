package channel

import (
	"github.com/zhbinary/thor"
	"github.com/zhbinary/thor/handler/base"
)

type ServerChannel interface {
	Channel
	Bind(launcher *thor.ServerBootstrap, addr string)
	Close()
	ShutDown()
	GetPipeline() *base.ChannelHandlerPipeline
}
