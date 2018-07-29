package base

type ChannelOutboundHandler interface {
	ChannelHandler
	channelWrite(msg interface{}) (interface{}, error)
}
