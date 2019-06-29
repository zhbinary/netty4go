//Created by zhbinary on 2019-01-09.
//Email: zhbinary@gmail.com
package types

type ChannelHandlerContext interface {
	ChannelInboundInvoker
	ChannelOutboundInvoker
	/**
	* Return the {@link Channel} which is bound to the {@link ChannelHandlerContext}.
	 */
	Channel() Channel

	/**
	 * The unique name of the {@link ChannelHandlerContext}.The name was used when then {@link ChannelHandler}
	 * was added to the {@link DefaultChannelPipeline}. This name can also be used to access the registered
	 * {@link ChannelHandler} from the {@link DefaultChannelPipeline}.
	 */
	Name() string

	/**
	 * The {@link ChannelHandler} that is bound this {@link ChannelHandlerContext}.
	 */
	Handler() ChannelHandler

	/**
	 * Return {@code true} if the {@link ChannelHandler} which belongs to this context was removed
	 * from the {@link DefaultChannelPipeline}. Note that this method is only meant to be called from with in the
	 * {@link EventLoop}.
	 */
	IsRemoved() bool

	//read(ctx ChannelHandlerContext) ChannelHandlerContext
	//
	//flush(ctx ChannelHandlerContext) ChannelHandlerContext

	/**
	 * Return the assigned {@link DefaultChannelPipeline}
	 */
	Pipeline() ChannelPipeline

	IsInbound() bool

	IsOutbound() bool
}
