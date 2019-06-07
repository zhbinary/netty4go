package types

import (
	"net"
)

type Channel interface {
	ChannelOutboundInvoker
	/**
     * Returns the globally unique identifier of this {@link Channel}.
     */
	Id() string

	/**
	 * Return the {@link EventLoop} this {@link Channel} was registered to.
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
     * Returns {@code true} if the {@link Channel} is open and may get active later
     */
	IsOpen() bool

	/**
	 * Returns {@code true} if the {@link Channel} is registered with an {@link EventLoop}.
	 */
	IsRegistered() bool

	/**
	 * Return {@code true} if the {@link Channel} is active and so connected.
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
     * Returns an <em>internal-use-only</em> object that provides unsafe operations.
     */
	Unsafe() Unsafe

	Config() ChannelConfig

	/**
     * Return the assigned {@link DefaultChannelPipeline}.
     */
	Pipeline() ChannelPipeline

	localAddress() net.Addr
	remoteAddress() net.Addr
	newUnsafe() Unsafe
	doRegister() error
	doDeRegister() error
	doBind(localAddress net.Addr) error
	doConnect(localAddress net.Addr, remoteAddress net.Addr) error
	doDisconnect() error
	doClose() error
	doBeginRead() error
}

type Unsafe interface {
	/**
	 * Return the {@link SocketAddress} to which is bound local or
	 * {@code null} if none.
	 */
	LocalAddress() net.Addr

	/**
	 * Return the {@link SocketAddress} to which is bound remote or
	 * {@code null} if none is bound yet.
	 */
	RemoteAddress() net.Addr

	/**
	 * Register the {@link Channel} of the {@link ChannelPromise} and notify
	 * the {@link ChannelFuture} once the registration was complete.
	 */
	Register(eventLoop EventLoop, promise Promise)

	/**
	 * Bind the {@link SocketAddress} to the {@link Channel} of the {@link ChannelPromise} and notify
	 * it once its done.
	 */
	Bind(localAddress net.Addr, promise Promise)

	/**
	 * Connect the {@link Channel} of the given {@link ChannelFuture} with the given remote {@link SocketAddress}.
	 * If a specific local {@link SocketAddress} should be used it need to be given as argument. Otherwise just
	 * pass {@code null} to it.
	 *
	 * The {@link ChannelPromise} will get notified once the connect operation was complete.
	 */
	Connect(remoteAddress net.Addr, localAddress net.Addr, promise Promise)

	/**
	 * Disconnect the {@link Channel} of the {@link ChannelFuture} and notify the {@link ChannelPromise} once the
	 * operation was complete.
	 */
	Disconnect(promise Promise)

	/**
	 * Close the {@link Channel} of the {@link ChannelPromise} and notify the {@link ChannelPromise} once the
	 * operation was complete.
	 */
	Close(promise Promise)

	/**
	 * Closes the {@link Channel} immediately without firing any events.  Probably only useful
	 * when registration attempt failed.
	 */
	CloseForcibly()

	/**
	 * Deregister the {@link Channel} of the {@link ChannelPromise} from {@link EventLoop} and notify the
	 * {@link ChannelPromise} once the operation was complete.
	 */
	Deregister(promise Promise)

	/**
	 * Schedules a read operation that fills the inbound buffer of the first {@link ChannelInboundHandler} in the
	 * {@link ChannelPipeline}.  If there's already a pending read operation, this method does nothing.
	 */
	BeginRead()

	/**
	 * Schedules a write operation.
	 */
	Write(msg interface{}, promise Promise)

	/**
	 * Flush out all write operations scheduled via {@link #write(Object, ChannelPromise)}.
	 */
	Flush()
}
