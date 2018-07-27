package channel

import (
	"net"
	"github.com/zhbinary/thor"
)

type TcpChannel struct {
	server  *thor.Server
	conn    *net.TCPConn
	readBuf []byte
	outChan chan []byte
	close   chan interface{}
}

func NewTcpChannel(s *thor.Server) *TcpChannel {
	return &TcpChannel{server: s, readBuf: make([]byte, 4096), close: make(chan interface{})}
}

func (c *TcpChannel) Listen(server *thor.Server, addr string) error {
	c.server = server
	localAddr, _ := net.ResolveTCPAddr("tcp", addr)
	listener, err := net.ListenTCP("tcp", localAddr)
	if err != nil {
		return err
	}
	conn, err := listener.AcceptTCP()
	if err != nil {
		return err
	}
	c.conn = conn
	c.loop()
	return nil
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
