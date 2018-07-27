package handler

type ChannelOutboundHandler interface {
	channelWrite(msg interface{})
}
