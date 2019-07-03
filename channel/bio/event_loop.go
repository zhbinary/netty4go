//Created by zhbinary on 2018/10/17.
package bio

import (
	"container/list"
	"github.com/zhbinary/heng/types"
)

type EventLoop struct {
	tasks   *list.List
	channel *SocketChannel
}

func NewEventLoop() *EventLoop {
	return &EventLoop{tasks: list.New()}
}

func (this *EventLoop) IsShutDown() bool {
	panic("implement me")
}

func (this *EventLoop) ShutdownGracefully() types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Next() types.EventLoop {
	panic("implement me")
}

func (this *EventLoop) Register(channel types.Channel) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Register0(promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Register1(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Submit(task types.Runnable) types.Future {
	panic("implement me")
}

func (this *EventLoop) Execute(task types.Runnable) {
	if task == nil {
		return
	}

	this.tasks.PushBack(task)
}

func (this *EventLoop) Parent() types.EventLoopGroup {
	panic("implement me")
}

func (this *EventLoop) startRoutine() {
}
