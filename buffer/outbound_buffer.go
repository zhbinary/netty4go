//Created by zhbinary on 2019-06-19.
//Email: zhbinary@gmail.com
package buffer

import (
	"container/list"
	"github.com/zhbinary/heng/types"
)

type OutboundBuffer struct {
	list  list.List
	front *list.Element
}

func (this *OutboundBuffer) AddMessage(msg interface{}, promise types.ChannelPromise) {
	this.list.PushBack(&OutboundEntry{Msg: msg, Promise: promise})
}

func (this *OutboundBuffer) Front() interface{} {
	this.front = this.list.Front()
	if this.front == nil {
		return nil
	}
	return this.front.Value
}

func (this *OutboundBuffer) RemoveFront() {
	if this.front != nil {
		this.list.Remove(this.front)
	}
}

type OutboundEntry struct {
	Msg     interface{}
	Promise types.ChannelPromise
}
