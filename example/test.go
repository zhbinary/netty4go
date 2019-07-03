package main

import (
	"fmt"
	"github.com/zhbinary/heng/buffer"
	"github.com/zhbinary/heng/channel/embedded"
	"github.com/zhbinary/heng/handler"
	"github.com/zhbinary/heng/types"
	"reflect"
)

func main() {
	byteBuf := buffer.NewHeapBytebuf(1024)
	for i := 0; i < 9; i++ {
		byteBuf.WriteUint8(uint8(i))
	}

	input := byteBuf.Duplicate()

	ch := embedded.NewChannel(NewIn1(), NewOut1())
	if !ch.WriteInbound(input) {
		fmt.Println("err")
	}

	if !ch.Finish() {
		fmt.Println("err")
	}

	read := ch.ReadInbound().(types.ByteBuf)
	if !reflect.DeepEqual(byteBuf.ReadableArray(), read.ReadableArray()) {
		fmt.Println("err buf")
	}

	if ch.ReadInbound() != nil {
		fmt.Println("err")
	}

	byteBufOut := buffer.NewHeapBytebuf(1024)
	for i := 0; i < 9; i++ {
		byteBufOut.WriteUint8(uint8(i))
	}

	output := byteBufOut.Duplicate()
	if !ch.WriteOutbound(output) {
		fmt.Println("err")
	}

	if !ch.Finish() {
		fmt.Println("err")
	}

	readOut := ch.ReadOutbound().(types.ByteBuf)
	if !reflect.DeepEqual(byteBufOut.ReadableArray(), readOut.ReadableArray()) {
		fmt.Println("err buf, isn't equal")
	}

	if ch.ReadOutbound() != nil {
		fmt.Println("err")
	}
}

type In1 struct {
	*handler.ChannelInboundHandlerAdapter
}

func NewIn1() *In1 {
	return &In1{ChannelInboundHandlerAdapter: handler.NewChannelInboundHandlerAdapter()}
}

func (this *In1) ChannelActive(ctx types.ChannelHandlerContext) {
	fmt.Println("In1 ChannelActive")
	ctx.FireChannelActive()
}

func (this *In1) ChannelRead(ctx types.ChannelHandlerContext, msg interface{}) {
	bf := msg.(types.ByteBuf)
	fmt.Println("In1 ChannelRead", bf.String())
	ctx.FireChannelRead(msg)
}

type Out1 struct {
	*handler.ChannelOutboundHandlerAdapter
}

func NewOut1() *Out1 {
	return &Out1{ChannelOutboundHandlerAdapter: handler.NewChannelOutboundHandlerAdapter()}
}

func (this *Out1) Write(ctx types.ChannelHandlerContext, msg interface{}, promise types.ChannelPromise) {
	fmt.Println("Out1 Write")
	ctx.Write0(msg, promise)
}
