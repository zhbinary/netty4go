// +build darwin dragonfly freebsd netbsd openbsd

//Created by zhbinary on 2018/10/17.
package kqueue

import (
	"syscall"
)

type WaitCb func(fd uint64, filter int16, data interface{})

type NetPoll struct {
	kq      int
	changes []syscall.Kevent_t
}

func NewPoll() (*NetPoll, error) {
	p := &NetPoll{}
	fd, err := syscall.Kqueue()
	if err != nil {
		return nil, err
	}

	p.kq = fd
	return p, nil
}

func (this *NetPoll) PollWait(cb WaitCb) error {
	events := make([]syscall.Kevent_t, 1024)
	for {
		n, err := syscall.Kevent(this.kq, nil, events, nil)
		if err != nil && err != syscall.EINTR {
			return err
		}

		for i := 0; i < n; i++ {
			ev := events[i]
			cb(ev.Ident, ev.Filter, ev.Data)
		}
	}
}

func (this *NetPoll) PollAddRead(fd uint64) {
	var changes []syscall.Kevent_t
	ev := syscall.Kevent_t{Ident: uint64(fd), Filter: syscall.EVFILT_READ, Flags: syscall.EV_ADD}
	changes = append(changes, ev)
	syscall.Kevent(this.kq, changes, nil, nil)
}

func (this *NetPoll) PollDelRead(fd uint64) {
	var changes []syscall.Kevent_t
	ev := syscall.Kevent_t{Ident: uint64(fd), Filter: syscall.EVFILT_READ, Flags: syscall.EV_DISABLE}
	changes = append(changes, ev)
	syscall.Kevent(this.kq, changes, nil, nil)
}

func (this *NetPoll) PollAddWrite(fd uint64) {
	var changes []syscall.Kevent_t
	ev := syscall.Kevent_t{Ident: uint64(uint64(fd)), Filter: syscall.EVFILT_WRITE, Flags: syscall.EV_ADD}
	changes = append(changes, ev)
	syscall.Kevent(this.kq, changes, nil, nil)
}

func (this *NetPoll) PollDelWrite(fd uint64) {
	var changes []syscall.Kevent_t
	ev := syscall.Kevent_t{Ident: uint64(fd), Filter: syscall.EVFILT_READ, Flags: syscall.EV_DISABLE}
	changes = append(changes, ev)
	syscall.Kevent(this.kq, changes, nil, nil)
}

func (this *NetPoll) PollAddReadWrite() {

}
