//Created by zhbinary on 2019-04-18.
//Email: zhbinary@gmail.com
package buffer

import (
	"github.com/zhbinary/heng/types"
	"sync"
)

type PooledByteBufAllocator struct {
	pool *sync.Pool
}

func NewPooledByteBufAllocator(maxCapacity int) *PooledByteBufAllocator {
	return &PooledByteBufAllocator{pool: &sync.Pool{New: func() interface{} {
		return NewHeadpBytebuf(maxCapacity)
	}}}
}

func (this *PooledByteBufAllocator) buffer(capacity int) types.ByteBuf {
	return this.pool.Get().(types.ByteBuf)
}
