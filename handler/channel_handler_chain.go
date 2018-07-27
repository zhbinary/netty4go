package handler

import (
	"container/list"
)

type ChannelHandlerChain struct {
	inboundList  list.List
	outboundList list.List
}

func NewHandlerChain() *ChannelHandlerChain {
	return &ChannelHandlerChain{inboundList: list.List{}, outboundList: list.List{}}
}

func (this *ChannelHandlerChain) AddLast(handler ChannelHandler) {
	switch handler.(type) {
	case ChannelInboundHandler:
		this.inboundList.PushBack(handler)
		break
	case ChannelOutboundHandler:
		this.outboundList.PushBack(handler)
		break
	}
}

func (this *ChannelHandlerChain) ReadChain(bytes []byte) {
	for e := this.inboundList.Front(); e != nil; e = e.Next() {
		(e.Value).(ChannelInboundHandler).channelRead(bytes)
	}
}

func (this *ChannelHandlerChain) WriteChain(msg interface{}) {
	for e := this.inboundList.Front(); e != nil; e = e.Next() {
		(e.Value).(ChannelOutboundHandler).channelWrite(msg)
	}
}
