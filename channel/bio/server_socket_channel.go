//Created by zhbinary on 2019-01-16.
//Email: zhbinary@gmail.com
package bio

import (
	"github.com/zhbinary/heng/channel"
	"github.com/zhbinary/heng/types"
	"net"
)

type ServerSocketChannel struct {
	channel.AbstractSocketChannel
	tcpListener *net.TCPListener
}

func (this *ServerSocketChannel) IsOpen() bool {
	panic("implement me")
}

func (this *ServerSocketChannel) IsRegistered() bool {
	panic("implement me")
}

func (this *ServerSocketChannel) IsActive() bool {
	panic("implement me")
}

func (this *ServerSocketChannel) localAddress() net.Addr {
	return nil
}

func (this *ServerSocketChannel) remoteAddress() net.Addr {
	return nil
}

func (this *ServerSocketChannel) newUnsafe() types.Unsafe {
	return &DefaultUnsafe{}
}

func (this *ServerSocketChannel) doRegister() error {
	panic("implement me")
}

func (this *ServerSocketChannel) doDeRegister() error {
	panic("implement me")
}

func (this *ServerSocketChannel) doBind(localAddress net.Addr) (err error) {
	this.tcpListener, err = net.ListenTCP("tcp", localAddress.(*net.TCPAddr))
	return
}

func (this *ServerSocketChannel) doConnect(localAddress net.Addr, remoteAddress net.Addr) error {
	panic("implement me")
}

func (this *ServerSocketChannel) doDisconnect() error {
	panic("implement me")
}

func (this *ServerSocketChannel) doClose() error {
	return this.tcpListener.Close()
}

func (this *ServerSocketChannel) doBeginRead() error {
	tcpConn, err := this.tcpListener.AcceptTCP()
	if err != nil {
		return err
	}
	child := &SocketChannel{tcpConn: tcpConn}
	this.Pipeline().FireChannelRead(child)
}
