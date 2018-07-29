package base

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

func (this *ChannelHandlerChain) FireRead(bytes []byte) {
	var data interface{}
	data = bytes
	for e := this.inboundList.Front(); e != nil; e = e.Next() {
		ret, err := (e.Value).(ChannelInboundHandler).channelRead(data)
		if err != nil {

		}
		data = ret
	}
}

func (this *ChannelHandlerChain) FireWrite(msg interface{}) {
	var data interface{}
	data = msg
	for e := this.inboundList.Front(); e != nil; e = e.Next() {
		ret, err := (e.Value).(ChannelOutboundHandler).channelWrite(data)
		if err != nil {

		}
		data = ret
	}
}


