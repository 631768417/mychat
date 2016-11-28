package tfClient

import (
	"fmt"
	"testing"

	//	"github.com/apache/thrift/lib/go/thrift"
	//	"chat.logger"
	. "chat.protocol"
)

func TestRemote(t *testing.T) {
	tid := NewTid()

	tid.Name = "734604"
	pwd := "e10adc3949ba59abbe56e057f20f883e"
	HttpClient(func(client *ITimClient) {
		r, er := client.TimRemoteUserAuth(tid, pwd)
		if er == nil && r != nil {
			fmt.Println(r)
			if r.ExtraMap != nil {
				if password, ok := r.ExtraMap["password"]; ok {
					if pwd == password {
						fmt.Print("ok")
					}
				}
				if extraAuth, ok := r.ExtraMap["extraAuth"]; ok {
					if pwd == extraAuth {
						fmt.Print("ok2")
					}
				}
			}
		}
	})
}
