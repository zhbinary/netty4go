//Created by zhbinary on 2019-06-18.
//Email: zhbinary@gmail.com
package embedded

import (
	"container/list"
	"github.com/zhbinary/heng/buffer"
	"github.com/zhbinary/heng/channel"
	"github.com/zhbinary/heng/types"
	"net"
)

const (
	StateOpen = iota
	StateActive
	StateClosed
)

type Channel struct {
	*channel.AbstractChannel
	handlers         []types.ChannelHandler
	state            int
	inboundMessages  list.List
	outboundMessages list.List
}

func NewChannel(handlers ...types.ChannelHandler) (ch *Channel) {
	ch = &Channel{state: StateOpen, handlers: handlers}
	unsafe := &Unsafe{AbstractUnsafe: channel.NewAbstractUnsafe(ch)}
	ch.AbstractChannel = channel.NewAbstractChannel("", NewEventLoop(), nil, unsafe)
	ch.initAbstractMethods()
	ch.setup()
	return ch
}

func (this *Channel) setup() {
	for _, handler := range this.handlers {
		this.Pipeline().AddLast("", handler)
	}
	this.EventLoop().Register(this)
}

func (this *Channel) initAbstractMethods() {
	pipeline := this.AbstractChannel.Pipeline().(*channel.DefaultChannelPipeline)
	pipeline.OnUnhandledChannelRead = this.onUnHandledChannelRead
}

//func (this *Channel) LocalAddress() net.Addr {
//	return nil
//}
//
//func (this *Channel) RemoteAddress() net.Addr {
//	return nil
//}

func (this *Channel) IsOpen() bool {
	return this.state == StateOpen
}

func (this *Channel) IsRegistered() bool {
	return true
}

func (this *Channel) IsActive() bool {
	return this.state == StateActive
}

func (this *Channel) IsWritable() bool {
	return true
}

func (this *Channel) DoRegister() (err error) {
	this.state = StateActive
	return
}

func (this *Channel) DoBind() (err error) {
	// Noop
	return
}

func (this *Channel) SetEventLoop(loop types.EventLoop) {
	//panic("implement me")
}

func (this *Channel) DoDisconnect() (err error) {
	// Noop
	this.DoClose()
	return
}

func (this *Channel) DoClose() (err error) {
	this.state = StateClosed
	return
}

func (this *Channel) DoDeregister() (err error) {
	// Noop
	return
}

func (this *Channel) DoBeginRead() (err error) {
	this.EventLoop().(*EventLoop).run()
	return
}

func (this *Channel) DoWrite(buffer types.OutboundBuffer) (err error) {
	for msg := buffer.Front(); msg != nil; msg = buffer.Front() {
		this.outboundMessages.PushBack(msg)
		buffer.RemoveFront()
	}
	return
}

func (this *Channel) WriteInbound(msgs ...interface{}) bool {
	pl := this.Pipeline()
	for _, msg := range msgs {
		pl.FireChannelRead(msg)
	}
	pl.FireChannelReadComplete()

	if this.inboundMessages.Len() == 0 {
		return false
	}
	return true
}

func (this *Channel) ReadInbound() interface{} {
	front := this.inboundMessages.Front()
	if front == nil {
		return nil
	}
	this.inboundMessages.Remove(front)
	return front.Value
}

func (this *Channel) WriteOutbound(msgs ...interface{}) bool {
	for _, msg := range msgs {
		this.Write(msg)
	}

	this.Flush()

	if this.outboundMessages.Len() == 0 {
		return false
	}
	return true
}

func (this *Channel) ReadOutbound() interface{} {
	front := this.outboundMessages.Front()
	if front == nil {
		return nil
	}
	this.outboundMessages.Remove(front)
	entry := front.Value.(*buffer.OutboundEntry)
	return entry.Msg
}

func (this *Channel) Finish() bool {
	this.Close()
	if this.outboundMessages.Len() != 0 || this.inboundMessages.Len() != 0 {
		return true
	}
	return false
}

func (this *Channel) onUnHandledChannelRead(msg interface{}) {
	this.inboundMessages.PushBack(msg)
}

type Unsafe struct {
	*channel.AbstractUnsafe
}

func (this *Unsafe) Connect(remoteAddress net.Addr, localAddress net.Addr, promise types.ChannelPromise) {

}
