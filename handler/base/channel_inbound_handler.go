package base

type ChannelInboundHandler interface {
	ChannelHandler
	channelRead(ctx *ChannelHandlerContext, data interface{}) error
	channelReadComplete(ctx *ChannelHandlerContext)
	channelConnected(ctx *ChannelHandlerContext)
	channelDisconnected(ctx *ChannelHandlerContext)
	userEventTriggered(ctx *ChannelHandlerContext, evnt interface{})
	exceptionCaught(ctx *ChannelHandlerContext, err error)
}
