//Created by zhbinary on 2019-06-17.
//Email: zhbinary@gmail.com
package handler

import (
	"github.com/zhbinary/heng/types"
	"net"
)

type ChannelOutboundHandlerAdapter struct {
	*ChannelHandlerAdapter
}

func (this *ChannelOutboundHandlerAdapter) Bind(ctx types.ChannelHandlerContext, localAddress net.Addr, promise types.ChannelPromise) {
	ctx.Bind0(localAddress, promise)
}

func (this *ChannelOutboundHandlerAdapter) Connect(ctx types.ChannelHandlerContext, localAddress net.Addr, remoteAddress net.Addr, promise types.ChannelPromise) {
	ctx.Connect0(localAddress, remoteAddress, promise)
}

func (this *ChannelOutboundHandlerAdapter) Disconnect(ctx types.ChannelHandlerContext, promise types.ChannelPromise) {
	ctx.Disconnect0(promise)
}

func (this *ChannelOutboundHandlerAdapter) Close(ctx types.ChannelHandlerContext, promise types.ChannelPromise) {
	ctx.Close0(promise)
}

func (this *ChannelOutboundHandlerAdapter) Deregister(ctx types.ChannelHandlerContext, promise types.ChannelPromise) {
	ctx.Deregister0(promise)
}

func (this *ChannelOutboundHandlerAdapter) Write(ctx types.ChannelHandlerContext, msg interface{}, promise types.ChannelPromise) {
	ctx.Write0(msg, promise)
}

func (this *ChannelOutboundHandlerAdapter) Read(ctx types.ChannelHandlerContext) {
	ctx.Read()
}

func (this *ChannelOutboundHandlerAdapter) Flush(ctx types.ChannelHandlerContext) {
	ctx.Flush()
}
