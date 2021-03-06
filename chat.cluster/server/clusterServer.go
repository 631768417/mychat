package clusterServer

import (
	"fmt"

	"chat.client"
	. "chat.common"
	"chat.logger"
	. "chat.protocol"
	"github.com/apache/thrift/lib/go/thrift"
)

type Controlloer struct {
	addr string
}

func ServerStart() {
	s := new(Controlloer)
	s.SetAddr(ClusterConf.RequestAddr)
	s.Server()
}

func (t *Controlloer) SetAddr(addr string) {
	t.addr = addr
}

func (t *Controlloer) ListenAddr() string {
	return t.addr
}

func (t *Controlloer) Server() {

	//构造 thrift 协议
	//缓存传输
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	//压缩二进制
	protocolFactory := thrift.NewTCompactProtocolFactory()
	serverTransport, err := thrift.NewTServerSocket(t.ListenAddr())

	if err != nil {
		logger.Error("server:", err.Error())
		panic(err.Error())
	}

	//RPC服务接口
	handler := new(client.MyImpl)
	processor := NewIProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	fmt.Println("cluster server listen:", t.ListenAddr())
	server.Serve()
}
