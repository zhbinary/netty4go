//Created by zhbinary on 2019-06-17.
//Email: zhbinary@gmail.com
package handler

import (
	"github.com/zhbinary/heng/types"
)

type ChannelInboundHandlerAdapter struct {
	*ChannelHandlerAdapter
}

func NewChannelInboundHandlerAdapter() *ChannelInboundHandlerAdapter {
	return &ChannelInboundHandlerAdapter{ChannelHandlerAdapter: &ChannelHandlerAdapter{}}
}

func (this *ChannelInboundHandlerAdapter) ChannelRead(ctx types.ChannelHandlerContext, msg interface{}) {
	ctx.FireChannelRead(msg)
}

func (this *ChannelInboundHandlerAdapter) ChannelReadComplete(ctx types.ChannelHandlerContext) {
	ctx.FireChannelReadComplete()
}

func (this *ChannelInboundHandlerAdapter) ChannelActive(ctx types.ChannelHandlerContext) {
	ctx.FireChannelActive()
}

func (this *ChannelInboundHandlerAdapter) ChannelInactive(ctx types.ChannelHandlerContext) {
	ctx.FireChannelInactive()
}

func (this *ChannelInboundHandlerAdapter) ChannelRegistered(ctx types.ChannelHandlerContext) {
	ctx.FireChannelRegistered()
}

func (this *ChannelInboundHandlerAdapter) ChannelUnregistered(ctx types.ChannelHandlerContext) {
	ctx.FireChannelUnregistered()
}

func (this *ChannelInboundHandlerAdapter) ChannelWritabilityChanged(ctx types.ChannelHandlerContext) {
	ctx.FireChannelWritabilityChanged()
}

func (this *ChannelInboundHandlerAdapter) UserEventTriggered(ctx types.ChannelHandlerContext, evt interface{}) {
	ctx.FireUserEventTriggered(evt)
}

func (this *ChannelInboundHandlerAdapter) ExceptionCaught(ctx types.ChannelHandlerContext, err error) {
	ctx.FireExceptionCaught(err)
}
