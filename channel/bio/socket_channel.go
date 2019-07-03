//Created by zhbinary on 2019-01-15.
//Email: zhbinary@gmail.com
package bio

import (
	"github.com/zhbinary/heng/channel"
	"github.com/zhbinary/heng/types"
	"net"
)

type SocketChannel struct {
	channel.AbstractChannel
	tcpConn *net.TCPConn
}

func (this *SocketChannel) SetEventLoop(loop types.EventLoop) {
	panic("implement me")
}

func (this *SocketChannel) DoRegister() (err error) {
	panic("implement me")
}

func (this *SocketChannel) DoBind() (err error) {
	panic("implement me")
}

func (this *SocketChannel) DoDisconnect() (err error) {
	panic("implement me")
}

func (this *SocketChannel) DoClose() (err error) {
	panic("implement me")
}

func (this *SocketChannel) DoDeregister() (err error) {
	panic("implement me")
}

func (this *SocketChannel) DoBeginRead() (err error) {
	panic("implement me")
}

func (this *SocketChannel) DoWrite(buffer types.OutboundBuffer) (err error) {
	panic("implement me")
}
