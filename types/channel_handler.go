package types

type ChannelHandler interface {
	HandlerAdded(ctx ChannelHandlerContext)
	HandlerRemoved(ctx ChannelHandlerContext)
}
