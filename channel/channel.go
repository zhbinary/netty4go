package channel

type Channel interface {
	Write()
	WriteAndFlush()
}

