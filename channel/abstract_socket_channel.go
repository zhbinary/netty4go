//Created by zhbinary on 2019-01-14.
//Email: zhbinary@gmail.com
package channel

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type AbstractSocketChannel struct {
	id            string
	parent        types.Channel
	eventLoop     types.EventLoop
	pipeline      types.ChannelPipeline
	unsafe        types.Unsafe
	localAddress  net.Addr
	remoteAddress net.Addr
	channelConfig types.ChannelConfig
}

func (this *AbstractSocketChannel) Bind(localAddress net.Addr, promise types.ChannelPromise) types.Future {
	return this.pipeline.Bind0(localAddress, promise)
}

func (this *AbstractSocketChannel) Connect(localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) types.Future {
	return this.pipeline.Connect0(localAddress, remoteAddress, promise)
}

func (this *AbstractSocketChannel) Disconnect(promise types.ChannelPromise) types.Future {
	return this.pipeline.Disconnect0(promise)
}

func (this *AbstractSocketChannel) Close(promise types.ChannelPromise) types.Future {
	return this.pipeline.Close0(promise)
}

func (this *AbstractSocketChannel) Deregister(promise types.ChannelPromise) types.Future {
	return this.pipeline.Deregister0(promise)
}

func (this *AbstractSocketChannel) Write(msg interface{}, promise types.ChannelPromise) types.Future {
	return this.pipeline.Write0(msg, promise)
}

func (this *AbstractSocketChannel) WriteAndFlush(msg interface{}, promise types.ChannelPromise) types.Future {
	return this.pipeline.WriteAndFlush0(msg, promise)
}

func (this *AbstractSocketChannel) Flush() {
	this.pipeline.Flush()
}

func (this *AbstractSocketChannel) Id() string {
	return this.id
}

func (this *AbstractSocketChannel) EventLoop() types.EventLoop {
	return this.eventLoop
}

func (this *AbstractSocketChannel) Parent() types.Channel {
	return this.parent
}

func (this *AbstractSocketChannel) LocalAddress() net.Addr {
	if this.localAddress == nil {
		this.localAddress = this.unsafe.LocalAddress()
	}
	return this.localAddress
}

func (this *AbstractSocketChannel) RemoteAddress() net.Addr {
	if this.remoteAddress == nil {
		this.remoteAddress = this.unsafe.RemoteAddress()
	}
	return this.remoteAddress
}

func (this *AbstractSocketChannel) IsWritable() bool {
	panic("implement me")
}

func (this *AbstractSocketChannel) Unsafe() types.Unsafe {
	return this.unsafe
}

func (this *AbstractSocketChannel) Pipeline() types.ChannelPipeline {
	return this.pipeline
}

func (this *AbstractSocketChannel) Config() types.ChannelConfig {
	return this.channelConfig
}
