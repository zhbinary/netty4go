package thor

import (
	"net"
	"sync"
	"github.com/zhbinary/thor/handler"
	"github.com/zhbinary/thor/channel"
)

type Server struct {
	address  string
	listener *net.Listener
	chain    *handler.ChannelHandlerChain
	channel  *channel.Channel
	grMu     sync.Mutex
	grWG     sync.WaitGroup
}

func Bind(address string) *Server {
	s := &Server{address: address}
	return s
}

func (s *Server) Channel(ch *channel.Channel) {
	s.channel = ch
}

func (s *Server) Option() *Server {
	return s
}

func (s *Server) Handler(handlers ...handler.ChannelHandler) *Server {
	for _, handler := range handlers {
		s.chain.AddLast(handler)
	}
	return s
}

func (s *Server) Sync() (*Server, error) {
	(*(s.channel)).Listen(s, s.address)
	//l, err := net.Listen("tcp", s.address)
	//if err != nil {
	//	return s, err
	//}
	//s.listener = &l
	//for {
	//	conn, err := l.Accept()
	//	if err != nil {
	//		return s, err
	//	}
	//	//conn.(*net.TCPConn).
	//	NewChannel(conn, s).loop()
	//}
	return s, nil
}

func (s *Server) GetChain() *handler.ChannelHandlerChain {
	return s.chain
}

func (s *Server) Go(f func()) {
	if f != nil {
		s.grWG.Add(1)
		go f()
	}
}
