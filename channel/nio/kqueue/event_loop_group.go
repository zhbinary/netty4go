//Created by zhbinary on 2018/10/17.
package kqueue

import (
	"github.com/zhbinary/heng/types"
)

type EventLoopGroup struct {
}

func (this *EventLoopGroup) IsShutDown() bool {
	panic("implement me")
}

func (this *EventLoopGroup) ShutdownGracefully() types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoopGroup) Next() types.EventLoop {
	panic("implement me")
}

func (this *EventLoopGroup) Register(channel types.Channel) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoopGroup) Register0(promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoopGroup) Register1(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoopGroup) Submit(task types.Runnable) types.Future {
	panic("implement me")
}

func (this *EventLoopGroup) Execute(task types.Runnable) {
	panic("implement me")
}
