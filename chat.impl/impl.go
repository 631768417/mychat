package impl

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"

	"chat.FW"
	"chat.cluster"
	"chat.cluster/route"
	. "chat.common"
	. "chat.connect"
	"chat.logger"
	. "chat.protocol"
	"chat.route"
	"chat.thrift/client"
	"chat.utils"
	"github.com/apache/thrift/lib/go/thrift"

	daoService "chat.mysql/service"
)

/**
I 接口的实现类
*/
type Impl struct {
	Ip     string
	Port   int
	Pub    string            //发布id
	Tu     *User             //当前用户
	Client thrift.TTransport //协议
}

// Parameters:
//  - Param
func (this *Impl) Stream(param *Param) (err error) {
	if param != nil {
		if param.GetInterflow() == "1" {
			this.Tu.Interflow = 1
		}
		if param.GetTLS() == "1" {
			this.Tu.TLS = 1
		}
		this.Tu.Version = param.GetVersion()
	}
	return
}

func (this *Impl) Starttls() (err error) {
	panic("")
	return
}

// Parameters:
//  - Tid
//  - Pwd
//登录方法
func (this *Impl) Login(tid *Tid, pwd string) (err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Warn("Login error", err)
			//			logger.Error(string(debug.Stack()))
		}
	}()
	isAuth := false
	if this.Tu.Fw == FW.AUTH {
		ack := NewAckBean()
		status200, typelogin := "200", "login"
		ack.AckStatus, ack.AckType = &status200, &typelogin
		this.Tu.SendAckBean(ack)
		return
	}
	if CF.MustAuth == 0 {
		isAuth = true
	} else {
		user_auth_url := CF.GetKV("user_auth_url", "")
		if len(user_auth_url) > 9 {
			isAuth = httpAuth(tid, pwd, user_auth_url)
		} else {
			b := daoService.ToAuth(tid, pwd)
			if b {
				isAuth = true
				logger.Debug("login is success:", tid.GetName())
			}
		}
	}

	if isAuth {
		ack := NewAckBean()
		this.Tu.UserTid = tid
		this.Tu.Fw = FW.AUTH
		this.Tu.Auth(tid)
		if cluster.IsCluster() {
			loginname, _ := GetLoginName(tid)
			cluster.SetLoginnameToCluster(loginname)
		}
		status200, typelogin := "200", "login"
		ack.AckStatus, ack.AckType = &status200, &typelogin
		this.Tu.SendAckBean(ack)
		_Presence(this, OnlinePBean(this.Tu.UserTid), false)
		go route.RouteOffLineMBean(this.Tu)
	} else {
		ack := NewAckBean()
		status400, typeType := "400", "login"
		ack.AckStatus, ack.AckType = &status400, &typeType
		this.Tu.SendAckBean(ack)
		panic(fmt.Sprint("loginname or pwd is error:", tid.GetName(), " | ", pwd))
	}
	return
}

// Parameters:
//  - Ab
//应答方法
func (this *Impl) Ack(ab *AckBean) (err error) {
	if this.Tu.Fw != FW.AUTH {
		panic(fmt.Sprint("not auth:", this.Tu.Fw))
	}
	this.Tu.OverLimit = 3
	go func() {
		defer func() {
			if err := recover(); err != nil {
			}
		}()
		if CF.ConfirmAck == 1 && ab != nil && ab.GetID() == this.Tu.LastSyncThreadId {
			timer := time.NewTicker(5 * time.Second)
			select {
			case <-timer.C:
				logger.Error("ack msg threadid over time", ab)
			case this.Tu.Sendflag <- ab.GetID():
			}
		}
	}()
	return
}

// Parameters:
//  - Pbean
func (this *Impl) Presence(pbean *PBean) (err error) {
	if CF.Presence != 1 {
		return
	}
	if this.Tu.Fw != FW.AUTH {
		panic("not auth")
	}
	//	logger.Debug("pbean", pbean)
	if this.Tu.UserType == 0 {
		pbean.FromTid = this.Tu.UserTid
		_type := pbean.GetType()
		switch _type {
		case "groupchat":
			pbean.LeaguerTid = this.Tu.UserTid
		}
	}
	return _Presence(this, pbean, true)
}

func _Presence(this *Impl, pbean *PBean, isAck bool) (err error) {
	defer func() {
		if er := recover(); er != nil {
			logger.Error(string(debug.Stack()))
			err = errors.New(fmt.Sprint(er))
		}
	}()
	if CF.Presence != 1 {
		return
	}
	//ThreadId 赋值成当前毫秒
	if pbean.GetThreadId() == "" {
		pbean.ThreadId = utils.TimeMills()
	}
	//isTotidExist := daoService.IsTidExist(pbean.GetToTid())
	_type := pbean.GetType()
	switch _type {
	case "groupchat":
		pbean.FromTid = pbean.ToTid
		//default:
		//	pbean.ToTid.Domain = pbean.FromTid.Domain //只能发送到相同domain的用户
	}

	mustRoute := false
	if cluster.IsCluster() && this.Tu.UserType == 0 {
		er := clusterRoute.ClusterRoutePBean(pbean)
		if er != nil {
			mustRoute = true
		}
	} else {
		mustRoute = true
	}
	if mustRoute {
		if pbean.GetToTid() == nil {
			route.RoutePBean(pbean)
		} else {
			route.RouteSinglePBean(pbean)
		}
		if isAck {
			ack := NewAckBean()
			id := pbean.ThreadId
			ack.ID = &id
			status200, typemessage := "200", "presence"
			ack.AckStatus, ack.AckType = &status200, &typemessage
			this.Tu.SendAckBean(ack)
		}
	}
	return
}

// Parameters:
//  - Mbean
//发消息
func (this *Impl) Message(mbean *MBean) (err error) {
	if this.Tu.Fw != FW.AUTH {
		panic("not auth")
	}
	//	logger.Debug("Message=====>", mbean)
	if this.Tu.UserType == 0 {
		mbean.FromTid = this.Tu.UserTid
		//		isTotidExist := daoService.IsTidExist(mbean.GetToTid())
		_type := mbean.GetType()
		switch _type {
		case "groupchat":
			//验证是否是群组的一员
			b := daoService.AuthMucmember(mbean.GetToTid(), this.Tu.UserTid)
			if !b {
				panic("auth room failed")
			}
		}
	}
	//	if isTotidExist {
	//		id, _, _ := route.RouteMBean(mbean, false, false)
	//		ack := NewAckBean()
	//		status200, typemessage := "200", "message"
	//		ack.AckStatus, ack.AckType = &status200, &typemessage
	//		ack.ExtraMap = make(map[string]string, 0)
	//		ack.ExtraMap["mid"] = fmt.Sprint(id)
	//		this.Tu.SendAckBean(ack)
	//	}
	return _Message(this, mbean)
}

//发消息核心方法
func _Message(this *Impl, mbean *MBean) (err error) {
	//ThreadId 赋值
	if mbean.GetThreadId() == "" {
		mbean.ThreadId = utils.TimeMills()
	}
	isTotidExist := daoService.IsTidExist(mbean.GetToTid())
	_type := mbean.GetType()
	switch _type {
	//群聊
	case "groupchat":
		mbean.FromTid = mbean.ToTid
		mbean.LeaguerTid = this.Tu.UserTid
		mbean.FromTid.Domain = this.Tu.UserTid.Domain
		mbean.ToTid = nil
		timestamp := utils.TimeMills()
		mbean.Timestamp = &timestamp
	default:
		//单聊
		mbean.ToTid.Domain = mbean.FromTid.Domain //只能发送到相同domain的用户
		timestamp := utils.TimeMills()
		mbean.Timestamp = &timestamp
	}

	if isTotidExist && mbean.GetToTid() != nil {
		mustRoute := true
		//是否是集群
		if cluster.IsCluster() {
			clusterBean := clusterRoute.OtherClusterUserBean(mbean.GetToTid())
			if this.Tu.UserType == 0 && clusterBean != nil {
				//置入集群路由
				er := clusterRoute.ClusterRouteMBean(mbean, clusterBean)
				if er != nil {
					mustRoute = true
				} else {
					mustRoute = false
				}
			} else {
				mustRoute = true
			}
		}

		if mustRoute {
			id, er, _ := route.RouteMBean(mbean, false, true)
			ack := NewAckBean()
			thid := mbean.ThreadId
			ack.ID = &thid
			if er == nil {
				status, typemessage := SC_SUCCESS, "message"
				ack.AckStatus, ack.AckType = &status, &typemessage
				ack.ExtraMap = make(map[string]string, 0)
				ack.ExtraMap["mid"] = fmt.Sprint(id)
			} else {
				status, typemessage := SC_FAILED, "message"
				ack.AckStatus, ack.AckType = &status, &typemessage
			}
			this.Tu.SendAckBean(ack)
		}
	}
	return
}

// Parameters:
//  - ThreadId
func (this *Impl) Ping(threadId string) (err error) {
	if this.Tu.Fw != FW.AUTH {
		panic("not auth")
	}
	//	logger.Debug("ping>>>>>", threadId)
	ab := NewAckBean()
	ab.ID = &threadId
	acktype, ackstatus := "ping", "200"
	ab.AckType, ab.AckStatus = &acktype, &ackstatus
	this.Tu.SendAckBean(ab)
	return
}

// Parameters:
//  - E
func (this *Impl) Error(e *Error) (err error) {
	panic("Error")
	return
}

func (this *Impl) Logout() (err error) {
	panic("Logout")
	return
}

// Parameters:
//  - Tid
//  - Pwd
func (this *Impl) Regist(tid *Tid, pwd string) (err error) {
	panic("error Regist")
	return
}

// Parameters:
//  - Tid
//  - Pwd
func (this *Impl) RemoteUserAuth(tid *Tid, pwd string, auth *MyAuth) (r *RemoteUserBean, err error) {
	panic("error RemoteUserAuth")
	return
}

// Parameters:
//  - Tid
func (this *Impl) RemoteUserGet(tid *Tid, auth *MyAuth) (r *RemoteUserBean, err error) {
	panic("error RemoteUserGet")
	return
}

// Parameters:
//  - Tid
//  - Ub
func (this *Impl) RemoteUserEdit(tid *Tid, ub *UserBean, auth *MyAuth) (r *RemoteUserBean, err error) {
	panic("error RemoteUserEdit")
	return
}

// Parameters:
//  - Pbean
func (this *Impl) ResponsePresence(pbean *PBean, auth *MyAuth) (r *ResponseBean, err error) {
	panic("ResponsePresence")
	return
}

// Parameters:
//  - Mbean
//拼装 回应数据
func (this *Impl) ResponseMessage(mbean *MBean, auth *MyAuth) (r *ResponseBean, err error) {
	r = NewResponseBean()
	fromDomain := mbean.GetFromTid().GetDomain()
	toDomain := mbean.GetToTid().GetDomain()
	//必须保证在同一个域发送
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
		isSinglePush := false
		if mbean.ExtraMap != nil {
			if pushType, ok := mbean.ExtraMap["pushType"]; ok {
				if pushType == "single" {
					isSinglePush = true
				}
				delete(mbean.ExtraMap, "pushType")
			}
		}

		id, er, offline := route.RouteMBean(mbean, isSinglePush, false)

		if er == nil {
			r.ExtraMap = make(map[string]string, 0)
			r.ExtraMap["mid"] = fmt.Sprint(id)
			r.ExtraMap["timestamp"] = timestamp
			if offline {
				r.ExtraMap["offline"] = "1"
			} else {
				r.ExtraMap["offline"] = "0"
			}
		}
	}
	return
}

func (this *Impl) MessageIq(msgIq *MessageIq, iqType string) (err error) {
	//	logger.Debug("MessageIq:", msgIq, " ", iqType)
	switch iqType {
	case "get":
		fidname := this.Tu.UserTid.GetName()
		domain := this.Tu.UserTid.Domain
		tidnames := msgIq.Tidlist
		limitcount := msgIq.Page.LimitCount
		fromstamp := msgIq.Page.FromTimeStamp
		tostamp := msgIq.Page.ToTimeStamp
		if tidnames != nil {
			for _, tidname := range tidnames {
				mbeans := daoService.LoadMBean(fidname, tidname, *domain, fromstamp, tostamp, *limitcount)
				//1合流 0不合流
				if this.Tu.Interflow > 0 {
					//消息合流发送
					mbeanlist := NewMBeanList()
					mbeanlist.ThreadId = utils.TimeMills()
					mbeanlist.MBeanList = mbeans
					this.Tu.Client.MessageList(mbeanlist)
				} else {
					//不合流发送
					if mbeans != nil {
						for _, mbean := range mbeans {
							er := this.Tu.Client.MessageResult_(mbean)
							if er != nil {
								break
							}
						}
					}
				}
			}
		}
	case "del":
		fidname := this.Tu.UserTid.GetName()
		domain := this.Tu.UserTid.Domain
		tidnames := msgIq.Tidlist
		mids := msgIq.Midlist
		if len(tidnames) == 1 && len(mids) == 1 {
			daoService.DelMBean(fidname, tidnames[0], *domain, mids[0])
		}
	case "delAll":
		fidname := this.Tu.UserTid.GetName()
		domain := this.Tu.UserTid.Domain
		tidnames := msgIq.Tidlist
		if len(tidnames) == 1 {
			daoService.DelAllMBean(fidname, tidnames[0], *domain)
		}
	default:
		panic("error iqType")
	}
	return
}

// Parameters:
//  - Mbean
func (this *Impl) MessageResult_(mbean *MBean) (err error) {
	logger.Debug("MessageResult_:", mbean)
	panic("error MessageResult_")
	return
}

func (this *Impl) Roser(roster *Roster) (err error) {
	logger.Debug("Roser:", roster)
	panic("error Roser")
	return
}

func (this *Impl) ResponseMessageIq(msgIq *MessageIq, iqType string, auth *MyAuth) (r *MBeanList, err error) {
	//	logger.Debug("ResponseMessageIq:", msgIq, iqType, auth)
	user_auth_url := CF.GetKV("user_auth_url", "")
	isAuth := false
	tid := NewTid()
	tid.Domain, tid.Name = auth.Domain, auth.GetUsername()
	pwd := auth.GetPwd()
	if user_auth_url != "" {
		isAuth = httpAuth(tid, pwd, user_auth_url)
	} else {
		b := daoService.ToAuth(tid, pwd)
		if b {
			isAuth = true
		}
	}
	if !isAuth {
		return
	}
	switch iqType {
	case "offline":
		r = NewMBeanList()
		mbeans := daoService.LoadOfflineMBean(tid)
		r.ThreadId = utils.TimeMills()
		r.MBeanList = mbeans
		mids := make([]interface{}, 0)
		for _, mbean := range mbeans {
			mids = append(mids, mbean.GetMid())
		}
		daoService.DelOfflineMBeanList(mids...)
		daoService.UpdateOffMessageList(mbeans, 1)
	case "get":
	}
	return
}

func httpAuth(tid *Tid, pwd, user_auth_url string) (isAuth bool) {
	var r *RemoteUserBean
	tfClient.HttpClient(func(client *IClient) (er error) {
		defer func() {
			if err := recover(); err != nil {
				er = errors.New(fmt.Sprint(err))
				logger.Error(string(debug.Stack()))
			}
		}()
		r, er = client.RemoteUserAuth(tid, pwd, nil)
		if er == nil && r != nil {
			logger.Debug(r)
			if r.ExtraMap != nil {
				if password, ok := r.ExtraMap["password"]; ok {
					if pwd == password {
						isAuth = true
					}
				}
				if extraAuth, ok := r.ExtraMap["extraAuth"]; ok {
					if pwd == extraAuth {
						isAuth = true
					}
				}
			}
		}
		return er
	}, user_auth_url)
	return
}

func (this *Impl) MessageList(mbeanList *MBeanList) (err error) {
	logger.Debug("MessageList:", mbeanList)
	panic("error MessageList")
	return
}

// Parameters:
//  - PbeanList
func (this *Impl) PresenceList(pbeanList *PBeanList) (err error) {
	logger.Debug("PresenceList:", pbeanList)
	panic("error PresenceList")
	return
}

func (this *Impl) ResponsePresenceList(pbeanList *PBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	logger.Debug("ResponsePresenceList:", pbeanList)
	panic("error ResponsePresenceList")
	return
}

// Parameters:
//  - MbeanList
//  - Auth
func (this *Impl) ResponseMessageList(mbeanList *MBeanList, auth *MyAuth) (r *ResponseBean, err error) {
	logger.Debug("ResponseMessageList:", mbeanList)
	panic("error ResponseMessageList")
	return
}

func (this *Impl) Property(tpb *PropertyBean) (err error) {
	if this.Tu.Fw != FW.AUTH {
		panic("not auth")
	}
	//	logger.Debug("Property:", tpb)
	interflow := tpb.GetInterflow()
	tls := tpb.GetTLS()
	if interflow == "1" {
		this.Tu.Interflow = 1
	}
	if tls == "1" {
		this.Tu.TLS = 1
	}
	return
}
