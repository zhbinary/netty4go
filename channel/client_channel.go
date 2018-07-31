package channel

type ClientChannel interface {
	Channel
	connect(addr string)
	disconnect()
}
