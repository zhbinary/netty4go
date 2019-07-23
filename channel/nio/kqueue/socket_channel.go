//Created by zhbinary on 2019-04-16.
//Email: zhbinary@gmail.com
package kqueue

import (
	"github.com/zhbinary/heng/types"
	"syscall"
)

type SocketChannel struct {
	AbstractSocketChannel
	readByteBuf  types.ByteBuf
	writeByteBuf types.ByteBuf
}

func NewSocketChannel() {

}

func (this *SocketChannel) pollReadReady(readBytes int) {
	for {
		if this.readByteBuf.WritableBytes() < readBytes {
			//this.readByteBuf.Increase()
		}
		writableBytes := this.readByteBuf.WritableBytes()
		n, err := syscall.Read(this.fd, this.readByteBuf.WritableArray())
		if err != nil {
			if err == syscall.EAGAIN {
				// Wait for next turn
				return
			}
			// Error occur, Close channel
			return
		}
		this.Pipeline().FireChannelRead(this.readByteBuf)

		// Continue
		if n < writableBytes || n == 0 {
			break
		}
	}
	this.Pipeline().FireChannelReadComplete()
}

func (this *SocketChannel) pollWriteReady() {
	//Pipeline().FireChannelWritabilityChanged()
	for {
		if this.writeByteBuf.ReadableBytes() == 0 {
			return
		}
		n, err := syscall.Write(Fd(), this.writeByteBuf.ReadableArray())
		if err != nil {
			if err == syscall.EAGAIN {
				// Wait for next turn
				return
			}
			// Error occur, Close channel
			return
		}
		this.writeByteBuf.SetReaderIndex(this.writeByteBuf.ReadableBytes() + n)
	}
}
