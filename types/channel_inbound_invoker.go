//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package types

type ChannelInboundInvoker interface {
	FireChannelRegistered()
	FireChannelUnregistered()
	FireChannelActive()
	FireChannelInactive()
	FireExceptionCaught(err error)
	FireUserEventTriggered(evt interface{})
	FireChannelRead(msg interface{})
	FireChannelReadComplete()
	FireChannelWritabilityChanged()
}
