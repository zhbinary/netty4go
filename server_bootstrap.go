//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package heng

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type ServerBootstrap struct {
	AbstractBootstrap
	childHandler types.ChannelHandler
}

func NewServerBootstrap(opts *Options) *ServerBootstrap {
	return &ServerBootstrap{opts: opts}
}

func (this *ServerBootstrap) Bind(localAddress net.Addr, promise types.Promise) types.Future {
	return nil
}

func (this *ServerBootstrap) Group(parentGroup types.EventLoopGroup, childGroup types.EventLoopGroup) {

}

func (this *ServerBootstrap) ChildHandler(childHandler types.ChannelHandler) {
	this.childHandler = childHandler
}
