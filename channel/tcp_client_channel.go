package channel

import (
	"github.com/zhbinary/thor"
	"net"
	"github.com/zhbinary/thor/handler/base"
)

type TcpClientChannel struct {
	serverLauncher *thor.ServerBootstrap
	conn           *net.TCPConn
	pipline        *base.ChannelHandlerPipeline
	readBuf        []byte
	outChan        chan []byte
	close          chan interface{}
}

func NewChannel(s *thor.ServerBootstrap, c *net.TCPConn) *TcpClientChannel {
	return &TcpClientChannel{serverLauncher: s, conn: c, readBuf: make([]byte, 4096), close: make(chan interface{})}
}

func (c *TcpClientChannel) loop() {
	c.serverLauncher.Go(c.readLoop)
	c.serverLauncher.Go(c.writeLoop)
}

func (c *TcpClientChannel) readLoop() {
	for {
		n, err := c.conn.Read(c.readBuf)
		if err != nil {
			c.pipline.ReadChain(c.readBuf[:n])
		}
	}
}

func (c *TcpClientChannel) writeLoop() {
	for {
		select {
		case out := <-c.outChan:
			c.conn.Write(out)
		case <-c.close:
			return
		}
	}
}
