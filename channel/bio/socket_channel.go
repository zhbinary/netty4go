//Created by zhbinary on 2019-01-15.
//Email: zhbinary@gmail.com
package bio

import (
	"github.com/zhbinary/heng/channel"
	"github.com/zhbinary/heng/types"
	"net"
)

type SocketChannel struct {
	channel.AbstractSocketChannel
	tcpConn *net.TCPConn
}

func (this *SocketChannel) IsActive() bool {
	panic(types.ErrUnsupportedOperation)
}

func (this *SocketChannel) IsOpen() bool {
	panic("implement me")
}

func (this *SocketChannel) IsRegistered() bool {
	panic("implement me")
}

func (this *SocketChannel) localAddress() net.Addr {
	return this.tcpConn.LocalAddr()
}

func (this *SocketChannel) remoteAddress() net.Addr {
	this.tcpConn.RemoteAddr()
}

func (this *SocketChannel) doRegister() error {
	panic("implement me")
}

func (this *SocketChannel) doDeRegister() error {
	panic("implement me")
}

func (this *SocketChannel) doBind(localAddress net.Addr) error {
	panic("implement me")
}

func (this *SocketChannel) doConnect(localAddress net.Addr, remoteAddress net.Addr) (err error) {
	this.tcpConn, err = net.DialTCP("tcp", localAddress.(*net.TCPAddr), remoteAddress.(*net.TCPAddr))
	return
}

func (this *SocketChannel) doDisconnect() error {
	return this.doClose()
}

func (this *SocketChannel) doClose() error {
	return this.tcpConn.Close()
}

func (this *SocketChannel) doBeginRead() error {
	buf := this.Config().GetByteBuf()
	_, err := buf.ReadOnce(this.tcpConn)
	if err != nil {
		return err
	}
	this.Pipeline().FireChannelRead(buf)
}

func (this *SocketChannel) newUnsafe() types.Unsafe {
	return &DefaultUnsafe{}
}
