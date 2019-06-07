//Created by zhbinary on 2018/10/17.
package types

type EventLoop interface {
	EventLoopGroup
	Parent() EventLoopGroup
}
