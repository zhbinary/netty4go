package channel

import (
	"github.com/zhbinary/thor"
	"net"
)

type TcpChannelContext struct {
	server *thor.Server
	close  chan interface{}
}

func NewTcpChannelContext(s *thor.Server) *TcpChannelContext {
	return &TcpChannelContext{server: s, close: make(chan interface{})}
}

func (c *TcpChannelContext) Listen(server *thor.Server, addr string) error {
	c.server = server
	localAddr, _ := net.ResolveTCPAddr("tcp", addr)
	listener, err := net.ListenTCP("tcp", localAddr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		NewChannel(server, conn).loop()
	}
	return nil
}
