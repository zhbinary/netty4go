package types

import (
	"net"
)

type ChannelOutboundHandler interface {
	ChannelHandler
	Bind(ctx ChannelHandlerContext, localAddress net.Addr, promise ChannelPromise)
	Connect(ctx ChannelHandlerContext, localAddress net.Addr, remoteAddress net.Addr, promise ChannelPromise)
	Disconnect(ctx ChannelHandlerContext, promise ChannelPromise)
	Close(ctx ChannelHandlerContext, promise ChannelPromise)
	Deregister(ctx ChannelHandlerContext, promise ChannelPromise)
	Write(ctx ChannelHandlerContext, msg interface{}, promise ChannelPromise)
	Read(ctx ChannelHandlerContext)
	Flush(ctx ChannelHandlerContext)
}
