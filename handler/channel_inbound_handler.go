package handler

type ChannelInboundHandler interface {
	ChannelHandler
	channelRead(data []byte) (interface{}, error)
}
