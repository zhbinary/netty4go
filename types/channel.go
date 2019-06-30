package types

import (
	"net"
)

type Channel interface {
	ChannelOutboundInvoker

	/**
	 * Returns an <em>internal-use-only</em> object that provides unsafe operations.
	 */
	Unsafe() Unsafe

	/**
	 * Returns the globally unique identifier of this {@link channel}.
	 */
	Id() string

	/**
	 * Return the {@link EventLoop} this {@link channel} was registered to.
	 */
	EventLoop() EventLoop

	/**
	 * Returns the parent of this channel.
	 *
	 * @return the parent channel.
	 *         {@code null} if this channel does not have a parent channel.
	 */
	Parent() Channel

	/**
	 * Returns {@code true} if the {@link channel} is open and may get active later
	 */
	IsOpen() bool

	/**
	 * Returns {@code true} if the {@link channel} is registered with an {@link EventLoop}.
	 */
	IsRegistered() bool

	/**
	 * Return {@code true} if the {@link channel} is active and so connected.
	 */
	IsActive() bool

	/**
	 * Returns the local address where this channel is bound to.  The returned
	 * {@link SocketAddress} is supposed to be down-cast into more concrete
	 * type such as {@link InetSocketAddress} to retrieve the detailed
	 * information.
	 *
	 * @return the local address of this channel.
	 *         {@code null} if this channel is not bound.
	 */
	LocalAddress() net.Addr

	/**
	 * Returns the remote address where this channel is connected to.  The
	 * returned {@link SocketAddress} is supposed to be down-cast into more
	 * concrete type such as {@link InetSocketAddress} to retrieve the detailed
	 * information.
	 *
	 * @return the remote address of this channel.
	 *         {@code null} if this channel is not connected.
	 *         If this channel is not connected but it can receive messages
	 *         from arbitrary remote addresses (e.g. {@link DatagramChannel},
	 *         use {@link DatagramPacket#recipient()} to determine
	 *         the origination of the received message as this method will
	 *         return {@code null}.
	 */
	RemoteAddress() net.Addr

	/**
	 * Returns {@code true} if and only if the I/O thread will perform the
	 * requested write operation immediately.  Any write requests made when
	 * this method returns {@code false} are queued until the I/O thread is
	 * ready to process the queued write requests.
	 */
	IsWritable() bool

	/**
	 * Return the assigned {@link DefaultChannelPipeline}.
	 */
	Pipeline() ChannelPipeline

	Config() ChannelConfig
}

type ChannelBundle interface {
	Channel
	ChannelInner
}

type ChannelInner interface {
	SetEventLoop(loop EventLoop)
	DoRegister() (err error)
	DoBind() (err error)
	DoDisconnect() (err error)
	DoClose() (err error)
	DoDeregister() (err error)
	DoBeginRead() (err error)
	DoWrite(buffer OutboundBuffer) (err error)
}

type Unsafe interface {
	/**
	 * Register the {@link channel} of the {@link ChannelPromise} and notify
	 * the {@link ChannelFuture} once the registration was complete.
	 */
	Register(eventLoop EventLoop, promise ChannelPromise)

	/**
	 * Bind the {@link SocketAddress} to the {@link channel} of the {@link ChannelPromise} and notify
	 * it once its done.
	 */
	Bind(localAddress net.Addr, promise ChannelPromise)

	/**
	 * Connect the {@link channel} of the given {@link ChannelFuture} with the given remote {@link SocketAddress}.
	 * If a specific local {@link SocketAddress} should be used it need to be given as argument. Otherwise just
	 * pass {@code null} to it.
	 *
	 * The {@link ChannelPromise} will get notified once the connect operation was complete.
	 */
	Connect(remoteAddress net.Addr, localAddress net.Addr, promise ChannelPromise)

	/**
	 * Disconnect the {@link channel} of the {@link ChannelFuture} and notify the {@link ChannelPromise} once the
	 * operation was complete.
	 */
	Disconnect(promise ChannelPromise)

	/**
	 * Close the {@link channel} of the {@link ChannelPromise} and notify the {@link ChannelPromise} once the
	 * operation was complete.
	 */
	Close(promise ChannelPromise)

	/**
	 * Closes the {@link channel} immediately without firing any events.  Probably only useful
	 * when registration attempt failed.
	 */
	CloseForcibly()

	/**
	 * Deregister the {@link channel} of the {@link ChannelPromise} from {@link EventLoop} and notify the
	 * {@link ChannelPromise} once the operation was complete.
	 */
	Deregister(promise ChannelPromise)

	/**
	 * Schedules a read operation that fills the inbound buffer of the first {@link ChannelInboundHandler} in the
	 * {@link ChannelPipeline}.  If there's already a pending read operation, this method does nothing.
	 */
	BeginRead()

	/**
	 * Schedules a write operation.
	 */
	Write(msg interface{}, promise ChannelPromise)

	/**
	 * Flush out all write operations scheduled via {@link #write(Object, ChannelPromise)}.
	 */
	Flush()
}
