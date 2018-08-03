package base

type ChannelInboundHandler interface {
	ChannelHandler
	ChannelRead(ctx *ChannelHandlerContext, data interface{})
	ChannelReadComplete(ctx *ChannelHandlerContext)
	ChannelConnected(ctx *ChannelHandlerContext)
	ChannelDisconnected(ctx *ChannelHandlerContext)
	UserEventTriggered(ctx *ChannelHandlerContext, evnt interface{})
	ExceptionCaught(ctx *ChannelHandlerContext, err error)
}
