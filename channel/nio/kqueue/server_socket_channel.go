//Created by zhbinary on 2019-04-16.
//Email: zhbinary@gmail.com
package kqueue

import "syscall"

type ServerSocketChannel struct {
	*AbstractSocketChannel
}

func NewServerSocketChannel() *ServerSocketChannel {
	ret := &ServerSocketChannel{}
	return ret
}

func (this *ServerSocketChannel) doRead() {
	for {
		nfd, _, err := syscall.Accept(this.fd)
		if err != nil {
			if err == syscall.EAGAIN {
				// Wait for next turn
				return
			}
			// Close channel
			return
		}

		if err := syscall.SetNonblock(nfd, true); err != nil {
			// Close channel
			return
		}
		this.Pipeline().FireChannelRead(nfd)
	}
	this.Pipeline().FireChannelReadComplete()
}
