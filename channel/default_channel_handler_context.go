package channel

import (
	"github.com/zhbinary/heng/concurrent"
	"github.com/zhbinary/heng/types"
	"net"
)

type AbstractChannelHandlerContext struct {
	Prev     *AbstractChannelHandlerContext
	Next     *AbstractChannelHandlerContext
	handler  types.ChannelHandler
	pipeline *DefaultChannelPipeline
	inbound  bool
	outbound bool
	name     string
}

func NewDefaultChannelHandlerContext(name string, handler types.ChannelHandler, pipeline *DefaultChannelPipeline) *AbstractChannelHandlerContext {
	return &AbstractChannelHandlerContext{name: name, handler: handler, pipeline: pipeline, inbound: IsInbound(handler), outbound: IsOutbound(handler)}
}

func (this *AbstractChannelHandlerContext) findContextInbound() *AbstractChannelHandlerContext {
	ctx := this.Next
	for ; ctx != nil && !ctx.IsInbound(); ctx = ctx.Next {
	}
	return ctx
}

func (this *AbstractChannelHandlerContext) findContextOutbound() *AbstractChannelHandlerContext {
	ctx := this.Prev
	for ; ctx != nil && !ctx.IsOutbound(); ctx = ctx.Prev {
	}
	return ctx
}

func (this *AbstractChannelHandlerContext) FireChannelRegistered() {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelRegistered()
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelRegistered() {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelRegistered(this)
	}
}

func (this *AbstractChannelHandlerContext) FireChannelUnregistered() {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelUnregistered()
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelUnregistered() {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelUnregistered(this)
	}
}

func (this *AbstractChannelHandlerContext) FireChannelActive() {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelActive()
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelActive() {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelActive(this)
	}
}

func (this *AbstractChannelHandlerContext) FireChannelInactive() {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelInactive()
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelInactive() {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelInactive(this)
	}
}

func (this *AbstractChannelHandlerContext) FireExceptionCaught(err error) {
	if next := this.findContextInbound(); next != nil {
		next.invokeExceptionCaught(err)
	}
}

func (this *AbstractChannelHandlerContext) invokeExceptionCaught(err error) {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ExceptionCaught(this, err)
	}
}

func (this *AbstractChannelHandlerContext) FireUserEventTriggered(evt interface{}) {
	if next := this.findContextInbound(); next != nil {
		next.invokeUserEventTriggered(evt)
	}
}

func (this *AbstractChannelHandlerContext) invokeUserEventTriggered(evt interface{}) {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.UserEventTriggered(this, evt)
	}
}

func (this *AbstractChannelHandlerContext) FireChannelRead(msg interface{}) {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelRead(msg)
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelRead(msg interface{}) {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelRead(this, msg)
	}
}

func (this *AbstractChannelHandlerContext) FireChannelReadComplete() {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelReadComplete()
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelReadComplete() {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelReadComplete(this)
	}
}

func (this *AbstractChannelHandlerContext) FireChannelWritabilityChanged() {
	if next := this.findContextInbound(); next != nil {
		next.invokeChannelWritabilityChanged()
	}
}

func (this *AbstractChannelHandlerContext) invokeChannelWritabilityChanged() {
	if inHandler, ok := this.handler.(types.ChannelInboundHandler); ok {
		inHandler.ChannelWritabilityChanged(this)
	}
}

func (this *AbstractChannelHandlerContext) Channel() types.Channel {
	return this.pipeline.Channel()
}

func (this *AbstractChannelHandlerContext) Name() string {
	return this.name
}

func (this *AbstractChannelHandlerContext) Handler() types.ChannelHandler {
	return this.handler
}

func (this *AbstractChannelHandlerContext) IsRemoved() bool {
	panic("implement me")
}

func (this *AbstractChannelHandlerContext) Pipeline() types.ChannelPipeline {
	return this.pipeline
}

func (this *AbstractChannelHandlerContext) Bind(localAddress net.Addr) types.ChannelFutrue {
	return this.Bind0(localAddress, this.newPromise())
}

func (this *AbstractChannelHandlerContext) Bind0(localAddress net.Addr, promise types.ChannelPromise) (future types.ChannelFutrue) {
	if next := this.findContextOutbound(); next != nil {
		next.invokeBind(localAddress, promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeBind(localAddress net.Addr, promise types.ChannelPromise) {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Bind(this, localAddress, promise)
	}
}

func (this *AbstractChannelHandlerContext) Connect(localAddress net.Addr, remoteAddress net.Addr) types.ChannelFutrue {
	return this.Connect0(localAddress, remoteAddress, this.newPromise())
}

func (this *AbstractChannelHandlerContext) Connect0(localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) types.ChannelFutrue {
	if next := this.findContextOutbound(); next != nil {
		next.invokeConnect(localAddress, remoteAddress, promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeConnect(localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Connect(this, localAddress, remoteAddress, promise)
	}
}

func (this *AbstractChannelHandlerContext) Disconnect() types.ChannelFutrue {
	return this.Disconnect0(this.newPromise())
}

func (this *AbstractChannelHandlerContext) Disconnect0(promise types.ChannelPromise) types.ChannelFutrue {
	if next := this.findContextOutbound(); next != nil {
		next.invokeDisconnect(promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeDisconnect(promise types.ChannelPromise) {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Disconnect(this, promise)
	}
}

func (this *AbstractChannelHandlerContext) Close() (future types.ChannelFutrue) {
	return this.Close0(this.newPromise())
}

func (this *AbstractChannelHandlerContext) Close0(promise types.ChannelPromise) types.ChannelFutrue {
	if next := this.findContextOutbound(); next != nil {
		next.invokeClose(promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeClose(promise types.ChannelPromise) {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Close(this, promise)
	}
}

func (this *AbstractChannelHandlerContext) Deregister() types.ChannelFutrue {
	return this.Deregister0(this.newPromise())
}

func (this *AbstractChannelHandlerContext) Deregister0(promise types.ChannelPromise) types.ChannelFutrue {
	if next := this.findContextOutbound(); next != nil {
		next.invokeDeregister(promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeDeregister(promise types.ChannelPromise) {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Deregister(this, promise)
	}
}

func (this *AbstractChannelHandlerContext) Read() {
	if next := this.findContextOutbound(); next != nil {
		next.invokeRead()
	}
}

func (this *AbstractChannelHandlerContext) invokeRead() {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Read(this)
	}
}

func (this *AbstractChannelHandlerContext) Write(msg interface{}) types.ChannelFutrue {
	return this.Write0(msg, this.newPromise())
}

func (this *AbstractChannelHandlerContext) Write0(msg interface{}, promise types.ChannelPromise) types.ChannelFutrue {
	if next := this.findContextOutbound(); next != nil {
		next.invokeWrite(msg, promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeWrite(msg interface{}, promise types.ChannelPromise) {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Write(this, msg, promise)
	}
}

func (this *AbstractChannelHandlerContext) WriteAndFlush(msg interface{}) types.ChannelFutrue {
	return this.WriteAndFlush0(msg, this.newPromise())
}

func (this *AbstractChannelHandlerContext) WriteAndFlush0(msg interface{}, promise types.ChannelPromise) types.ChannelFutrue {
	if next := this.findContextOutbound(); next != nil {
		next.invokeWrite(msg, promise)
	}
	return promise
}

func (this *AbstractChannelHandlerContext) invokeWriteAndFlush(msg interface{}, promise types.ChannelPromise) {
	this.invokeWrite(msg, promise)
	this.invokeFlush()
}

func (this *AbstractChannelHandlerContext) Flush() {
	if next := this.findContextOutbound(); next != nil {
		next.invokeFlush()
	}
}

func (this *AbstractChannelHandlerContext) invokeFlush() {
	if outHandler, ok := this.handler.(types.ChannelOutboundHandler); ok {
		outHandler.Flush(this)
	}
}

func (this *AbstractChannelHandlerContext) IsInbound() bool {
	return this.inbound
}

func (this *AbstractChannelHandlerContext) IsOutbound() bool {
	return this.outbound
}

func (this *AbstractChannelHandlerContext) newPromise() types.ChannelPromise {
	return concurrent.NewDefaultChannelPromise(this.Channel())
}

func IsInbound(handler types.ChannelHandler) bool {
	_, ok := handler.(types.ChannelInboundHandler)
	return ok
}

func IsOutbound(handler types.ChannelHandler) bool {
	_, ok := handler.(types.ChannelOutboundHandler)
	return ok
}
