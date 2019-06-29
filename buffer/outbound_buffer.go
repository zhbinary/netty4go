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
	this.list.PushBack(&OutboundEntry{msg: msg, promise: promise})
}

func (this *OutboundBuffer) Front() interface{} {
	this.front = this.list.Front()
	return this.front.Value
}

func (this *OutboundBuffer) RemoveFront() {
	this.list.Remove(this.front)
}

type OutboundEntry struct {
	msg     interface{}
	promise types.ChannelPromise
}
