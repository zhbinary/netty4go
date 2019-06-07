//Created by zhbinary on 2019-01-18.
//Email: zhbinary@gmail.com
package channel

import (
	"github.com/zhbinary/heng/types"
)

type DefaultChannelConfig struct {
}

func (this *DefaultChannelConfig) GetByteBuf() types.ByteBuf {
	panic("implement me")
}

func (this *DefaultChannelConfig) GetOption() *ChannelOption {
	panic("implement me")
}

func (this *DefaultChannelConfig) SetOption(option *ChannelOption) {
	panic("implement me")
}

func (this *DefaultChannelConfig) GetOptions() []*ChannelOption {
	panic("implement me")
}

func (this *DefaultChannelConfig) SetOptions(options []*ChannelOption) {
	panic("implement me")
}
