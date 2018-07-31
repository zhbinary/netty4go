package channel

import "github.com/zhbinary/thor"

type ChannelContext interface {
	Listen(server *thor.ServerBootstrap, addr string) error
}
