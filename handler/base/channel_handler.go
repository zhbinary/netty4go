package base

type ChannelHandler interface {
	handlerAdded()
	handlerRemoeved()
	exceptionCaught(err error)
}
