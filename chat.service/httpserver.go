package service

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime/debug"
	"strings"
	"time"

	. "chat.common"
	. "chat.impl"
	"chat.logger"
	daoService "chat.mysql/service"
	"chat.protocol"
	"github.com/apache/thrift/lib/go/thrift"
)

func Httpserver() {
	if CF.GetHttpPort() <= 0 {
		return
	}
	http.HandleFunc("/my", my)
	http.HandleFunc("/info", info)
	http.HandleFunc("/uinfo", userInfo)
	http.HandleFunc("/hi", hbaseclient)
	s := &http.Server{
		Addr:           fmt.Sprint(":", CF.GetHttpPort()),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("httpserver start listen:", CF.GetHttpPort())
	s.ListenAndServe()
}

func my(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("err:", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	logger.Debug("RemoteAddr:", r.RemoteAddr)
	logger.Debug("X-Forwarded-For:", r.Header.Get("X-Forwarded-For"))
	logger.Debug("ContentLength:", r.ContentLength)
	X_Forwarded_For := r.Header.Get("X-Forwarded-For")
	ss := strings.Split(r.RemoteAddr, ":")
	ipaddr := ss[0]
	if X_Forwarded_For != "" && X_Forwarded_For != "127.0.0.1" {
		ipaddr = X_Forwarded_For
	}
	if r.ContentLength >= 100*1024*1024 {
		return
	}
	if !daoService.AllowHttpIp(ipaddr) {
		logger.Info("ipaddr is not allow", "[", ipaddr, "]")
		return
	}
	if "POST" == r.Method {
		protocolFactory := thrift.NewTCompactProtocolFactory()
		//		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		transport := thrift.NewStreamTransport(r.Body, w)
		inProtocol := protocolFactory.GetProtocol(transport)
		outProtocol := protocolFactory.GetProtocol(transport)
		handler := &Impl{Ip: ipaddr}
		processor := protocol.NewIProcessor(handler)
		processor.Process(inProtocol, outProtocol)
	}
}
