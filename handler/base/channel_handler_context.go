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

func (this *ChannelHandlerContext) findContextInbound() *ChannelHandlerContext {
	ctx := this.next
	for ; !ctx.next.inbound; ctx = ctx.next {
	}
	return ctx
}

func (this *ChannelHandlerContext) findContextOutbound() *ChannelHandlerContext {
	ctx := this.next
	for ; !ctx.next.outbound; ctx = ctx.next {
	}
	return ctx
}

func (this *ChannelHandlerContext) FireChannelRead(msg interface{}) {
	next := this.findContextInbound()
	handler := next.handler
	if handler != nil {
		if inboundHandler, ok := (*handler).(*ChannelInboundHandler); ok {
			(*inboundHandler).ChannelRead(next, msg)
		}
		next.FireChannelRead(msg)
	}
}

func (this *ChannelHandlerContext) FireChannelReadComplete() {

}

func (this *ChannelHandlerContext) FireChannelConnected() {

}

func (this *ChannelHandlerContext) FireChannelDisconnected() {

}

func (this *ChannelHandlerContext) FireUserEventTriggered(evnt interface{}) {

}

func (this *ChannelHandlerContext) FireExceptionCaught(err error) {

}

func (this *ChannelHandlerContext) write(msg interface{}) {

}
