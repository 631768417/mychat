package service

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"database/sql"
	"sync"

	"chat.mysql/dao"
	"chat.mysql/dao/basedao"
	"github.com/apache/thrift/lib/go/thrift"

	"chat.DB"
	. "chat.Map"
	. "chat.common"

	"chat.connect"
	"chat.hbase"
	"chat.hbaseService"
	"chat.logger"
	. "chat.protocol"
	"chat.utils"
)

var authProviderDB *sql.DB
var once sync.Once
var domainmap *HashTable = NewHashTable()

func initAuthProviderDB() {
	logger.Info("initAuthProviderDB")
	authProviderDB, _ = DB.GetDB(CF.GetKV("my.mysql.connection", ""), 100, 10)
}

func InitDaoservice() {
	AddConf()
	updateVersion()
	if CF.DataBase == 1 {
		hbase.Init()
	}
}

/*保存离线信息列表*/
func SaveOfflineMBeanList(mbeans []*MBean) {
	if mbeans != nil && len(mbeans) > 0 {
		for _, mbean := range mbeans {
			SaveOfflineMBean(mbean)
		}
	}
}

func SaveOfflineMBean(mbean *MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SaveOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		hbaseService.SaveOfflineMBean(mbean)
	} else {
		_SaveOfflineMBean(mbean)
	}
}

/*保存离线信息*/
func _SaveOfflineMBean(mbean *MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SaveOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if mbean.GetType() == "groupchat" {
		_saveOfflineMucBean(mbean)
	} else {
		_saveOfflineMBean(mbean)
	}
}

func _saveOfflineMBean(mbean *MBean) {
	my_offline := dao.New_offline()
	mid, _ := strconv.Atoi(mbean.GetMid())
	my_offline.SetMid(int64(mid))
	my_offline.SetDomain(mbean.FromTid.GetDomain())
	my_offline.SetFromuser(mbean.GetFromTid().GetName())
	my_offline.SetCreatetime(utils.NowTime())
	my_offline.SetUsername(mbean.GetToTid().GetName())
	my_offline.SetStamp(utils.TimeMills())
	mbean.Offline = NewTime()
	mbean.Offline.Timestamp = mbean.Timestamp
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	base64string := utils.Base64Encode(stanza)
	length := len([]byte(base64string))
	my_offline.SetStanza(base64string)
	my_offline.SetMessage_size(int64(length))
	my_offline.Insert()
	go UpdateOffMessage(mbean, 0)
}

func _saveOfflineMucBean(mbean *MBean) {
	my_mucoffline := dao.New_mucoffline()
	my_mucoffline.SetCreatetime(utils.NowTime())
	my_mucoffline.SetMid(utils.Atoi64(mbean.GetMid()))
	my_mucoffline.SetDomain(mbean.GetFromTid().GetDomain())
	my_mucoffline.SetUsername(mbean.GetToTid().GetName())
	my_mucoffline.SetStamp(mbean.GetTimestamp())
	my_mucoffline.SetRoomid(mbean.GetFromTid().GetName())
	my_mucoffline.SetMsgtype(int64(mbean.GetMsgType()))
	my_mucoffline.Insert()
}

/*load 离线信息*/
func LoadOfflineMBean(tid *Tid) (mbeans []*MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("LoadOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if CF.DataBase == 1 {
		return hbaseService.LoadOfflineMBean(tid)
	} else {
		return _LoadOfflineMBean(tid)
	}
}

func _LoadOfflineMBean(tid *Tid) (mbeans []*MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("LoadOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_offline := dao.New_offline()
	my_offline.Where(my_offline.Domain.EQ(tid.GetDomain()), my_offline.Username.EQ(tid.GetName()))
	my_offline.OrderBy(my_offline.Id.Asc())
	offlines, err := my_offline.Selects()
	if err == nil {
		mbeans = make([]*MBean, 0)
		for _, of := range offlines {
			var timmbean *MBean = NewMBean()
			bb, er := utils.Base64Decode(of.GetStanza())
			if er == nil {
				thrift.NewTDeserializer().Read(timmbean, []byte(bb))
				mbeans = append(mbeans, timmbean)
			} else {
				logger.Error("Base64Decode:", er)
			}
		}
	}
	return
}

func LoadOfflineMucMBean(tid *Tid) (mbeans []*MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("LoadOfflineMucMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		return hbaseService.LoadOfflineMucMBean(tid)
	} else {
		return _LoadOfflineMucMBean(tid)
	}
}

func _LoadOfflineMucMBean(tid *Tid) (mbeans []*MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("LoadOfflineMucMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_mucoffline := dao.New_mucoffline()
	my_mucoffline.Where(my_mucoffline.Domain.EQ(tid.GetDomain()), my_mucoffline.Username.EQ(tid.GetName()))
	my_mucoffline.OrderBy(my_mucoffline.Id.Desc())
	mucofflines, err := my_mucoffline.Selects()
	if err == nil && mucofflines != nil && len(mucofflines) > 0 {
		mids := make([]interface{}, 0)
		for _, mucoffline := range mucofflines {
			mids = append(mids, mucoffline.GetMid())
		}
		my_mucmessage := dao.New_mucmessage()
		my_mucmessage.Where(my_mucmessage.Id.IN(mids...))
		mucmessages, err := my_mucmessage.Selects()
		if err == nil && mucmessages != nil && len(mucmessages) > 0 {
			mbeans := make([]*MBean, 0)
			for _, mucmsg := range mucmessages {
				var timmbean *MBean = NewMBean()
				bb, er := utils.Base64Decode(mucmsg.GetStanza())
				if er == nil {
					thrift.NewTDeserializer().Read(timmbean, []byte(bb))
					mbeans = append(mbeans, timmbean)
				} else {
					logger.Error("Base64Decode:", er)
				}
			}
		}
	}
	return
}

func LoadMucmember(roomid *Tid) (tids []*Tid) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return nil
	}
	mucRoomSQL := CF.GetKV("my.mysql.mucRoomSQL", "")
	if mucRoomSQL == "" {
		my_mucmember := dao.New_mucmember()
		my_mucmember.Where(my_mucmember.Domain.EQ(roomid.GetDomain()), my_mucmember.Roomtid.EQ(roomid.GetName()))
		my_mucmembers, err := my_mucmember.Selects()
		if err == nil && my_mucmembers != nil && len(my_mucmembers) > 0 {
			tids = make([]*Tid, 0)
			for _, r := range my_mucmembers {
				tid := NewTid()
				domain := roomid.GetDomain()
				tid.Domain = &domain
				tid.Name = r.GetTidname()
				tids = append(tids, tid)
			}
		}
	} else {
		provider()
		if authProviderDB == nil {
			logger.Error("authProviderDB is nil")
			return nil
		}
		gbbeans, err := basedao.Query(authProviderDB, mucRoomSQL, roomid.GetName())
		if err == nil && gbbeans != nil && len(gbbeans) > 0 {
			for _, gbbean := range gbbeans {
				uname := gbbean.FieldMapName["username"].ValueString()
				tid := NewTid()
				domain := roomid.GetDomain()
				tid.Domain = &domain
				tid.Name = uname
				tids = append(tids, tid)
			}
		}
	}
	return
}

func AuthMucmember(roomid, tid *Tid) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return true
	}
	mucAuthSQL := CF.GetKV("my.mysql.mucAuthSQL", "")
	if mucAuthSQL == "" {
		my_mucmember := dao.New_mucmember()
		my_mucmember.Where(my_mucmember.Domain.EQ(roomid.GetDomain()), my_mucmember.Roomtid.EQ(roomid.GetName()), my_mucmember.Tidname.EQ(tid.GetName()))
		my_mucmember.Limit(0, 1)
		gbbeans, err := my_mucmember.QueryBeen(my_mucmember.Id.Count())
		if err == nil && gbbeans != nil && len(gbbeans) > 0 {
			b = true
		}
	} else {
		provider()
		if authProviderDB == nil {
			logger.Error("authProviderDB is nil")
			return
		}
		gbbeans, err := basedao.Query(authProviderDB, mucAuthSQL, roomid.GetName(), tid.GetName())
		if err == nil && gbbeans != nil && len(gbbeans) > 0 {
			b = true
		}
	}
	return
}

func DelOfflineMBean(mid *string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if CF.DataBase == 1 {
		hbaseService.DelOfflineMBean(mid)
	} else {
		_DelOfflineMBean(mid)
	}
}

/*删除指定信息*/
func _DelOfflineMBean(mid *string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_offline := dao.New_offline()
	my_offline.Where(my_offline.Mid.EQ(mid))
	my_offline.Delete()
}

func DelOfflineMucMBean(mid *string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if CF.DataBase == 1 {
		hbaseService.DelOfflineMucMBean(mid)
	} else {
		_DelOfflineMucMBean(mid)
	}
}

func _DelOfflineMucMBean(mid *string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_mucoffline := dao.New_mucoffline()
	my_mucoffline.Where(my_mucoffline.Mid.EQ(mid))
	my_mucoffline.Delete()
}

func DelOfflineMBeanList(mids ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMBeanList,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if CF.DataBase == 1 {
		hbaseService.DelOfflineMBeanList(mids...)
	} else {
		_DelOfflineMBeanList(mids...)
	}
}

/*删除指定信息列表*/
func _DelOfflineMBeanList(mids ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMBeanList,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_offline := dao.New_offline()
	my_offline.Where(my_offline.Mid.IN(mids...))
	my_offline.Delete()
}

func DelOfflineMucMBeanList(mids ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMucMBeanList,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if CF.DataBase == 1 {
		hbaseService.DelOfflineMucMBeanList(mids...)
	} else {
		_DelOfflineMucMBeanList(mids...)
	}
}

func _DelOfflineMucMBeanList(mids ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("DelOfflineMucMBeanList,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_mucoffline := dao.New_mucoffline()
	my_mucoffline.Where(my_mucoffline.Mid.IN(mids...))
	my_mucoffline.Delete()
}

/*保存信息*/
func SaveMBean(mbean *MBean) (mid string, timestamp string, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SaveMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if CF.DataBase == 1 {
		return hbaseService.SaveMBean(mbean)
	} else {
		return _saveMBean(mbean, 1, 1)
	}

}

func SaveSingleMBean(mbean *MBean) (mid string, timestamp string, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SaveMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		return hbaseService.SaveSingleMBean(mbean)
	} else {
		return _SaveSingleMBean(mbean)
	}
}

/*保存信息*/
func _SaveSingleMBean(mbean *MBean) (mid string, timestamp string, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SaveMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		if mbean.GetMid() == "" {
			mid = fmt.Sprint(utils.GetRand(100000000))
			mbean.Mid = &mid
			timestamp = mbean.GetTimestamp()
		}
		return
	}
	fromname := mbean.FromTid.GetName()
	toname := mbean.ToTid.GetName()
	small, large := 0, 0
	if toname > fromname {
		large = 1
	} else {
		small = 1
	}
	return _saveMBean(mbean, small, large)
}

/*保存信息*/
func _saveMBean(mbean *MBean, small, large int) (mid string, timestamp string, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("_saveMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		if mbean.GetMid() == "" {
			mid := fmt.Sprint(utils.GetRand(100000000))
			mbean.Mid = &mid
		}
		return
	}
	domain := mbean.GetFromTid().GetDomain()
	fromname := mbean.GetFromTid().GetName()
	toname := mbean.GetToTid().GetName()
	message := dao.New_message()
	chatid := utils.Chatid(fromname, toname, domain)
	message.SetChatid(chatid)
	timestamp = mbean.GetTimestamp()
	message.SetStamp(timestamp)
	message.SetCreatetime(utils.NowTime())
	message.SetFromuser(fromname)
	message.SetTouser(toname)
	message.SetSmall(int64(small))
	message.SetLarge(int64(large))
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	stanzastr := string(utils.Base64Encode(stanza))
	message.SetStanza(stanzastr)
	message.Insert()
	mess := dao.New_message()
	mess.Where(mess.Stamp.EQ(timestamp), mess.Chatid.EQ(chatid))
	mess, err = mess.Select()
	if err == nil {
		mid = fmt.Sprint(mess.GetId())
		mbean.Mid = &mid
	}
	return
}

func SaveMucMBean(mbean *MBean) (mid string, err error) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("SaveMucMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		return hbaseService.SaveMucMBean(mbean)
	} else {
		return _SaveMucMBean(mbean)
	}
}

func _SaveMucMBean(mbean *MBean) (mid string, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprint(er))
			logger.Error("SaveMucMBean,", er)
			logger.Error(string(debug.Stack()))
		}
	}()
	my_mucmessage := dao.New_mucmessage()
	my_mucmessage.SetStamp(mbean.GetTimestamp())
	my_mucmessage.SetFromuser(mbean.GetLeaguerTid().GetName())
	my_mucmessage.SetRoomtidname(mbean.GetFromTid().GetName())
	my_mucmessage.SetDomain(mbean.GetLeaguerTid().GetDomain())
	my_mucmessage.SetMsgtype(int64(mbean.GetMsgType()))
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	stanzastr := string(utils.Base64Encode(stanza))
	my_mucmessage.SetStanza(stanzastr)
	my_mucmessage.SetCreatetime(utils.NowTime())
	my_mucmessage.Insert()

	mucmessage := dao.New_mucmessage()
	mucmessage.Where(mucmessage.Stamp.EQ(mbean.GetTimestamp()), mucmessage.Fromuser.EQ(mbean.LeaguerTid.GetName()), mucmessage.Domain.EQ(mbean.LeaguerTid.GetDomain()), mucmessage.Roomtidname.EQ(mbean.GetFromTid().GetName()))
	mucmessage, err = mucmessage.Select(mucmessage.Id)
	if err == nil {
		mid = fmt.Sprint(mucmessage.GetId())
		mbean.Mid = &mid
	}
	return
}

func UpdateOffMessage(mbean *MBean, status int) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("UpdateOffMessage", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		hbaseService.UpdateOffMessage(mbean, status)
	} else {
		_UpdateOffMessage(mbean, status)
	}
}

/**
  离线信息发送成功后 更新 small或large 状态
*/
func _UpdateOffMessage(mbean *MBean, status int) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("UpdateOffMessage", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	//	domain := mbean.GetFromTid().GetDomain()
	fromname := mbean.GetFromTid().GetName()
	toname := mbean.GetToTid().GetName()
	//	chatid := utils.Chatid(fromname, toname, domain)
	message := dao.New_message()
	if toname < fromname {
		message.SetSmall(int64(status))
	} else {
		message.SetLarge(int64(status))
	}
	message.Where(message.Id.EQ(mbean.GetMid()))
	message.Update()
}

func UpdateOffMessageList(mbeans []*MBean, status int) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("UpdateOffMessageList", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		hbaseService.UpdateOffMessageList(mbeans, status)
	} else {
		_UpdateOffMessageList(mbeans, status)
	}
}

/**
  离线信息发送成功后 更新 small或large 状态  列表
*/
func _UpdateOffMessageList(mbeans []*MBean, status int) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("UpdateOffMessageList", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	if len(mbeans) == 0 {
		return
	}
	fromname := mbeans[0].GetFromTid().GetName()
	toname := mbeans[0].GetToTid().GetName()
	message := dao.New_message()
	if toname < fromname {
		message.SetSmall(int64(status))
	} else {
		message.SetLarge(int64(status))
	}
	mids := make([]interface{}, 0)
	for _, mbean := range mbeans {
		mids = append(mids, mbean.GetMid())
	}
	message.Where(message.Id.IN(mids...))
	message.Update()
}

func LoadMBean(fidname, tidname, domain string, fromstamp, tostamp *string, limitcount int32) (tms []*MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		return hbaseService.LoadMBean(fidname, tidname, domain, fromstamp, tostamp, limitcount)
	} else {
		return _LoadMBean(fidname, tidname, domain, fromstamp, tostamp, limitcount)
	}
}

/***/
func _LoadMBean(fidname, tidname, domain string, fromstamp, tostamp *string, limitcount int32) (tms []*MBean) {
	logger.Debug("LoadMBean:", fidname, " ", tidname, " ", domain, " ", fromstamp, " ", tostamp, " ", limitcount)
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return nil
	}
	chatid := utils.Chatid(fidname, tidname, domain)
	isLarge := fidname > tidname
	message := dao.New_message()
	wheres := make([]*basedao.Where, 0)

	if fromstamp != nil && tostamp != nil {
		wheres = append(wheres, message.Stamp.Between(*fromstamp, *tostamp))
	} else if fromstamp != nil {
		wheres = append(wheres, message.Stamp.GT(*fromstamp))
	} else if tostamp != nil {
		wheres = append(wheres, message.Stamp.LT(*tostamp))
	}
	wheres = append(wheres, message.Chatid.EQ(chatid))
	if isLarge {
		wheres = append(wheres, message.Large.EQ(1))
	} else {
		wheres = append(wheres, message.Small.EQ(1))
	}
	message.Where(wheres...)
	message.OrderBy(message.Id.Desc())
	if limitcount > 0 {
		message.Limit(0, limitcount)
	}
	messageArray, err := message.Selects()
	if err == nil && messageArray != nil {
		tms = make([]*MBean, 0)
		for _, msg := range messageArray {
			tm := new(MBean)
			bb, er := utils.Base64Decode(msg.GetStanza())
			if er == nil {
				thrift.NewTDeserializer().Read(tm, bb)
				mid := fmt.Sprint(msg.GetId())
				tm.Mid = &mid
				tms = append(tms, tm)
			} else {
				logger.Error("Base64Decode:", er)
			}
		}
	}
	return
}

func DelMBean(fidname, tidname, domain, mid string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		hbaseService.DelMBean(fidname, tidname, domain, mid)
	} else {
		_DelMBean(fidname, tidname, domain, mid)
	}
}

func _DelMBean(fidname, tidname, domain, mid string) {
	logger.Debug("DelMBean:", fidname, " ", tidname, " ", domain, " ", mid)
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	chatid := utils.Chatid(fidname, tidname, domain)
	isLarge := fidname > tidname
	timMessage := dao.New_message()
	if isLarge {
		timMessage.SetLarge(0)
	} else {
		timMessage.SetSmall(0)
	}
	timMessage.Where(timMessage.Chatid.EQ(chatid), timMessage.Id.EQ(mid))
	timMessage.Update()
}

func DelAllMBean(fidname, tidname, domain string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.DataBase == 1 {
		hbaseService.DelAllMBean(fidname, tidname, domain)
	} else {
		_DelAllMBean(fidname, tidname, domain)
	}
}

func _DelAllMBean(fidname, tidname, domain string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	chatid := utils.Chatid(fidname, tidname, domain)
	isLarge := fidname > tidname
	timMessage := dao.New_message()
	if isLarge {
		timMessage.SetLarge(0)
	} else {
		timMessage.SetSmall(0)
	}
	timMessage.Where(timMessage.Chatid.EQ(chatid))
	timMessage.Update()
}

///*lastTime 时间之后的消息*/
//func LoadMBean(fid, tid *Tid, lastTime time.Time) (mbeans []*MBean) {
//	return
//}

/**ip地址是否被限制*/
func AllowHttpIp(ip string) bool {
	return true
}

func IsTidExist(tid *Tid) bool {
	return true
}

func ToAuth(tid *Tid, pwd string) (b bool) {
	if CF.MustAuth == 0 {
		return true
	}
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	authProvider_passwordSQL := CF.GetKV("my.mysql.passwordSQL", "")
	if authProvider_passwordSQL == "" {
		b = _auth(tid, pwd)
	} else {
		provider()
		if authProviderDB == nil {
			logger.Error("authProviderDB is nil")
			return false
		}
		for i := 0; i < 5; i++ {
			index := ""
			if i > 0 {
				index = fmt.Sprint(i)
			}
			authProvider_passwordSQL := CF.GetKV(fmt.Sprint("my.mysql.passwordSQL", index), "")
			if authProvider_passwordSQL == "" {
				continue
			}
			b = _auth4Sql(authProvider_passwordSQL, tid, pwd)
			if b {
				break
			}
		}
	}
	return
}

func _auth4Sql(authProvider_passwordSQL string, tid *Tid, pwd string) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	provider()
	if authProviderDB == nil {
		logger.Error("authProviderDB is nil")
		return false
	}
	gbbean, err := basedao.Query(authProviderDB, authProvider_passwordSQL, tid.GetName())
	if err == nil && gbbean != nil && len(gbbean) == 1 {
		if bean, ok := gbbean[0].FieldMapName["password"]; ok {
			switch CF.GetKV("authProvider.passwordType", "") {
			case "plain":
				b = eqString(bean.ValueString(), pwd)
			case "md5":
				b = eqString(bean.ValueString(), utils.MD5(pwd))
			case "sha1":
				b = eqString(bean.ValueString(), utils.Sha1(pwd))
			default:
				b = eqString(bean.ValueString(), pwd)
			}
		}
	}
	return
}

func _auth(tid *Tid, pwd string) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	loginname, _ := connect.GetLoginName(tid)
	my_user := dao.New_user()
	my_user.Where(my_user.Loginname.EQ(loginname))
	user, err := my_user.Select()
	if err == nil && user != nil {
		switch CF.GetKV("authProvider.passwordType", "") {
		case "plain":
			b = eqString(user.GetEncryptedpassword(), pwd)
		case "md5":
			b = eqString(user.GetEncryptedpassword(), utils.MD5(pwd))
		case "sha1":
			b = eqString(user.GetEncryptedpassword(), utils.Sha1(pwd))
		default:
			b = eqString(user.GetEncryptedpassword(), pwd)
		}
	}
	return
}

func eqString(s1, s2 string) bool {
	return strings.ToUpper(s1) == strings.ToUpper(s2)
}

func provider() {
	if authProviderDB == nil && CF.GetKV("my.mysql.connection", "") != "" {
		once.Do(initAuthProviderDB)
	}
}

func CheckDomain(domain string) bool {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return true
	}
	d := domainmap.Get(domain)
	if d != nil {
		if (time.Now().UnixNano()/1000000000 - d.(int64)) < 10*60 {
			return true
		} else {
			domainmap.Del(domain)
		}
	}
	my_domain := dao.New_domain()
	my_domain.Where(my_domain.Domain.EQ(domain))
	var err error
	my_domain, err = my_domain.Select()
	if err == nil && my_domain != nil && my_domain.GetId() > 0 {
		domainmap.Put(domain, time.Now().UnixNano()/1000000000)
		return true
	}
	return false
}

func AddConf() {
	logger.Debug("Addconf ok")
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	my_config := dao.New_config()
	confs, err := my_config.Selects()
	if err == nil && confs != nil && len(confs) > 0 {
		for _, conf := range confs {
			if conf.GetKeyword() != "" && conf.GetValuestr() != "" {
				CF.KV[conf.GetKeyword()] = conf.GetValuestr()
			}
		}
	}
	my_property := dao.New_property()
	propertys, err := my_property.Selects()
	if err == nil && propertys != nil && len(propertys) > 0 {
		for _, property := range propertys {
			if property.GetKeyword() != "" && (property.GetValueint() > 0 || property.GetValuestr() != "") {
				if property.GetValuestr() != "" {
					CF.KV[property.GetKeyword()] = property.GetValuestr()
				} else if property.GetValueint() > 0 {
					CF.KV[property.GetKeyword()] = fmt.Sprint(property.GetValueint())
				}
			}
		}
	}
}

//
func GetOnlineRoser(fromtid *Tid) (tids []*Tid) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return nil
	}
	domain := fromtid.GetDomain()
	fromname := fromtid.GetName()
	//	logger.Debug(domain, " ", fromname)
	authProvider_rosterSql := CF.GetKV("my.mysql.rosterSQL", "")
	loginname, _ := connect.GetLoginName(fromtid)
	if authProvider_rosterSql == "" {
		my_roster := dao.New_roster()
		my_roster.Where(my_roster.Loginname.EQ(loginname))
		rosters, err := my_roster.Selects()
		if err == nil && rosters != nil && len(rosters) > 0 {
			tids = make([]*Tid, 0)
			for _, r := range rosters {
				tid := NewTid()
				//domain := fromtid.GetDomain()
				tid.Domain = &domain
				tid.Name = r.GetRostername()
				tids = append(tids, tid)
			}
		}
	} else {
		provider()
		if authProviderDB == nil {
			logger.Error("authProviderDB is nil")
			return nil
		}
		gbbeans, err := basedao.Query(authProviderDB, authProvider_rosterSql, fromname)
		if err == nil && gbbeans != nil && len(gbbeans) > 0 {
			for _, gbbean := range gbbeans {
				uname := gbbean.FieldMapName["roster"].ValueString()
				tid := NewTid()
				domain := fromtid.GetDomain()
				tid.Domain = &domain
				tid.Name = uname
				tids = append(tids, tid)
			}
		}
	}
	return
}

func updateVersion() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	domain := dao.New_config()
	domain.Where(domain.Keyword.EQ("version"))
	td, err := domain.Select()
	if err == nil && td != nil && td.GetId() > 0 {
		domain = dao.New_config()
		domain.SetValuestr(fmt.Sprint(VersionCode))
		domain.SetRemark(fmt.Sprint(VersionName, " | ", VersionCode, " | ", utils.NowTime()))
		domain.Where(domain.Id.EQ(td.GetId()))
		domain.Update()
	} else {
		domain = dao.New_config()
		domain.SetValuestr(fmt.Sprint(VersionCode))
		domain.SetRemark(fmt.Sprint(VersionName, " | ", VersionCode, " | ", utils.NowTime()))
		domain.SetCreatetime(utils.NowTime())
		domain.SetKeyword("version")
		domain.Insert()
	}
}
