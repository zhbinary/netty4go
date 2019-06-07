//Created by zhbinary on 2019-04-18.
//Email: zhbinary@gmail.com
package buffer

import "github.com/zhbinary/heng/types"

type UnpooledBytebufAllocator struct {
}

func (this *UnpooledBytebufAllocator) buffer(capacity int) types.ByteBuf {
	return &HeapByteBuf{}
}
