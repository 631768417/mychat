package tfClient

import (
	"errors"
	"fmt"
	"runtime/debug"

	"chat.logger"
	. "chat.protocol"
	"github.com/apache/thrift/lib/go/thrift"
)

func HttpClient(f func(*IClient) error, urlstr string) (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprint(er))
			logger.Error(er)
			logger.Error(string(debug.Stack()))
		}
	}()
	if urlstr != "" {
		logger.Debug("httpClient url:", urlstr)
		transport, err := thrift.NewTHttpPostClient(urlstr)
		defer transport.Close()
		if err == nil {
			factory := thrift.NewTCompactProtocolFactory()
			transport.Open()
			itimClient := NewIClientFactory(transport, factory)
			err = f(itimClient)
		}
	}
	return
}

func HttpClient2(f func(*IClient) error, user_auth_url string) (err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprint(er))
			logger.Error(er)
			logger.Error(string(debug.Stack()))
		}
	}()
	if user_auth_url != "" {
		logger.Debug("httpClient url:", user_auth_url)
		transport, err := thrift.NewTHttpPostClient(user_auth_url)
		defer transport.Close()
		if err == nil {
			factory := thrift.NewTCompactProtocolFactory()
			transport.Open()
			itimClient := NewIClientFactory(transport, factory)
			err = f(itimClient)
		}
	} else {
		err = errors.New("httpclient url is null")
	}
	return
}

func TcpClient(f func(*IClient), urlstr string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(err)
		}
	}()
	if urlstr != "" {
		logger.Debug("tcpClient addr:", urlstr)
		transport, err := thrift.NewTSocket(urlstr)
		defer transport.Close()
		if err == nil {
			protocolFactory := thrift.NewTCompactProtocolFactory()
			itimClient := NewIClientFactory(transport, protocolFactory)
			f(itimClient)
		}
	}
}
