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
	channelBundle types.ChannelBundle
	buffer        *buffer.OutboundBuffer
}

func NewAbstractUnsafe(channelBundle types.ChannelBundle) *AbstractUnsafe {
	return &AbstractUnsafe{channelBundle: channelBundle, buffer: &buffer.OutboundBuffer{}}
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

	if !this.channelBundle.IsActive() {
		promise.SetFailure(errors.New("registered to an event loop already"))
	}

	this.channelBundle.SetEventLoop(eventLoop)
	//this.channelBundle.eventLoop = eventLoop
	if err := this.channelBundle.DoRegister(); err != nil {
		promise.SetFailure(err)
	} else {
		promise.SetSuccess()
	}

	this.channelBundle.Pipeline().FireChannelRegistered()
}

func (this *AbstractUnsafe) Bind(localAddress net.Addr, promise types.ChannelPromise) {
	if err := this.channelBundle.DoBind; err != nil {
		promise.SetFailure(err())
	} else {
		promise.SetSuccess()
	}
}

func (this *AbstractUnsafe) Connect(remoteAddress net.Addr, localAddress net.Addr, promise types.ChannelPromise) {
	panic("Implement me")
}

func (this *AbstractUnsafe) Disconnect(promise types.ChannelPromise) {
	if err := this.channelBundle.DoDisconnect; err != nil {
		promise.SetFailure(err())
	} else {
		promise.SetSuccess()
	}
}

func (this *AbstractUnsafe) Close(promise types.ChannelPromise) {
	if err := this.channelBundle.DoClose; err != nil {
		promise.SetFailure(err())
	} else {
		promise.SetSuccess()
	}
}

func (this *AbstractUnsafe) CloseForcibly() {
	this.channelBundle.DoClose()
}

func (this *AbstractUnsafe) Deregister(promise types.ChannelPromise) {
	if err := this.channelBundle.DoDeregister(); err != nil {
		promise.SetFailure(err)
	} else {
		promise.SetSuccess()
		if this.channelBundle.IsActive() {
			this.channelBundle.Pipeline().FireChannelInactive()
		}

		if this.channelBundle.IsRegistered() {
			this.channelBundle.Pipeline().FireChannelUnregistered()
		}
	}
}

func (this *AbstractUnsafe) BeginRead() {
	if !this.channelBundle.IsActive() {
		return
	}

	if err := this.channelBundle.DoBeginRead(); err != nil {
		this.channelBundle.Pipeline().FireExceptionCaught(err)
	}
}

func (this *AbstractUnsafe) Write(msg interface{}, promise types.ChannelPromise) {
	this.buffer.AddMessage(msg, promise)
}

func (this *AbstractUnsafe) Flush() {
	if err := this.channelBundle.DoWrite(this.buffer); err != nil {
	}
}
