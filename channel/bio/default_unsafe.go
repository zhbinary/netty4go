//Created by zhbinary on 2019-01-16.
//Email: zhbinary@gmail.com
package bio

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type DefaultUnsafe struct {
	channel types.Channel
}

func (this *DefaultUnsafe) LocalAddress() net.Addr {
}

func (this *DefaultUnsafe) RemoteAddress() net.Addr {
}

func (this *DefaultUnsafe) Register(eventLoop types.EventLoop, promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) Bind(localAddress net.Addr, promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) Connect(remoteAddress net.Addr, localAddress net.Addr, promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) Disconnect(promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) Close(promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) CloseForcibly() {
	panic("implement me")
}

func (this *DefaultUnsafe) Deregister(promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) BeginRead() {
	panic("implement me")
}

func (this *DefaultUnsafe) Write(msg interface{}, promise types.Promise) {
	panic("implement me")
}

func (this *DefaultUnsafe) Flush() {
	panic("implement me")
}
