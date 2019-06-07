//Created by zhbinary on 2019-04-20.
//Email: zhbinary@gmail.com
package heng

import (
	"github.com/zhbinary/heng/concurrent"
	"github.com/zhbinary/heng/types"
	"net"
	"reflect"
)

type AbstractBootstrap struct {
	group       types.EventLoopGroup
	handler     types.ChannelHandler
	opts        *Options
	attrs       map[string]interface{}
	channelType reflect.Type
	initChannel func(channel types.Channel)
}

func (this *AbstractBootstrap) Bind(remote net.Addr) {
	regFuture := this.initAndRegister()
	channel := regFuture.Channel()
	if regFuture.IsDone() {
		promise := concurrent.NewDefaultChannelPromise()
		channel.Bind(, promise)
	} else {
		regFuture.AddListener(func(future types.Future) {
			promise := concurrent.NewDefaultChannelPromise()
			channel.Bind(, promise)
		})
	}
}

func (this *AbstractBootstrap) Bind0(local net.Addr, remote net.Addr) {

}

func (this *AbstractBootstrap) Channel(t reflect.Type) {
	this.channelType = t
}

func (this *AbstractBootstrap) newChannel() types.Channel {
	return reflect.New(this.channelType).Interface().(types.Channel)
}

func (this *AbstractBootstrap) Group(group types.EventLoopGroup) {
	this.group = group
}

func (this *AbstractBootstrap) Handler(handler types.ChannelHandler) {
	this.handler = handler
}

func (this *AbstractBootstrap) initAndRegister() types.ChannelFutrue {
	channel := this.newChannel()
	this.initChannel(channel)

	promise := concurrent.NewDefaultChannelPromise()
	return this.group.Register(channel, promise)
}
