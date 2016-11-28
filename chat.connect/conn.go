package connect

import (
	"errors"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	. "chat.FW"
	. "chat.common"
	"chat.logger"

	. "chat.protocol"
	. "chat.utils"
)

type Pool struct {
	// 注册了的连接器
	pool map[*User]bool
	//	pool *HashTable
	// 在线用户  domain+name  用户池
	poolUser map[string]map[*User]bool
	//	poolUser *HashTable
	// 从连接器中注册请求
	register chan *User
	// 从连接器中注销请求
	unregister chan *User
	// 锁
	rwLock *sync.RWMutex
}

func (this *Pool) AddAuthUser(tu *User) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("AddAuthUser,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	if _, ok := this.pool[tu]; ok {
		loginname, _ := GetLoginName(tu.UserTid)
		if len(this.poolUser[loginname]) == 0 {
			this.poolUser[loginname] = make(map[*User]bool)
		}
		this.poolUser[loginname][tu] = true
	}
}

func (this *Pool) AddSingleAuthUser(tu *User) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("AddSingleAuthUser,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	loginname, _ := GetLoginName(tu.UserTid)
	if pu, ok := this.poolUser[loginname]; ok {
		if pu != nil && len(pu) > 0 {
			for k, _ := range pu {
				k.Close()
				delete(pu, k)
				delete(this.pool, k)
			}
		}
	}
	if _, ok := this.pool[tu]; ok {
		if this.poolUser[loginname] == nil {
			this.poolUser[loginname] = make(map[*User]bool)
		}
		this.poolUser[loginname][tu] = true
	}
}

func (this *Pool) GetAllLoginName() (ss []string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("GetAllLoginName,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	//	this.rwLock.RLock()
	//	defer this.rwLock.RUnlock()
	ss = make([]string, 0)
	for k, _ := range this.poolUser {
		ss = append(ss, k)
	}
	return
}

func (this *Pool) DeleteUser(tu *User) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DeleteUser,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	this.rwLock.Lock()
	defer this.rwLock.Unlock()
	tu.Close()
	if tu.UserTid != nil {
		loginname, _ := GetLoginName(tu.UserTid)
		if tumap, ok := this.poolUser[loginname]; ok {
			if _, ok := tumap[tu]; ok {
				delete(tumap, tu)
				if len(tumap) == 0 {
					delete(this.poolUser, loginname)
				}
			}
		}
	}
	if _, ok := this.pool[tu]; ok {
		delete(this.pool, tu)
	}
}

func (t *Pool) Len4PU() int {
	t.rwLock.RLock()
	defer t.rwLock.RUnlock()
	return len(t.poolUser)
}

func (t *Pool) Len4P() int {
	t.rwLock.RLock()
	defer t.rwLock.RUnlock()
	return len(t.pool)
}

func (t *Pool) PrintUsersInfo() string {
	t.rwLock.RLock()
	defer t.rwLock.RUnlock()
	str := fmt.Sprintln(len(t.pool), "======>", len(t.poolUser))
	for _, vmap := range t.poolUser {
		for tu, _ := range vmap {
			str = fmt.Sprintln(str, tu.UserTid.GetName(), " # ", TimeMills2TimeFormat(tu.IdCardNo), " # ", tu.UserTid.GetResource())
		}
	}
	return str
}

func (t *Pool) GetLoginUser(loginname string) (tus []*User) {
	t.rwLock.RLock()
	defer t.rwLock.RUnlock()
	if tumap, ok := t.poolUser[loginname]; ok {
		tus = make([]*User, 0)
		for tu, _ := range tumap {
			tus = append(tus, tu)
		}
	}
	return
}

func (t *Pool) AddConnect(c *User) {
	t.rwLock.Lock()
	defer t.rwLock.Unlock()
	t.pool[c] = true
}

type User struct {
	UserTid          *Tid
	Client           *IClient
	Fw               FLOW
	OverLimit        int
	IdCardNo         string
	IsClose          bool
	Sendflag         chan string
	Sync             *sync.Mutex
	LastSyncThreadId string
	UserType         int // 0 client  1 cluster client
	TLS              int //
	Interflow        int //0为不合流 1为合流
	Version          int16
}

func (t *User) Auth(tid *Tid) (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Auth,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("Auth err")
		}
	}()
	t.Sync.Lock()
	defer t.Sync.Unlock()
	t.UserTid = tid
	if CF.SingleClient == 1 {
		TP.AddSingleAuthUser(t)
	} else {
		TP.AddAuthUser(t)
	}
	return
}

func (t *User) SendMBean(mbean *MBean) (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("sendMBean error:", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("sendMBean err")
		}
	}()
	t.Sync.Lock()
	defer t.Sync.Unlock()
	if t.IsClose {
		er = errors.New("timuser is close")
		return
	}
	if CF.ConfirmAck == 1 {
		timer := time.NewTicker(3 * time.Second)
		t.LastSyncThreadId = mbean.GetThreadId()
		er = t.Client.Message(mbean)
		if er == nil {
			select {
			case <-timer.C:
				er = errors.New(fmt.Sprint(t.UserTid.GetName(), ", send ack overtime:", mbean.GetThreadId(), "  ", t))
				logger.Error("send ack overtime:", mbean.GetThreadId())
			case threadId := <-t.Sendflag:
				if t.LastSyncThreadId != threadId {
					er = errors.New(fmt.Sprint("error msg ack threadid:", t.LastSyncThreadId, "!=", threadId))
					logger.Error("error msg ack threadid:", t.LastSyncThreadId, "!=", threadId)
				}
			}
		} else {
			logger.Error("sendMBean:", er.Error())
		}
	} else {
		er = t.Client.Message(mbean)
	}
	if er != nil {
		t.IsClose = true
		t.Fw = CLOSE
	}
	return
}

func (t *User) SendMBeanList(mbeans []*MBean) (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendMBeanList error:", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("SendMBeanList err")
		}
	}()
	t.Sync.Lock()
	defer t.Sync.Unlock()
	if t.IsClose {
		er = errors.New("timuser is close")
		return
	}
	mbeanList := NewMBeanList()
	mbeanList.MBeanList = mbeans
	mbeanList.ThreadId = TimeMills()
	if CF.ConfirmAck == 1 {
		timer := time.NewTicker(5 * time.Second)
		t.LastSyncThreadId = mbeanList.GetThreadId()
		er = t.Client.MessageList(mbeanList)
		select {
		case <-timer.C:
			er = errors.New(fmt.Sprint("send ack overtime:", mbeanList.GetThreadId()))
			logger.Error("send ack overtime:", mbeanList.GetThreadId())
		case threadId := <-t.Sendflag:
			if mbeanList.GetThreadId() != threadId {
				er = errors.New(fmt.Sprint("error msg ack threadid:", mbeanList.GetThreadId(), "!=", threadId))
				logger.Error("error msg ack threadid:", mbeanList.GetThreadId(), "!=", threadId)
			}
		}
	} else {
		er = t.Client.MessageList(mbeanList)
	}
	if er != nil {
		t.IsClose = true
		t.Fw = CLOSE
	}
	return
}

func (t *User) Ping() (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Ping,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("Ping err")
		}
	}()
	t.Sync.Lock()
	defer t.Sync.Unlock()
	er = t.Client.Ping(TimeMills())
	return
}

func (t *User) SendPBean(pbean *PBean) (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("sendPBean,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("sendPBean err")
		}
	}()
	if CF.Presence != 1 {
		return
	}
	t.Sync.Lock()
	defer t.Sync.Unlock()
	er = t.Client.Presence(pbean)
	return
}

func (t *User) SendPBeanList(pbean []*PBean) (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SendPBeanList,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("SendPBeanList err")
		}
	}()
	if CF.Presence != 1 {
		return
	}
	t.Sync.Lock()
	defer t.Sync.Unlock()
	pbeanList := NewPBeanList()
	pbeanList.ThreadId = TimeMills()
	pbeanList.PBeanList = pbean
	er = t.Client.PresenceList(pbeanList)
	return
}

func (t *User) SendAckBean(ackBean *AckBean) (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("sendAckBean,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("sendAckBean err")
		}
	}()
	t.Sync.Lock()
	defer t.Sync.Unlock()
	er = t.Client.Ack(ackBean)
	return
}

func (t *User) Close() (er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Close,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("Close err")
		}
	}()
	t.Sync.Lock()
	defer t.Sync.Unlock()
	er = t.Client.Transport.Close()
	return
}

//获取登录名
func GetLoginName(tid *Tid) (loginname string, er error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("GetLoginName,", err)
			logger.Error(string(debug.Stack()))
			er = errors.New("GetLoginName err")
		}
	}()
	if tid == nil || tid.GetDomain() == "" || tid.GetName() == "" {
		return "", errors.New("tid error")
	}
	domain := tid.GetDomain()
	name := tid.GetName()
	//登录名 =  domain + my盐值 + name
	loginname = MD5(fmt.Sprint(domain, "my", name))
	return
}

var TP = Pool{
	register:   make(chan *User),
	unregister: make(chan *User),
	pool:       make(map[*User]bool),
	poolUser:   make(map[string]map[*User]bool),
	rwLock:     new(sync.RWMutex),
}
