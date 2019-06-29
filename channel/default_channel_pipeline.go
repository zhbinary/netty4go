package channel

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type DefaultChannelPipeline struct {
	channel types.Channel
	head    *AbstractChannelHandlerContext
	tail    *AbstractChannelHandlerContext

	OnUnhandledChannelRead func(msg interface{})
}

func NewDefaultChannelPipeline(channel types.Channel) *DefaultChannelPipeline {
	//if Channel == nil {
	//	
	//}
	pipeline := &DefaultChannelPipeline{channel: channel}
	pipeline.initAbstractMethods()
	head := newHead(pipeline).AbstractChannelHandlerContext
	tail := newTail(pipeline).AbstractChannelHandlerContext
	pipeline.head = head
	pipeline.tail = tail
	head.Next = tail
	tail.Prev = head
	return pipeline
}

func (this *DefaultChannelPipeline) initAbstractMethods() {
	this.OnUnhandledChannelRead = this.onUnhandledChannelRead
}

func (this *DefaultChannelPipeline) Channel() types.Channel {
	return this.channel
}

func (this *DefaultChannelPipeline) newContext(name string, handler types.ChannelHandler) *AbstractChannelHandlerContext {
	return NewDefaultChannelHandlerContext(name, handler, this)
}

func (this *DefaultChannelPipeline) FireChannelRegistered() {
	this.head.invokeChannelRegistered()
}

func (this *DefaultChannelPipeline) FireChannelUnregistered() {
	this.head.invokeChannelUnregistered()
}

func (this *DefaultChannelPipeline) FireChannelActive() {
	this.head.invokeChannelActive()
}

func (this *DefaultChannelPipeline) FireChannelInactive() {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireExceptionCaught(err error) {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireUserEventTriggered(evt interface{}) {
	panic("implement me")
}

func (this *DefaultChannelPipeline) FireChannelRead(msg interface{}) {
	this.head.invokeChannelRead(msg)
}

func (this *DefaultChannelPipeline) FireChannelReadComplete() {
	this.head.invokeChannelReadComplete()
}

func (this *DefaultChannelPipeline) FireChannelWritabilityChanged() {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Bind(localAddress net.Addr) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Bind0(localAddress net.Addr, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Connect(localAddress net.Addr, remoteAddress net.Addr) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Connect0(localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Disconnect() types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Disconnect0(promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Close() types.ChannelFutrue {
	return this.tail.Close()
}

func (this *DefaultChannelPipeline) Close0(promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Deregister() types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Deregister0(promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Read() {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Write(msg interface{}) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Write0(msg interface{}, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) WriteAndFlush(msg interface{}) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) WriteAndFlush0(msg interface{}, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *DefaultChannelPipeline) Flush() {
	panic("implement me")
}

func (this *DefaultChannelPipeline) AddLast(name string, handler types.ChannelHandler) {
	newCtx := this.newContext(name, handler)
	prev := this.tail.Prev
	newCtx.Prev = prev
	newCtx.Next = this.tail
	prev.Next = newCtx
	this.tail.Prev = newCtx
}

func (this *DefaultChannelPipeline) onUnhandledChannelRead(msg interface{}) {
	// Release message
}

type TailContext struct {
	*AbstractChannelHandlerContext
}

func newTail(pipeline *DefaultChannelPipeline) (ctx *TailContext) {
	return &TailContext{AbstractChannelHandlerContext: &AbstractChannelHandlerContext{pipeline: pipeline}}
}

func (this *TailContext) Handler() types.ChannelHandler {
	return this
}

func (this *TailContext) HandlerAdded(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) HandlerRemoved(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) ChannelRead(ctx types.ChannelHandlerContext, msg interface{}) {
	this.pipeline.OnUnhandledChannelRead(msg)
}

func (this *TailContext) ChannelReadComplete(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) ChannelActive(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) ChannelInactive(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) ChannelRegistered(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) ChannelUnregistered(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) ChannelWritabilityChanged(ctx types.ChannelHandlerContext) {
}

func (this *TailContext) UserEventTriggered(ctx types.ChannelHandlerContext, evt interface{}) {
}

func (this *TailContext) ExceptionCaught(ctx types.ChannelHandlerContext, err error) {
}

type HeadContext struct {
	*AbstractChannelHandlerContext
	unsafe types.Unsafe
}

func newHead(pipeline *DefaultChannelPipeline) (ctx *HeadContext) {
	return &HeadContext{AbstractChannelHandlerContext: &AbstractChannelHandlerContext{pipeline: pipeline}, unsafe: pipeline.channel.Unsafe()}
}

func (this *HeadContext) Handler() types.ChannelHandler {
	return this
}

func (this *HeadContext) HandlerAdded(ctx types.ChannelHandlerContext) {
}

func (this *HeadContext) HandlerRemoved(ctx types.ChannelHandlerContext) {
}

func (this *HeadContext) Bind(ctx types.ChannelHandlerContext, localAddress net.Addr, promise types.ChannelPromise) {
	this.unsafe.Bind(localAddress, promise)
}

func (this *HeadContext) Connect(ctx types.ChannelHandlerContext, localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) {
	this.unsafe.Connect(localAddress, remoteAddress, promise)
}

func (this *HeadContext) Disconnect(ctx types.ChannelHandlerContext, promise types.ChannelPromise) {
	this.unsafe.Disconnect(promise)
}

func (this *HeadContext) Close(ctx types.ChannelHandlerContext, promise types.ChannelPromise) {
	this.unsafe.Close(promise)
}

func (this *HeadContext) Deregister(ctx types.ChannelHandlerContext, promise types.ChannelPromise) {
	this.unsafe.Deregister(promise)
}

func (this *HeadContext) Write(ctx types.ChannelHandlerContext, msg interface{}, promise types.ChannelPromise) {
	this.unsafe.Write(msg, promise)
}

func (this *HeadContext) Read(ctx types.ChannelHandlerContext) {
	this.unsafe.BeginRead()
}

func (this *HeadContext) Flush(ctx types.ChannelHandlerContext) {
	this.unsafe.Flush()
}
