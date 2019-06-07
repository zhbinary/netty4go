//Created by zhbinary on 2019-01-10.
//Email: zhbinary@gmail.com
package heng

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type Bootstrap struct {
}

func NewClient(opts *Options) *Bootstrap {
}

func (this *Bootstrap) Connect(remoteAddress net.Addr, promise types.Promise) types.Future {
	return nil
}
