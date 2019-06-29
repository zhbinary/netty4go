//Created by zhbinary on 2019-06-17.
//Email: zhbinary@gmail.com
package channel

import (
	"fmt"
	"github.com/zhbinary/heng/buffer"
	"github.com/zhbinary/heng/channel/embedded"
	"github.com/zhbinary/heng/handler"
	"github.com/zhbinary/heng/types"
	"testing"
)

func TestDefaultChannelPipeline_FireChannelActive(t *testing.T) {
	byteBuf := buffer.NewHeapBytebuf(1024)
	for i := 0; i < 9; i++ {
		byteBuf.WriteUint8(uint8(i))
	}

	ch := embedded.NewChannel(&In1{})
	if !ch.WriteInbound(byteBuf) {
		t.Error()
	}

	if !ch.Finish() {
		t.Error()
	}
}

type In1 struct {
	*handler.ChannelInboundHandlerAdapter
}

func (this *In1) ChannelActive(ctx types.ChannelHandlerContext) {
	fmt.Println("In1 ChannelActive")
}
