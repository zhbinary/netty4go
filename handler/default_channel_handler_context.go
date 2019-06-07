package handler

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type DefaultChannelHandlerContext struct {
	prev     *DefaultChannelHandlerContext
	next     *DefaultChannelHandlerContext
	handler  types.ChannelHandler
	pipeline types.ChannelPipeline
	inbound  bool
	outbound bool
	name     string
}

func NewDefaultChannelHandlerContext(name string, handler types.ChannelHandler, pipeline types.ChannelPipeline) types.ChannelHandlerContext {
	return &DefaultChannelHandlerContext{name: name, handler: handler, pipeline: pipeline, inbound: IsInbound(handler), outbound: IsOutbound(handler)}
}

func (this *DefaultChannelHandlerContext) findContextInbound() types.ChannelHandlerContext {
	ctx := this.Next()
	for ; !ctx.Next().IsInbound(); ctx = ctx.Next() {
	}
	return ctx
}

func (this *DefaultChannelHandlerContext) findContextOutbound() types.ChannelHandlerContext {
	ctx := this.Next()
	for ; !ctx.Next().IsOutbound(); ctx = ctx.Next() {
	}
	return ctx
}

func (this *DefaultChannelHandlerContext) FireChannelRegistered() types.ChannelInboundInvoker {
	this.Channel().EventLoop().Execute(func() {
		if inHandler, ok := this.findContextInbound().Handler().(types.ChannelInboundHandler); ok {
			// todo pass next handler context, not this
			inHandler.ChannelRegistered(this)
		} else {

		}
	}, nil)
	return this
}

func (this *DefaultChannelHandlerContext) FireChannelUnregistered() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireChannelActive() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireChannelInactive() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireExceptionCaught(err error) types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireUserEventTriggered(evt interface{}) types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireChannelRead(msg interface{}) types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireChannelReadComplete() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) FireChannelWritabilityChanged() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Channel() types.Channel {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Name() string {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Handler() types.ChannelHandler {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) IsRemoved() bool {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Pipeline() types.ChannelPipeline {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Bind(localAddress net.Addr, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Connect(localAddress net.Addr, remoteAddress net.Addr, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Disconnect(promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Close(promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Deregister(promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Write(msg interface{}, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) WriteAndFlush(msg interface{}, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Flush() {
	panic("implement me")
}

func (this *DefaultChannelHandlerContext) Next() types.ChannelHandlerContext {
	return this.next
}

func (this *DefaultChannelHandlerContext) Prev() types.ChannelHandlerContext {
	return this.prev
}

func (this *DefaultChannelHandlerContext) IsInbound() bool {
	return this.inbound
}

func (this *DefaultChannelHandlerContext) IsOutbound() bool {
	return this.outbound
}

func IsInbound(handler types.ChannelHandler) bool {
	_, ok := handler.(types.ChannelInboundHandler)
	return ok
}

func IsOutbound(handler types.ChannelHandler) bool {
	_, ok := handler.(types.ChannelOutboundHandler)
	return ok
}
