package base

import "errors"

type ChannelHandlerPipeline struct {
	head *ChannelHandlerContext
	tail *ChannelHandlerContext
}

func NewHandlerChain() *ChannelHandlerPipeline {
	head := newHead()
	tail := newTail()
	ptr := &ChannelHandlerPipeline{head: head, tail: tail}
	ptr.head.next = tail
	ptr.tail.prev = head
	return ptr
}

func (this *ChannelHandlerPipeline) AddLast(handler *ChannelHandler) error {
	if handler == nil {
		return errors.New("")
	}
	newCtx := NewChannelHandlerContext("", handler, this)
	newCtx.prev = this.tail.prev
	newCtx.next = this.tail
	this.tail.prev.next = newCtx
	this.tail.prev = newCtx
	return nil
}

func newHead() *ChannelHandlerContext {
	return &ChannelHandlerContext{}
}

func newTail() *ChannelHandlerContext {
	return &ChannelHandlerContext{}
}

func (this *ChannelHandlerPipeline) FireChannelRead(msg interface{}) {
	this.head.FireChannelRead(msg)
}

func (this *ChannelHandlerPipeline) FireChannelReadComplete() {

}

func (this *ChannelHandlerPipeline) FireChannelConnected() {

}

func (this *ChannelHandlerPipeline) FireChannelDisconnected() {

}

func (this *ChannelHandlerPipeline) FireUserEventTriggered(evnt interface{}) {

}

func (this *ChannelHandlerPipeline) FireExceptionCaught(err error) {

}

func (this *ChannelHandlerPipeline) write(msg interface{}) {

}
