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

func NewDefaultChannelPromise() types.ChannelPromise {
	return &DefaultChannelPromise{AbstractPromise: NewAbstractPromise()}
}

func (this *DefaultChannelPromise) Channel() types.Channel {
	return this.channel
}
