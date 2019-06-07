//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package types

type ChannelPipeline interface {
	ChannelInboundInvoker
	ChannelOutboundInvoker
	AddLast(name string, handler ChannelHandler)
	Channel() Channel
}
