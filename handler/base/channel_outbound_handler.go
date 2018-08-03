package base

type ChannelOutboundHandler interface {
	ChannelHandler
	write(msg interface{})
}
