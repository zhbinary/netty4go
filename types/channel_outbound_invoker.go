//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package types

import (
	"net"
)

type ChannelOutboundInvoker interface {
	Bind(localAddress net.Addr) (future ChannelFutrue)
	Bind0(localAddress net.Addr, promise ChannelPromise) (future ChannelFutrue)
	Connect(localAddress net.Addr, remoteAddress net.Addr) (future ChannelFutrue)
	Connect0(localAddress net.Addr, remoteAddress net.Addr, promise ChannelPromise) (future ChannelFutrue)
	Disconnect() (future ChannelFutrue)
	Disconnect0(promise ChannelPromise) (future ChannelFutrue)
	Close() (future ChannelFutrue)
	Close0(promise ChannelPromise) (future ChannelFutrue)
	Deregister() (future ChannelFutrue)
	Deregister0(promise ChannelPromise) (future ChannelFutrue)
	Read()
	Write(msg interface{}) (future ChannelFutrue)
	Write0(msg interface{}, promise ChannelPromise) (future ChannelFutrue)
	WriteAndFlush(msg interface{}) (future ChannelFutrue)
	WriteAndFlush0(msg interface{}, promise ChannelPromise) (future ChannelFutrue)
	Flush()
}
