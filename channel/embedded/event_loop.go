//Created by zhbinary on 2019-06-18.
//Email: zhbinary@gmail.com
package embedded

import (
	"github.com/zhbinary/heng/concurrent"
	"github.com/zhbinary/heng/types"
)

type EventLoop struct {
	ch      chan types.Runnable
	started bool
}

func NewEventLoop() *EventLoop {
	return &EventLoop{ch: make(chan types.Runnable, 128)}
}

func (this *EventLoop) IsShutDown() bool {
	return false
}

func (this *EventLoop) ShutdownGracefully() types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Next() types.EventLoop {
	return this
}

func (this *EventLoop) Register(channel types.Channel) types.ChannelFutrue {
	promise := concurrent.NewDefaultChannelPromise(channel)
	return this.Register0(promise)
}

func (this *EventLoop) Register0(promise types.ChannelPromise) types.ChannelFutrue {
	promise.Channel().Unsafe().Register(this, promise)
	return promise
}

func (this *EventLoop) Register1(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Submit(task types.Runnable) types.Future {
	panic("implement me")
}

func (this *EventLoop) Execute(task types.Runnable) {
	this.startLoop()
	this.ch <- task
}

func (this *EventLoop) Parent() types.EventLoopGroup {
	panic("implement me")
}

func (this *EventLoop) startLoop() {
	if !this.started {
		this.started = true
		this.run()
	}
}

func (this *EventLoop) run() {
	for {
		select {
		case task := <-this.ch:
			task()
		}
	}
}
