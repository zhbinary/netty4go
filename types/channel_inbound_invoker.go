//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package types

type ChannelInboundInvoker interface {
	FireChannelRegistered() ChannelInboundInvoker
	FireChannelUnregistered() ChannelInboundInvoker
	FireChannelActive() ChannelInboundInvoker
	FireChannelInactive() ChannelInboundInvoker
	FireExceptionCaught(err error) ChannelInboundInvoker
	FireUserEventTriggered(evt interface{}) ChannelInboundInvoker
	FireChannelRead(msg interface{}) ChannelInboundInvoker
	FireChannelReadComplete() ChannelInboundInvoker
	FireChannelWritabilityChanged() ChannelInboundInvoker
}
