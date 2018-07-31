package channel

import (
	"github.com/zhbinary/thor"
	"net"
	"github.com/zhbinary/thor/handler/base"
)

type TcpServerChannel struct {
	serverLauncher *thor.ServerBootstrap
	pipline        *base.ChannelHandlerPipeline
}

func NewTcpServerChannel() *TcpServerChannel {
	return &TcpServerChannel{pipline: &base.ChannelHandlerPipeline{}}
}

func (this *TcpServerChannel) bind(launcher *thor.ServerBootstrap, addr string) error {
	this.serverLauncher = launcher
	localAddr, _ := net.ResolveTCPAddr("tcp", addr)
	listener, err := net.ListenTCP("tcp", localAddr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return err
		}
		NewChannel(launcher, conn).loop()
	}
	return nil
}

func (this *TcpServerChannel) GetPipeline() *base.ChannelHandlerPipeline {
	return this.pipline
}
