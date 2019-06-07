//Created by zhbinary on 2019-04-18.
//Email: zhbinary@gmail.com
package types

type ByteBufAllocator interface {
	buffer(capacity int) ByteBuf
}
