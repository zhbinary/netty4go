//Created by zhbinary on 2019-01-18.
//Email: zhbinary@gmail.com
package types

import (
	"github.com/zhbinary/heng/channel"
)

type ChannelConfig interface {
	GetOption() *channel.ChannelOption
	SetOption(option *channel.ChannelOption)
	GetOptions() []*channel.ChannelOption
	SetOptions(options []*channel.ChannelOption)
	GetByteBuf() ByteBuf
	//getAllocator()
	//SetAllocator()
}
