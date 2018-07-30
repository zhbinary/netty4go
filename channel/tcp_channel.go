package channel

import (
	"github.com/zhbinary/thor"
	"net"
	"github.com/zhbinary/thor/handler/base"
)

type TcpChannel struct {
	server  *thor.Server
	conn    *net.TCPConn
	pipline *base.ChannelHandlerPipeline
	readBuf []byte
	outChan chan []byte
	close   chan interface{}
}

func NewChannel(s *thor.Server, c *net.TCPConn) *TcpChannel {
	return &TcpChannel{server: s, conn: c, readBuf: make([]byte, 4096), close: make(chan interface{})}
}

func (c *TcpChannel) loop() {
	c.server.Go(c.readLoop)
	c.server.Go(c.writeLoop)
}

func (c *TcpChannel) readLoop() {
	for {
		n, err := c.conn.Read(c.readBuf)
		if err != nil {
			chain := c.server.GetChain
			chain().ReadChain(c.readBuf[:n])
		}
	}
}

func (c *TcpChannel) writeLoop() {
	for {
		select {
		case out := <-c.outChan:
			c.conn.Write(out)
		case <-c.close:
			return
		}
	}
}
