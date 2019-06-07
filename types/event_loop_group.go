//Created by zhbinary on 2018/10/17.
package types

type EventLoopGroup interface {
	IsShutDown() bool
	ShutdownGracefully(promise ChannelPromise) Future
	AwaitTermination() bool
	Next() EventLoop
	Register(channel Channel, promise ChannelPromise) ChannelFutrue
	Execute(task Task, promise ChannelPromise) Future
}
