//Created by zhbinary on 2018/10/17.
package types

type EventLoopGroup interface {
	IsShutDown() bool
	ShutdownGracefully() ChannelFutrue
	Next() EventLoop
	Register(channel Channel) ChannelFutrue
	Register0(promise ChannelPromise) ChannelFutrue
	Register1(channel Channel, promise ChannelPromise) ChannelFutrue
	Submit(task Runnable) Future
	Execute(task Runnable)
}
