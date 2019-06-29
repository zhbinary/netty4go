//Created by zhbinary on 2019-06-19.
//Email: zhbinary@gmail.com
package channel

import (
	"errors"
	"github.com/zhbinary/heng/buffer"
	"github.com/zhbinary/heng/types"
	"net"
)

type AbstractChannel struct {
	id         string
	registered bool
	eventLoop  types.EventLoop
	parent     types.Channel
	pipeline   *DefaultChannelPipeline
	unsafe     types.Unsafe

	DoRegister   func() error
	DoBind       func() error
	DoDisconnect func() error
	DoClose      func() error
	DoDeregister func() error
	DoBeginRead  func() error
	DoWrite      func(buffer *buffer.OutboundBuffer) error
}

func NewAbstractChannel(id string, loop types.EventLoop, parent types.Channel, unsafe types.Unsafe) *AbstractChannel {
	ch := &AbstractChannel{id: id, eventLoop: loop, parent: parent, unsafe: unsafe}
	ch.pipeline = NewDefaultChannelPipeline(ch)
	return ch
}

func (this *AbstractChannel) Bind(localAddress net.Addr) types.ChannelFutrue {
	return this.pipeline.Bind(localAddress)
}

func (this *AbstractChannel) Bind0(localAddress net.Addr, promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.Bind0(localAddress, promise)
}

func (this *AbstractChannel) Connect(localAddress net.Addr, remoteAddress net.Addr) types.ChannelFutrue {
	return this.pipeline.Connect(localAddress, remoteAddress)
}

func (this *AbstractChannel) Connect0(localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.Connect0(localAddress, remoteAddress, promise)
}

func (this *AbstractChannel) Disconnect() types.ChannelFutrue {
	return this.pipeline.Disconnect()
}

func (this *AbstractChannel) Disconnect0(promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.Disconnect0(promise)
}

func (this *AbstractChannel) Close() types.ChannelFutrue {
	return this.pipeline.Close()
}

func (this *AbstractChannel) Close0(promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.Close0(promise)
}

func (this *AbstractChannel) Deregister() types.ChannelFutrue {
	return this.pipeline.Deregister()
}

func (this *AbstractChannel) Deregister0(promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.Deregister0(promise)
}

func (this *AbstractChannel) Read() {
	this.pipeline.Read()
}

func (this *AbstractChannel) Write(msg interface{}) types.ChannelFutrue {
	return this.pipeline.Write(msg)
}

func (this *AbstractChannel) Write0(msg interface{}, promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.Write0(msg, promise)
}

func (this *AbstractChannel) WriteAndFlush(msg interface{}) types.ChannelFutrue {
	return this.pipeline.WriteAndFlush(msg)
}

func (this *AbstractChannel) WriteAndFlush0(msg interface{}, promise types.ChannelPromise) types.ChannelFutrue {
	return this.pipeline.WriteAndFlush0(msg, promise)
}

func (this *AbstractChannel) Flush() {
	this.pipeline.Flush()
}

func (this *AbstractChannel) Id() string {
	return this.id
}

func (this *AbstractChannel) EventLoop() types.EventLoop {
	return this.eventLoop
}

func (this *AbstractChannel) Parent() types.Channel {
	return this.parent
}

func (this *AbstractChannel) IsOpen() bool {
	panic("implement me")
}

func (this *AbstractChannel) IsRegistered() bool {
	return this.registered
}

func (this *AbstractChannel) IsActive() bool {
	panic("implement me")
}

func (this *AbstractChannel) LocalAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractChannel) RemoteAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractChannel) IsWritable() bool {
	panic("implement me")
}

func (this *AbstractChannel) Unsafe() types.Unsafe {
	return this.unsafe
}

func (this *AbstractChannel) Config() types.ChannelConfig {
	panic("implement me")
}

func (this *AbstractChannel) Pipeline() types.ChannelPipeline {
	return this.pipeline
}

type AbstractUnsafe struct {
	Channel types.Channel
	Buffer  *buffer.OutboundBuffer
}

func NewAbstractUnsafe(channel types.Channel) *AbstractUnsafe {
	return &AbstractUnsafe{Channel: channel, Buffer: &buffer.OutboundBuffer{}}
}

func (this *AbstractUnsafe) LocalAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractUnsafe) RemoteAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractUnsafe) Register(eventLoop types.EventLoop, promise types.ChannelPromise) {
	if eventLoop == nil {
		promise.SetFailure(errors.New("eventLoop is nil"))
	}

	ch := this.Channel.(*AbstractChannel)
	if !ch.IsActive() {
		promise.SetFailure(errors.New("registered to an event loop already"))
	}

	this.Channel.eventLoop = eventLoop
	if err := this.Channel.DoRegister(); err != nil {
		promise.SetFailure(err)
	} else {
		promise.SetSuccess()
	}

	this.Channel.Pipeline().FireChannelRegistered()
}

func (this *AbstractUnsafe) Bind(localAddress net.Addr, promise types.ChannelPromise) {
	if err := this.Channel.DoBind; err != nil {
		promise.SetFailure(err())
	} else {
		promise.SetSuccess()
	}
}

func (this *AbstractUnsafe) Connect(remoteAddress net.Addr, localAddress net.Addr, promise types.ChannelPromise) {
	panic("Implement me")
}

func (this *AbstractUnsafe) Disconnect(promise types.ChannelPromise) {
	if err := this.Channel.DoDisconnect; err != nil {
		promise.SetFailure(err())
	} else {
		promise.SetSuccess()
	}
}

func (this *AbstractUnsafe) Close(promise types.ChannelPromise) {
	if err := this.Channel.DoClose; err != nil {
		promise.SetFailure(err())
	} else {
		promise.SetSuccess()
	}
}

func (this *AbstractUnsafe) CloseForcibly() {
	this.Channel.DoClose()
}

func (this *AbstractUnsafe) Deregister(promise types.ChannelPromise) {
	if err := this.Channel.DoDeregister(); err != nil {
		promise.SetFailure(err)
	} else {
		promise.SetSuccess()
		if this.Channel.IsActive() {
			this.Channel.Pipeline().FireChannelInactive()
		}

		if this.Channel.IsRegistered() {
			this.Channel.Pipeline().FireChannelUnregistered()
		}
	}
}

func (this *AbstractUnsafe) BeginRead() {
	if !this.Channel.IsActive() {
		return
	}

	if err := this.Channel.DoBeginRead(); err != nil {
		this.Channel.Pipeline().FireExceptionCaught(err)
	}
}

func (this *AbstractUnsafe) Write(msg interface{}, promise types.ChannelPromise) {
	this.Buffer.AddMessage(msg, promise)
}

func (this *AbstractUnsafe) Flush() {
	if err := this.Channel.DoWrite(this.Buffer); err != nil {
	}
}

//type AbstractUnsafe struct {
//	Channel *AbstractChannel
//	Buffer  *buffer.OutboundBuffer
//}
//
//func (this *AbstractUnsafe) LocalAddress() net.Addr {
//	panic("implement me")
//}
//
//func (this *AbstractUnsafe) RemoteAddress() net.Addr {
//	panic("implement me")
//}
//
//func (this *AbstractUnsafe) Register(eventLoop types.EventLoop, promise types.ChannelPromise) {
//	if eventLoop == nil {
//		promise.SetFailure(errors.New("eventLoop is nil"))
//	}
//
//	if !this.Channel.IsActive() {
//		promise.SetFailure(errors.New("registered to an event loop already"))
//	}
//
//	this.Channel.eventLoop = eventLoop
//	if err := this.Channel.DoRegister(); err != nil {
//		promise.SetFailure(err)
//	} else {
//		promise.SetSuccess()
//	}
//
//	this.Channel.Pipeline().FireChannelRegistered()
//}
//
//func (this *AbstractUnsafe) Bind(localAddress net.Addr, promise types.ChannelPromise) {
//	if err := this.Channel.DoBind; err != nil {
//		promise.SetFailure(err())
//	} else {
//		promise.SetSuccess()
//	}
//}
//
//func (this *AbstractUnsafe) Connect(remoteAddress net.Addr, localAddress net.Addr, promise types.ChannelPromise) {
//	panic("Implement me")
//}
//
//func (this *AbstractUnsafe) Disconnect(promise types.ChannelPromise) {
//	if err := this.Channel.DoDisconnect; err != nil {
//		promise.SetFailure(err())
//	} else {
//		promise.SetSuccess()
//	}
//}
//
//func (this *AbstractUnsafe) Close(promise types.ChannelPromise) {
//	if err := this.Channel.DoClose; err != nil {
//		promise.SetFailure(err())
//	} else {
//		promise.SetSuccess()
//	}
//}
//
//func (this *AbstractUnsafe) CloseForcibly() {
//	this.Channel.DoClose()
//}
//
//func (this *AbstractUnsafe) Deregister(promise types.ChannelPromise) {
//	if err := this.Channel.DoDeregister(); err != nil {
//		promise.SetFailure(err)
//	} else {
//		promise.SetSuccess()
//		if this.Channel.IsActive() {
//			this.Channel.Pipeline().FireChannelInactive()
//		}
//
//		if this.Channel.IsRegistered() {
//			this.Channel.Pipeline().FireChannelUnregistered()
//		}
//	}
//}
//
//func (this *AbstractUnsafe) BeginRead() {
//	if !this.Channel.IsActive() {
//		return
//	}
//
//	if err := this.Channel.DoBeginRead(); err != nil {
//		this.Channel.Pipeline().FireExceptionCaught(err)
//	}
//}
//
//func (this *AbstractUnsafe) Write(msg interface{}, promise types.ChannelPromise) {
//	this.Buffer.AddMessage(msg, promise)
//}
//
//func (this *AbstractUnsafe) Flush() {
//	if err := this.Channel.DoWrite(this.Buffer); err != nil {
//	}
//}
