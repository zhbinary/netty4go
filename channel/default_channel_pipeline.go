package channel

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type DefaultChannelPipeline struct {
	channel types.Channel
	head    types.ChannelHandlerContext
	tail    types.ChannelHandlerContext
}

func NewDefaultChannelPipeline() *DefaultChannelPipeline {
	head := newHead()
	tail := newTail()
	ptr := &DefaultChannelPipeline{head: head, tail: tail}
	head.Next() = tail
	tail.Prev() = head
	return ptr
}

func (this *DefaultChannelPipeline) Channel() types.Channel {
	return this.channel
}

func (this *DefaultChannelPipeline) newContext(name string, handler types.ChannelHandler) types.ChannelHandlerContext {
	return nil
}

func newHead() types.ChannelHandlerContext {
	return nil
}

func newTail() types.ChannelHandlerContext {
	return nil
}

func (this *DefaultChannelPipeline) FireChannelRegistered() types.ChannelInboundInvoker {
	this.head.FireChannelRegistered()
	return this
}

func (this *DefaultChannelPipeline) FireChannelUnregistered() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireChannelActive() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireChannelInactive() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireExceptionCaught(err error) types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireUserEventTriggered(evt interface{}) types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireChannelRead(msg interface{}) types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireChannelReadComplete() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireChannelWritabilityChanged() types.ChannelInboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Bind(localAddress net.Addr) types.Future {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Connect(localAddress net.Addr, remoteAddress net.Addr) types.Future {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Disconnect() types.Future {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Close() types.Future {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Deregister() types.Future {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Write(msg interface{}) types.Future {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Flush() types.ChannelOutboundInvoker {
	panic("implement me")
}

func (this *DefaultChannelPipeline) AddLast(name string, handler types.ChannelHandler) {
	newCtx := this.newContext(name, handler)
	prev := this.tail.Prev()
	newCtx.Prev() = prev
	newCtx.Next() = this.tail
	prev.Next() = newCtx
	this.tail.Prev() = newCtx
}
