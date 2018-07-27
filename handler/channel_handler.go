package handler

type ChannelHandler interface {
	//handlerAdded()
	//handlerRemoeved()
	exceptionCaught(err error)
}
