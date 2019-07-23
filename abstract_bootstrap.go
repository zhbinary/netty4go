//Created by zhbinary on 2019-04-20.
//Email: zhbinary@gmail.com
package netty4go

import (
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
	addr := &net.TCPAddr{IP: net.IP{}, Port: 8280}
	if regFuture.IsDone() {
		channel.Bind(addr)
	} else {
		regFuture.AddListener(&types.FutureListenerAdapter{OperationCompleteCb: func(future types.Future) {
			channel.Bind(addr)
		}})
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
	return this.group.Register(channel)
}
