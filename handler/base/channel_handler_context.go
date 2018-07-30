package base

type ChannelHandlerContext struct {
	prev     *ChannelHandlerContext
	next     *ChannelHandlerContext
	handler  *ChannelHandler
	pipeline *ChannelHandlerPipeline
	inbound  bool
	outbound bool
	name     string
}

func NewChannelHandlerContext(name string, handler *ChannelHandler, pipeline *ChannelHandlerPipeline) *ChannelHandlerContext {
	return &ChannelHandlerContext{name: name, handler: handler, pipeline: pipeline, inbound: IsInbound(handler), outbound: IsOutbound(handler)}
}

func IsInbound(handler ChannelHandler) bool {
	_, ok := handler.(ChannelInboundHandler)
	return ok
}

func IsOutbound(handler ChannelHandler) bool {
	_, ok := handler.(ChannelOutboundHandler)
	return ok
}

func (this *ChannelHandlerContext) getHandler() *ChannelHandler {
	return this.handler
}
