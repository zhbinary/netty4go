package thor

import (
	"sync"
	"github.com/zhbinary/thor/channel"
	"github.com/zhbinary/thor/handler/base"
)

type ServerBootstrap struct {
	address       string
	serverChannel *channel.ServerChannel
	pipeline      *base.ChannelHandlerPipeline
	grMu          sync.Mutex
	grWG          sync.WaitGroup
}

func NewBootstrap() *ServerBootstrap {
	return &ServerBootstrap{}
}

func (s *ServerBootstrap) Bind(address string) {
	s.address = address
	(*(s.serverChannel)).Bind(s, s.address)
}

func (s *ServerBootstrap) SetChannel(ch *channel.ServerChannel) {
	s.serverChannel = ch
}

func (s *ServerBootstrap) SetOption() *ServerBootstrap {
	return s
}

func (s *ServerBootstrap) SetHandler(handlers ... *base.ChannelHandler) *ServerBootstrap {
	for _, handler := range handlers {
		(*(s.serverChannel)).GetPipeline().AddLast(handler)
	}
	return s
}

func (s *ServerBootstrap) Go(f func()) {
	if f != nil {
		s.grWG.Add(1)
		go f()
	}
}
