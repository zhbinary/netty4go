package types

type OutboundBuffer interface {
	AddMessage(msg interface{}, promise ChannelPromise)
	Front() interface{}
	RemoveFront()
}
