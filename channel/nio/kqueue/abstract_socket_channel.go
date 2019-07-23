//Created by zhbinary on 2019-04-16.
//Email: zhbinary@gmail.com
package kqueue

import "github.com/zhbinary/heng/channel"

type AbstractSocketChannel struct {
	*channel.AbstractChannel
	fd int
}

func (this *AbstractSocketChannel) Fd() int {
	return this.fd
}
