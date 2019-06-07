//Created by zhbinary on 2019-04-09.
//Email: zhbinary@gmail.com
package channel

import (
	"github.com/zhbinary/heng/types"
)

type AbstractEventLoop struct {
	parent types.EventLoopGroup
	ch     chan types.Task
	// implement by sub class
	start func()
}

func (this *AbstractEventLoop) IsShutDown() bool {
	panic("implement me")
}

func (this *AbstractEventLoop) ShutdownGracefully(promise types.ChannelPromise) types.Future {
	panic("implement me")
}

func (this *AbstractEventLoop) AwaitTermination() bool {
	panic("implement me")
}

func (this *AbstractEventLoop) Next() types.EventLoop {
	panic("implement me")
}

func (this *AbstractEventLoop) Register(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *AbstractEventLoop) Submit(task types.Task, promise types.ChannelPromise) types.Future {
	this.ch <- task
	promise.SetSuccess(nil)
	return promise
}

func (this *AbstractEventLoop) Parent() types.EventLoopGroup {
	return this.parent
}
