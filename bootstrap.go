//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package netty4go

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type Bootstrap struct {
	*AbstractBootstrap
}

func NewClient(opts *Options) *Bootstrap {

}

func (this *Bootstrap) Connect(remoteAddress net.Addr, promise types.Promise) types.Future {
	this.initAndRegister()
	return nil
}
