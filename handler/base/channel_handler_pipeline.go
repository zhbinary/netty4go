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

func (this *ChannelHandlerPipeline) FireChannelRead() {

}

//func (this *ChannelHandlerPipeline) AddLast(handler ChannelHandler) {
//	switch handler.(type) {
//	case ChannelInboundHandler:
//		this.inboundList.PushBack(handler)
//		break
//	case ChannelOutboundHandler:
//		this.outboundList.PushBack(handler)
//		break
//	}
//}
//
//func (this *ChannelHandlerPipeline) FireRead(bytes []byte) {
//	var data interface{}
//	data = bytes
//	for e := this.inboundList.Front(); e != nil; e = e.Next() {
//		ret, err := (e.Value).(ChannelInboundHandler).channelRead(data)
//		if err != nil {
//
//		}
//		data = ret
//	}
//}
//
//func (this *ChannelHandlerPipeline) FireWrite(msg interface{}) {
//	var data interface{}
//	data = msg
//	for e := this.inboundList.Front(); e != nil; e = e.Next() {
//		ret, err := (e.Value).(ChannelOutboundHandler).channelWrite(data)
//		if err != nil {
//
//		}
//		data = ret
//	}
//}
//
//func Fire(i interface{})  {
//
//}
