package channel

import "github.com/zhbinary/thor"

type Channel interface {
	Listen(server *thor.Server, addr string) error
}
