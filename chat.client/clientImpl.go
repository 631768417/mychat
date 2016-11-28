package client

import (
	"errors"
	"fmt"
	//	"time"
	//	"runtime/debug"

	. "chat.common"
	"chat.logger"
	daoService "chat.mysql/service"
	. "chat.protocol"
	"chat.route"
	"chat.utils"
)

type MyImpl struct {
	Ip   string
	Port int
	Pub  string //发布id
}

// Parameters:
//  - Param
func (this *MyImpl) Stream(param *Param) (err error) {
	panic("error")
	return
}
func (this *MyImpl) Starttls() (err error) {
	panic("error")
	return
}

// Parameters:
//  - Tid
//  - Pwd
func (this *MyImpl) Login(tid *Tid, pwd string) (err error) {
	logger.Debug("Login:", tid, pwd)
	panic("error")
	return
}

// Parameters:
//  - Ab
func (this *MyImpl) Ack(ab *AckBean) (err error) {
	logger.Debug("Ack=========>", ab)
	panic("error")
	return
}

// Parameters:
//  - Pbean
func (this *MyImpl) Presence(pbean *PBean) (err error) {
	logger.Debug(pbean)
	panic("error")
	return
}

// Parameters:
//  - Mbean
func (this *MyImpl) Message(mbean *MBean) (err error) {
	logger.Debug(mbean)
	panic("error")
	return
}

// Parameters:
//  - ThreadId
func (this *MyImpl) Ping(threadId string) (err error) {
	panic("error")
	return
}

// Parameters:
//  - E
func (this *MyImpl) Error(e *Error) (err error) {
	panic("error")
	return
}
func (this *MyImpl) Logout() (err error) {
	panic("error")
	return
}

// Parameters:
//  - Tid
//  - Pwd
func (this *MyImpl) Regist(tid *Tid, pwd string) (err error) {
	panic("error")
	return
}

// Parameters:
//  - Tid
//  - Pwd
func (this *MyImpl) RemoteUserAuth(tid *Tid, pwd string, auth *MyAuth) (r *RemoteUserBean, err error) {
	panic("error")
	return
}

// Parameters:
//  - Tid
func (this *MyImpl) RemoteUserGet(tid *Tid, auth *MyAuth) (r *RemoteUserBean, err error) {
	panic("error")
	return
}

// Parameters:
//  - Tid
//  - Ub
func (this *MyImpl) RemoteUserEdit(tid *Tid, ub *UserBean, auth *MyAuth) (r *RemoteUserBean, err error) {
	panic("error")
	return
}

// Parameters:
//  - Pbean
func (this *MyImpl) ResponsePresence(pbean *PBean, auth *MyAuth) (r *ResponseBean, err error) {
	logger.Debug("ResponsePresence", pbean, auth)
	if !checkAuth(auth) {
		err = errors.New(fmt.Sprint("cluster auth fail:", auth))
		return
	}
	go _MyResponsePresence(pbean, auth)
	return
}

func _MyResponsePresence(pbean *PBean, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if er := recover(); er != nil {
			logger.Error("error:", er)
		}
	}()
	if pbean.GetToTid() == nil {
		route.RoutePBean(pbean)
	} else {
		route.RouteSinglePBean(pbean)
	}
	return
}

// Parameters:
//  - Mbean
func (this *MyImpl) ResponseMessage(mbean *MBean, auth *MyAuth) (r *ResponseBean, err error) {
	logger.Debug("ResponseMessage", mbean, auth)
	if !checkAuth(auth) {
		err = errors.New(fmt.Sprint("cluster auth fail:", auth))
		return
	}
	go _MyResponseMessage(mbean, auth)
	return
}

func _MyResponseMessage(mbean *MBean, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if er := recover(); er != nil {
			logger.Error("error:", er)
		}
	}()
	//	r = NewResponseBean()
	fromDomain := mbean.GetFromTid().GetDomain()
	toDomain := mbean.GetToTid().GetDomain()
	if fromDomain == toDomain {
		if !daoService.CheckDomain(fromDomain) {
			logger.Error("domain check fail:", fromDomain)
			return
		}
	} else {
		logger.Error("fromDomain != toDomain", fromDomain, " ", toDomain)
		return
	}
	isTotidExist := daoService.IsTidExist(mbean.GetToTid())
	mbean.ToTid.Domain = mbean.FromTid.Domain //只能发送到相同domain的用户
	timestamp := utils.TimeMills()
	mbean.Timestamp = &timestamp
	if isTotidExist {
		_, err, _ = route.RouteMBean(mbean, false, false)
	} else {
		err = errors.New("ResponseMessage totid not exist")
	}
	return
}

func (this *MyImpl) MessageIq(timMsgIq *MessageIq, iqType string) (err error) {
	logger.Debug("MessageIq:", timMsgIq, " ", iqType)
	panic("error")
	return
}

// Parameters:
//  - Mbean
func (this *MyImpl) MessageResult_(mbean *MBean) (err error) {
	logger.Debug("MessageResult_:", mbean)
	panic("error")
	return
}

func (this *MyImpl) Roser(roster *Roster) (err error) {
	logger.Debug("Roser:", roster)
	panic("error")
	return
}

func checkAuth(a *MyAuth) bool {
	if a.GetDomain() == ClusterConf.Domain && a.GetUsername() == ClusterConf.Username && a.GetPwd() == ClusterConf.Password {
		return true
	}
	return false
}

func (this *MyImpl) ResponseMessageIq(timMsgIq *MessageIq, iqType string, auth *MyAuth) (r *MBeanList, err error) {
	logger.Debug("ResponseMessageIq:", timMsgIq, iqType, auth)
	panic("error ResponseMessageIq")
	return
}

func (this *MyImpl) MessageList(mbeanList *MBeanList) (err error) {
	logger.Debug("MessageList:", mbeanList)
	panic("error MessageList")
	return
}

// Parameters:
//  - PbeanList
func (this *MyImpl) PresenceList(pbeanList *PBeanList) (err error) {
	logger.Debug("PresenceList:", pbeanList)
	panic("error PresenceList")
	return
}

func (this *MyImpl) ResponsePresenceList(pbeanList *PBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	if !checkAuth(auth) {
		err = errors.New(fmt.Sprint("cluster ResponsePresenceList fail:", auth))
		return
	}
	defer func() {
		if er := recover(); er != nil {
			logger.Error("error:", er)
		}
	}()
	go _MyResponsePresenceList(pbeanList, auth)
	return
}

func _MyResponsePresenceList(pbeanList *PBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if er := recover(); er != nil {
			logger.Error("error:", er)
		}
	}()
	if pbeanList != nil && pbeanList.GetPBeanList() != nil && len(pbeanList.GetPBeanList()) > 0 {
		if ClusterConf.Interflow > 0 && len(pbeanList.GetPBeanList()) > 1 {
			route.RoutePBeanList(pbeanList.GetPBeanList())
		} else {
			for _, pbean := range pbeanList.GetPBeanList() {
				_MyResponsePresence(pbean, auth)
			}
		}
	}
	return
}

// Parameters:
//  - MbeanList
//  - Auth
func (this *MyImpl) ResponseMessageList(mbeanList *MBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	if !checkAuth(auth) {
		err = errors.New(fmt.Sprint("cluster ResponseMessageList fail:", auth))
		return
	}
	defer func() {
		if er := recover(); er != nil {
			logger.Error("error:", er)
		}
	}()
	go _MyResponseMessageList(mbeanList, auth)
	return
}

func _MyResponseMessageList(mbeanList *MBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	defer func() {
		if er := recover(); er != nil {
			logger.Error("error:", er)
		}
	}()
	if mbeanList != nil && mbeanList.GetMBeanList() != nil && len(mbeanList.GetMBeanList()) > 0 {
		if ClusterConf.Interflow > 0 && len(mbeanList.GetMBeanList()) > 1 {
			route.RouteMBeanList(mbeanList.GetMBeanList(), true)
		} else {
			for _, mbean := range mbeanList.GetMBeanList() {
				route.RouteMBean(mbean, false, true)
			}
		}
	}
	return
}

func (this *MyImpl) Property(tpb *PropertyBean) (err error) {
	logger.Debug("Property:", tpb)
	return
}
