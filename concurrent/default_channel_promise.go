//Created by zhbinary on 2019-01-14.
//Email: zhbinary@gmail.com
package concurrent

import (
	"github.com/zhbinary/heng/types"
)

type DefaultChannelPromise struct {
	*AbstractPromise
	channel types.Channel
}

func NewDefaultChannelPromise(channel types.Channel) types.ChannelPromise {
	return &DefaultChannelPromise{AbstractPromise: NewAbstractPromise(), channel: channel}
}

func (this *DefaultChannelPromise) Channel() types.Channel {
	return this.channel
}
