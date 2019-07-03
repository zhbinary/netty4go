//Created by zhbinary on 2019-01-16.
//Email: zhbinary@gmail.com
package bio

import (
	"github.com/zhbinary/heng/channel"
	"github.com/zhbinary/heng/types"
	"net"
)

type ServerSocketChannel struct {
	channel.AbstractChannel
	tcpListener *net.TCPListener
}

func (this *ServerSocketChannel) SetEventLoop(loop types.EventLoop) {
	panic("implement me")
}

func (this *ServerSocketChannel) DoRegister() (err error) {
	panic("implement me")
}

func (this *ServerSocketChannel) DoBind() (err error) {
	this.tcpListener, err = net.ListenTCP("tcp", localAddress.(*net.TCPAddr))
}

func (this *ServerSocketChannel) DoDisconnect() (err error) {
	panic("implement me")
}

func (this *ServerSocketChannel) DoClose() (err error) {
	return this.tcpListener.Close()
}

func (this *ServerSocketChannel) DoDeregister() (err error) {
	panic("implement me")
}

func (this *ServerSocketChannel) DoBeginRead() (err error) {
	tcpConn, err := this.tcpListener.AcceptTCP()
	if err != nil {
		return err
	}
	child := &SocketChannel{tcpConn: tcpConn}
	this.Pipeline().FireChannelRead(child)
	return
}

func (this *ServerSocketChannel) DoWrite(buffer types.OutboundBuffer) (err error) {
	panic("implement me")
}

func (this *ServerSocketChannel) IsOpen() bool {
	panic("implement me")
}

func (this *ServerSocketChannel) IsActive() bool {
	panic("implement me")
}

func (this *ServerSocketChannel) LocalAddress() net.Addr {
	panic("implement me")
}

func (this *ServerSocketChannel) RemoteAddress() net.Addr {
	panic("implement me")
}

func (this *ServerSocketChannel) IsWritable() bool {
	panic("implement me")
}
