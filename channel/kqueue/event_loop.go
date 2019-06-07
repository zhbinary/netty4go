//Created by zhbinary on 2018/10/17.
package kqueue

import (
	"github.com/zhbinary/heng/types"
	"syscall"
)

type EventLoop struct {
	ch       chan types.Task
	parent   types.EventLoopGroup
	netPoll  *NetPoll
	channels map[uint64]*AbstractSocketChannel
}

func NewEventLoop() (*EventLoop, error) {
	el := &EventLoop{}
	poll, err := NewPoll()
	if err != nil {
		return nil, err
	}
	el.ch = make(chan types.Task)
	el.netPoll = poll
	return el, nil
}

func (this *EventLoop) IsShutDown() bool {
	panic("implement me")
}

func (this *EventLoop) ShutdownGracefully(promise types.ChannelPromise) types.Future {
	panic("implement me")
}

func (this *EventLoop) AwaitTermination() bool {
	panic("implement me")
}

func (this *EventLoop) Next() types.EventLoop {
	panic("implement me")
}

func (this *EventLoop) Register(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	channel.Unsafe().Register(this, promise)
}

func (this *EventLoop) Submit(task types.Task, promise types.ChannelPromise) types.Future {
	this.ch <- task
}

func (this *EventLoop) Parent() types.EventLoopGroup {
	return this.parent
}

func (this *EventLoop) run() {
	for {
		select {
		case task := <-this.ch:
			task()
		default:
			this.netPoll.PollWait(func(fd uint64, filter int16, data interface{}) {
				if filter == syscall.EVFILT_READ {
					this.channels[fd].pollReadReady()
				} else if filter == syscall.EVFILT_WRITE {
					this.channels[fd].pollWriteReady()
				}
			})
		}
	}
}

func (this *EventLoop) add(channel *AbstractSocketChannel) error {
	this.channels[channel.Fd()] = channel
	this.netPoll.PollAddRead(channel.Fd())
	return nil
}

func (this *EventLoop) mod(channel *AbstractSocketChannel) error {
	this.netPoll.PollAddRead(channel.Fd())
	return nil
}

func (this *EventLoop) del(channel *AbstractSocketChannel) error {
	delete(this.channels, channel.Fd())
	return nil
}
