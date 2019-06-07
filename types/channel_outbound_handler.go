package types

import (
	"net"
)

type ChannelOutboundHandler interface {
	ChannelHandler
	Bind(ctx ChannelHandlerContext, localAddress net.Addr, promise Promise)
	Connect(ctx ChannelHandlerContext, localAddress net.Addr, remoteAddress net.Addr, promise Promise)
	Disconnect(ctx ChannelHandlerContext, promise Promise)
	Close(ctx ChannelHandlerContext, promise Promise)
	Deregister(ctx ChannelHandlerContext, promise Promise)
	Write(ctx ChannelHandlerContext, msg interface{}, promise Promise)
	Read(ctx ChannelHandlerContext)
	Flush(ctx ChannelHandlerContext)
}
