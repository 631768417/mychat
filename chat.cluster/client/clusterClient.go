package clusterClient

import (
	"os"
	"runtime/debug"
	"sync"

	"chat.logger"
	. "chat.protocol"
	"github.com/apache/thrift/lib/go/thrift"
)

type ClusterClient struct {
	myclient *IClient
	lock     *sync.RWMutex
	Weight   int
}

func (this *ClusterClient) Close() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Close,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.Weight = 0
	if this.myclient != nil {
		this.myclient.Transport.Flush()
		this.myclient.Transport.Close()
	}
}

func (this *ClusterClient) SendMBean(mbean *MBean, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.lock.Lock()
	defer this.lock.Unlock()
	r, err = this.myclient.ResponseMessage(mbean, auth)
	return
}

func (this *ClusterClient) SendMBeanList(mbeanList *MBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendMBeanList,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.lock.Lock()
	defer this.lock.Unlock()
	r, err = this.myclient.ResponseMessageList(mbeanList, auth)
	return
}

func (this *ClusterClient) SendPBean(pbean *PBean, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendPBean error:", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.lock.Lock()
	defer this.lock.Unlock()
	r, err = this.myclient.ResponsePresence(pbean, auth)
	return
}

func (this *ClusterClient) SendPBeanList(pbeanList *PBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendPBeanList error:", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.lock.Lock()
	defer this.lock.Unlock()
	r, err = this.myclient.ResponsePresenceList(pbeanList, auth)
	return
}

func NewClusterClient(addr string) (clusterClient *ClusterClient, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("ClusterClient,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	transportFactory := thrift.NewTBufferedTransportFactory(1024)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport, er := thrift.NewTSocket(addr)
	if er != nil {
		logger.Error(os.Stderr, "error resolving address:", err)
		err = er
		return
	}
	useTransport := transportFactory.GetTransport(transport)
	myclient := NewIClientFactory(useTransport, protocolFactory)
	if er = transport.Open(); er != nil {
		logger.Error(os.Stderr, "Error opening socket to ", addr, " ", er)
		err = er
		return
	}
	clusterClient = new(ClusterClient)
	clusterClient.myclient = myclient
	clusterClient.lock = new(sync.RWMutex)
	clusterClient.Weight = 1
	return
}
