//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package types

import (
	"net"
)

type ChannelOutboundInvoker interface {
	Bind(localAddress net.Addr, promise Promise) Future
	Connect(localAddress net.Addr, remoteAddress net.Addr, promise Promise) Future
	Disconnect(promise Promise) Future
	Close(promise Promise) Future
	Deregister(promise Promise) Future
	Write(msg interface{}, promise Promise) Future
	WriteAndFlush(msg interface{}, promise Promise) Future
	Flush()
}
