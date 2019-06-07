//Created by zhbinary on 2019-04-16.
//Email: zhbinary@gmail.com
package kqueue

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type AbstractSocketChannel struct {
	fd             int
	pollReadReady  func()
	pollWriteReady func()
	pollRdHupReady func()
}

func (this *AbstractSocketChannel) Bind(localAddress net.Addr, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) Connect(localAddress net.Addr, remoteAddress net.Addr, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) Disconnect(promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) Close(promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) Deregister(promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) Write(msg interface{}, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) WriteAndFlush(msg interface{}, promise types.Promise) types.Future {
	panic("implement me")
}

func (this *AbstractSocketChannel) Flush() {
	panic("implement me")
}

func (this *AbstractSocketChannel) Id() string {
	panic("implement me")
}

func (this *AbstractSocketChannel) EventLoop() types.EventLoop {
	panic("implement me")
}

func (this *AbstractSocketChannel) Parent() types.Channel {
	panic("implement me")
}

func (this *AbstractSocketChannel) IsOpen() bool {
	panic("implement me")
}

func (this *AbstractSocketChannel) IsRegistered() bool {
	panic("implement me")
}

func (this *AbstractSocketChannel) IsActive() bool {
	panic("implement me")
}

func (this *AbstractSocketChannel) LocalAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractSocketChannel) RemoteAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractSocketChannel) IsWritable() bool {
	panic("implement me")
}

func (this *AbstractSocketChannel) Unsafe() types.Unsafe {
	panic("implement me")
}

func (this *AbstractSocketChannel) Config() types.ChannelConfig {
	panic("implement me")
}

func (this *AbstractSocketChannel) Pipeline() types.ChannelPipeline {
	panic("implement me")
}

func (this *AbstractSocketChannel) localAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractSocketChannel) remoteAddress() net.Addr {
	panic("implement me")
}

func (this *AbstractSocketChannel) newUnsafe() types.Unsafe {
	panic("implement me")
}

func (this *AbstractSocketChannel) doRegister() error {
	panic("implement me")
}

func (this *AbstractSocketChannel) doDeRegister() error {
	panic("implement me")
}

func (this *AbstractSocketChannel) doBind(localAddress net.Addr) error {
	panic("implement me")
}

func (this *AbstractSocketChannel) doConnect(localAddress net.Addr, remoteAddress net.Addr) error {
	panic("implement me")
}

func (this *AbstractSocketChannel) doDisconnect() error {
	panic("implement me")
}

func (this *AbstractSocketChannel) doClose() error {
	panic("implement me")
}

func (this *AbstractSocketChannel) doBeginRead() error {
	panic("implement me")
}

func (this *AbstractSocketChannel) Fd() int {
	return this.fd
}
