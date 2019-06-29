//Created by zhbinary on 2019-04-09.
//Email: zhbinary@gmail.com
package channel

import (
	"github.com/zhbinary/heng/types"
)

type AbstractEventLoop struct {
	parent types.EventLoopGroup
	ch     chan types.Runnable
	// implement by sub class
	start func()
}

func (this *AbstractEventLoop) IsShutDown() bool {
	panic("implement me")
}

func (this *AbstractEventLoop) ShutdownGracefully() types.ChannelFutrue {
	panic("implement me")
}

func (this *AbstractEventLoop) Next() types.EventLoop {
	panic("implement me")
}

func (this *AbstractEventLoop) Register(channel types.Channel) types.ChannelFutrue {
	panic("implement me")
}

func (this *AbstractEventLoop) Register0(promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *AbstractEventLoop) Register1(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *AbstractEventLoop) Submit(task types.Runnable) types.Future {
	panic("implement me")
}

func (this *AbstractEventLoop) Execute(task types.Runnable) {
	this.ch <- task
}

func (this *AbstractEventLoop) Parent() types.EventLoopGroup {
	return this.parent
}
