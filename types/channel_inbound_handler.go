package types

type ChannelInboundHandler interface {
	ChannelHandler
	ChannelRead(ctx ChannelHandlerContext, msg interface{})
	ChannelReadComplete(ctx ChannelHandlerContext)
	ChannelActive(ctx ChannelHandlerContext)
	ChannelInactive(ctx ChannelHandlerContext)
	ChannelRegistered(ctx ChannelHandlerContext)
	ChannelUnregistered(ctx ChannelHandlerContext)
	ChannelWritabilityChanged(ctx ChannelHandlerContext)
	UserEventTriggered(ctx ChannelHandlerContext, evt interface{})
	ExceptionCaught(ctx ChannelHandlerContext, err error)
}
