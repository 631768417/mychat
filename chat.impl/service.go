package impl

import (
	. "chat.protocol"
	"chat.utils"
)

func newTid(name string, domain, resource *string) *Tid {
	tid := NewTid()
	tid.Domain = domain
	tid.Name = name
	tid.Resource = resource
	return tid
}

func OnlinePBean(tid *Tid) (pbean *PBean) {
	pbean = NewPBean()
	pbean.ThreadId = utils.TimeMills()
	pbean.FromTid = tid
	show, status := "online", "probe"
	pbean.Show, pbean.Status = &show, &status
	return
}

func OfflinePBean(tid *Tid) (pbean *PBean) {
	pbean = NewPBean()
	pbean.ThreadId = utils.TimeMills()
	pbean.FromTid = tid
	show, status := "offline", "unavailable"
	pbean.Show, pbean.Status = &show, &status
	return
}
