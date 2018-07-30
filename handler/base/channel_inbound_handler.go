package base

type ChannelInboundHandler interface {
	ChannelHandler
	channelRead(data interface{}) (interface{}, error)
	channelReadComplete()
	exceptionCaught(err error)
}
