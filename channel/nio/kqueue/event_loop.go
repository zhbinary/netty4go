//Created by zhbinary on 2018/10/17.
package kqueue

import (
	"github.com/zhbinary/heng/buffer"
	"github.com/zhbinary/heng/concurrent"
	"github.com/zhbinary/heng/types"
	"syscall"
)

type EventLoop struct {
	parent          types.EventLoopGroup
	tasks           chan types.Runnable
	kqFd            int
	changes, events []syscall.Kevent_t
	channels        map[uint64]types.Channel
}

func NewEventLoop() (el *EventLoop, err error) {
	fd, err := syscall.Kqueue()
	if err != nil {
		return
	}
	el = &EventLoop{tasks: make(chan types.Runnable), kqFd: fd}
	return
}

func (this *EventLoop) IsShutDown() bool {
	panic("implement me")
}

func (this *EventLoop) ShutdownGracefully() types.ChannelFutrue {
	panic("implement me")
}

func (this *EventLoop) Next() types.EventLoop {
	return this.parent.Next()
}

func (this *EventLoop) Register(channel types.Channel) types.ChannelFutrue {
	promise := concurrent.NewDefaultChannelPromise(channel)
	this.Register0(promise)
	return promise
}

func (this *EventLoop) Register0(promise types.ChannelPromise) types.ChannelFutrue {
	promise.Channel().Unsafe().Register(this, promise)
	return promise
}

func (this *EventLoop) Register1(channel types.Channel, promise types.ChannelPromise) types.ChannelFutrue {
	channel.Unsafe().Register(this, promise)
	return promise
}

func (this *EventLoop) Submit(task types.Runnable) types.Future {
	panic("implement me")
}

func (this *EventLoop) Execute(task types.Runnable) {
	this.tasks <- task
}

func (this *EventLoop) Parent() types.EventLoopGroup {
	return this.parent
}

func (this *EventLoop) run() {
	var timeout *syscall.Timespec
	for {
		n, err := syscall.Kevent(this.kqFd, this.changes, this.events, timeout)
		if err != nil {

		}

		if n > 0 {
			this.processKeys()
		}

	}

	this.cleanChanges()
}

func (this *EventLoop) addRead(fd uint64) {
	event := syscall.Kevent_t{Ident: fd, Filter: syscall.EVFILT_READ, Flags: syscall.EV_ADD}
	this.changes = append(this.changes, event)
}

func (this *EventLoop) processKeys() {
	for _, event := range this.events {
		switch event.Filter {
		case syscall.EVFILT_READ:
			this.readReady(&event)
		case syscall.EVFILT_WRITE:
		}
	}
}

func (this *EventLoop) readReady(event *syscall.Kevent_t) {
	// Pass data to pipeline
	channel := this.channels[event.Ident]

	buf := buffer.NewHeapBytebuf(int(event.Data))
	// Read through unsafe
}

func (this *EventLoop) writeReady() {
	// flush out buffer to socket
}

func (this *EventLoop) cleanChanges() {
	this.changes = nil
}
