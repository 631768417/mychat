package hbaseService

import (
	"errors"
	"fmt"
	"runtime/debug"

	. "chat.common"
	"chat.hbase"
	"chat.logger"
	. "chat.protocol"
	"chat.utils"
	"github.com/apache/thrift/lib/go/thrift"
)

/*保存离线信息列表*/
func SaveOfflineMBeanList(mbeans []*MBean) {
	if mbeans != nil && len(mbeans) > 0 {
		for _, mbean := range mbeans {
			SaveOfflineMBean(mbean)
		}
	}
}

/*保存离线信息*/
func SaveOfflineMBean(mbean *MBean) {
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
	/**
	my_offline := dao.NewTim_offline()
	mid, _ := strconv.Atoi(mbean.GetMid())
	my_offline.SetMid(int64(mid))
	my_offline.SetDomain(mbean.FromTid.GetDomain())
	my_offline.SetFromuser(mbean.GetFromTid().GetName())
	my_offline.SetCreatetime(utils.NowTime())
	my_offline.SetUsername(mbean.GetToTid().GetName())
	my_offline.SetStamp(utils.TimeMills())
	mbean.Offline = NewTimTime()
	mbean.Offline.Timestamp = mbean.Timestamp
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	base64string := base64Util.Base64Encode(stanza)
	length := len([]byte(base64string))
	my_offline.SetStanza(base64string)
	my_offline.SetMessage_size(int64(length))
	my_offline.Insert()
	*/
	my_offline := new(hbase.Tim_offline)
	my_offline.Mid = fmt.Sprint(mbean.GetMid())
	my_offline.Domain = mbean.FromTid.GetDomain()
	my_offline.Fromuser = mbean.FromTid.GetName()
	my_offline.Createtime = utils.NowTime()
	my_offline.Username = mbean.ToTid.GetName()
	my_offline.Stamp = utils.TimeMills()
	mbean.Offline = NewTime()
	mbean.Offline.Timestamp = mbean.Timestamp
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	base64string := utils.Base64Encode(stanza)
	length := len([]byte(base64string))
	my_offline.Stanza = base64string
	my_offline.Message_size = fmt.Sprint(length)
	my_offline.IndexMid = my_offline.Mid
	my_offline.IndexDomainUsername = utils.MD5(fmt.Sprint(my_offline.Domain, "_idx_", my_offline.Username))
	my_offline.Insert()
	go UpdateOffMessage(mbean, 0)
}

func _saveOfflineMucBean(mbean *MBean) {
	/***
	my_mucoffline := dao.NewTim_mucoffline()
	my_mucoffline.SetCreatetime(utils.NowTime())
	my_mucoffline.SetMid(utils.Atoi64(mbean.GetMid()))
	my_mucoffline.SetDomain(mbean.GetFromTid().GetDomain())
	my_mucoffline.SetUsername(mbean.GetToTid().GetName())
	my_mucoffline.SetStamp(mbean.GetTimestamp())
	my_mucoffline.SetRoomid(mbean.GetFromTid().GetName())
	my_mucoffline.SetMsgtype(int64(mbean.GetMsgType()))
	my_mucoffline.Insert()
	*/
	my_mucoffline := new(hbase.Tim_mucoffline)
	my_mucoffline.Createtime = utils.NowTime()
	my_mucoffline.Mid = mbean.GetMid()
	my_mucoffline.Domain = mbean.GetFromTid().GetDomain()
	my_mucoffline.Username = mbean.GetToTid().GetName()
	my_mucoffline.Stamp = mbean.GetTimestamp()
	my_mucoffline.Roomid = mbean.GetFromTid().GetName()
	my_mucoffline.Msgtype = fmt.Sprint(mbean.GetMsgType())
	my_mucoffline.IndexMid = my_mucoffline.Mid
	my_mucoffline.IndexDomainUsername = utils.MD5(fmt.Sprint(my_mucoffline.Domain, "_idx_", my_mucoffline.Username))
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
	/**
	my_offline := dao.NewTim_offline()
	my_offline.Where(my_offline.Domain.EQ(tid.GetDomain()), my_offline.Username.EQ(tid.GetName()))
	my_offline.OrderBy(my_offline.Id.Asc())
	offlines, err := my_offline.Selects()
	if err == nil {
		mbeans = make([]*MBean, 0)
		for _, of := range offlines {
			var timmbean *MBean = NewMBean()
			bb, er := base64Util.Base64Decode(of.GetStanza())
			if er == nil {
				thrift.NewTDeserializer().Read(timmbean, []byte(bb))
				mbeans = append(mbeans, timmbean)
			} else {
				logger.Error("Base64Decode:", er)
			}
		}
	}*/
	my_offline := new(hbase.Tim_offline)
	bean := new(hbase.Bean)
	bean.Family = "index"
	bean.Qualifier = utils.MD5(fmt.Sprint(tid.GetDomain(), "_idx_", tid.GetName()))
	rs, er := hbase.ScansFromRow(my_offline.Tablename(), []*hbase.Bean{bean}, 0, false)
	if er == nil {
		mbeans = make([]*MBean, 0)
		for _, r := range rs {
			//printResult(r)
			t := new(hbase.Tim_offline)
			hbase.Result2object(r, t)
			var timmbean *MBean = NewMBean()
			bb, er := utils.Base64Decode(t.Stanza)
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

func printResult(result *hbase.TResult_) {
	for _, resultColumnValue := range result.GetColumnValues() {
		logger.Error("printResult===>", hbase.Bytes2hex(result.GetRow()), " | ", "family==", string(resultColumnValue.GetFamily()), " | ", "qualifier==", string(resultColumnValue.GetQualifier()), " | ", "value==", string(resultColumnValue.GetValue()), " | ", "timestamp==", resultColumnValue.GetTimestamp())
	}
}

func LoadOfflineMucMBean(tid *Tid) (mbeans []*MBean) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("LoadOfflineMucMBean,", err)
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	/***
	my_mucoffline := dao.NewTim_mucoffline()
	my_mucoffline.Where(my_mucoffline.Domain.EQ(tid.GetDomain()), my_mucoffline.Username.EQ(tid.GetName()))
	my_mucoffline.OrderBy(my_mucoffline.Id.Desc())
	mucofflines, err := my_mucoffline.Selects()
	if err == nil && mucofflines != nil && len(mucofflines) > 0 {
		mids := make([]interface{}, 0)
		for _, mucoffline := range mucofflines {
			mids = append(mids, mucoffline.GetMid())
		}
		my_mucmessage := dao.NewTim_mucmessage()
		my_mucmessage.Where(my_mucmessage.Id.IN(mids...))
		mucmessages, err := my_mucmessage.Selects()
		if err == nil && mucmessages != nil && len(mucmessages) > 0 {
			mbeans := make([]*MBean, 0)
			for _, mucmsg := range mucmessages {
				var timmbean *MBean = NewMBean()
				bb, er := base64Util.Base64Decode(mucmsg.GetStanza())
				if er == nil {
					thrift.NewTDeserializer().Read(timmbean, []byte(bb))
					mbeans = append(mbeans, timmbean)
				} else {
					logger.Error("Base64Decode:", er)
				}
			}
		}
	}**/
	my_mucoffline := new(hbase.Tim_mucoffline)
	rs, er := hbase.Selects(my_mucoffline.Tablename(), "index", utils.MD5(fmt.Sprint(tid.GetDomain(), "_idx_", tid.GetName())), 0, false)
	if er == nil {
		mbeans = make([]*MBean, 0)
		mids := make([]int64, 0)
		for _, r := range rs {
			t := new(hbase.Tim_mucoffline)
			hbase.Result2object(r, t)
			mids = append(mids, utils.Atoi64(t.Mid))
		}
		my_mucmessage := new(hbase.Tim_mucmessage)
		rss, err := hbase.SelectByRows(my_mucmessage.Tablename(), mids)
		if err == nil {
			for _, r := range rss {
				t := new(hbase.Tim_mucmessage)
				hbase.Result2object(r, t)
				var timmbean *MBean = NewMBean()
				bb, er := utils.Base64Decode(t.Stanza)
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

/***
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
		my_mucmember := dao.NewTim_mucmember()
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
*/

/***
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
		my_mucmember := dao.NewTim_mucmember()
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
*/

/*删除指定信息*/
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
	/***
	my_offline := dao.NewTim_offline()
	my_offline.Where(my_offline.Mid.EQ(mid))
	my_offline.Delete()*/
	row := utils.Atoi64(*mid)
	my_offline := new(hbase.Tim_offline)
	my_offline.Delete(row)
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
	//	my_mucoffline := dao.NewTim_mucoffline()
	//	my_mucoffline.Where(my_mucoffline.Mid.EQ(mid))
	//	my_mucoffline.Delete()
	my_mucoffline := new(hbase.Tim_mucoffline)
	my_mucoffline.Delete(utils.Atoi64(*mid))
}

/*删除指定信息列表*/
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
	//	my_offline := dao.NewTim_offline()
	//	my_offline.Where(my_offline.Mid.IN(mids...))
	//	my_offline.Delete()
	//	rows := make([]int64, 0)
	beans := make([]*hbase.Bean, 0)
	for _, mid := range mids {
		bean := new(hbase.Bean)
		bean.Family = "index"
		bean.Qualifier = fmt.Sprint(mid)
		beans = append(beans, bean)
	}
	my_offline := new(hbase.Tim_offline)
	//	my_offline.Deletes(rows)
	my_offline.DeleteByBean(beans)
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
	//	my_mucoffline := dao.NewTim_mucoffline()
	//	my_mucoffline.Where(my_mucoffline.Mid.IN(mids...))
	//	my_mucoffline.Delete()
	rows := make([]int64, 0)
	for _, mid := range mids {
		rows = append(rows, utils.Atoi64(fmt.Sprint(mid)))
	}
	my_mucoffline := new(hbase.Tim_mucoffline)
	my_mucoffline.Deletes(rows)
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
	return _saveMBean(mbean, 1, 1)
}

/*保存信息*/
func SaveSingleMBean(mbean *MBean) (mid string, timestamp string, err error) {
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
	/**
	domain := mbean.GetFromTid().GetDomain()
	fromname := mbean.GetFromTid().GetName()
	toname := mbean.GetToTid().GetName()
	message := dao.NewTim_message()
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
	stanzastr := string(base64Util.Base64Encode(stanza))
	message.SetStanza(stanzastr)
	message.Insert()
	mess := dao.NewTim_message()
	mess.Where(mess.Stamp.EQ(timestamp), mess.Chatid.EQ(chatid))
	var err error
	mess, err = mess.Select()
	if err == nil {
		mid = fmt.Sprint(mess.GetId())
		mbean.Mid = &mid
	}*/
	domain := mbean.GetFromTid().GetDomain()
	fromname := mbean.GetFromTid().GetName()
	toname := mbean.GetToTid().GetName()
	message := new(hbase.Tim_message)
	chatid := utils.Chatid(fromname, toname, domain)
	message.Chatid = chatid
	timestamp = mbean.GetTimestamp()
	message.Stamp = timestamp
	message.Createtime = utils.NowTime()
	message.Fromuser = fromname
	message.Touser = toname
	message.Small = fmt.Sprint(small)
	message.Large = fmt.Sprint(large)
	message.Msgtype = fmt.Sprint(mbean.GetMsgType())
	message.Msgmode = "1"
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	stanzastr := string(utils.Base64Encode(stanza))
	message.Stanza = stanzastr
	message.IndexChatid = chatid
	var row int64
	row, err = message.Insert()
	if err == nil {
		mid = fmt.Sprint(row)
	}
	return
}

func SaveMucMBean(mbean *MBean) (mid string, err error) {
	defer func() {
		if er := recover(); er != nil {
			err = errors.New(fmt.Sprint(er))
			logger.Error("SaveMucMBean,", er)
			logger.Error(string(debug.Stack()))
		}
	}()
	/***
	my_mucmessage := dao.NewTim_mucmessage()
	my_mucmessage.SetStamp(mbean.GetTimestamp())
	my_mucmessage.SetFromuser(mbean.GetLeaguerTid().GetName())
	my_mucmessage.SetRoomtidname(mbean.GetFromTid().GetName())
	my_mucmessage.SetDomain(mbean.GetLeaguerTid().GetDomain())
	my_mucmessage.SetMsgtype(int64(mbean.GetMsgType()))
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	stanzastr := string(base64Util.Base64Encode(stanza))
	my_mucmessage.SetStanza(stanzastr)
	my_mucmessage.SetCreatetime(utils.NowTime())
	my_mucmessage.Insert()

	mucmessage := dao.NewTim_mucmessage()
	mucmessage.Where(mucmessage.Stamp.EQ(mbean.GetTimestamp()), mucmessage.Fromuser.EQ(mbean.LeaguerTid.GetName()), mucmessage.Domain.EQ(mbean.LeaguerTid.GetDomain()), mucmessage.Roomtidname.EQ(mbean.GetFromTid().GetName()))
	var err error
	mucmessage, err = mucmessage.Select(mucmessage.Id)
	if err == nil {
		mid = fmt.Sprint(mucmessage.GetId())
		mbean.Mid = &mid
	}*/
	my_mucmessage := new(hbase.Tim_mucmessage)
	my_mucmessage.Stamp = mbean.GetTimestamp()
	my_mucmessage.Fromuser = mbean.GetLeaguerTid().GetName()
	my_mucmessage.Roomtidname = mbean.GetFromTid().GetName()
	my_mucmessage.Domain = mbean.GetLeaguerTid().GetDomain()
	my_mucmessage.Msgtype = fmt.Sprint(mbean.GetMsgType())
	stanza, _ := thrift.NewTSerializer().Write(mbean)
	stanzastr := string(utils.Base64Encode(stanza))
	my_mucmessage.Stanza = stanzastr
	my_mucmessage.Createtime = utils.NowTime()
	my_mucmessage.IndexFromuserDomain = utils.MD5(fmt.Sprint(my_mucmessage.Fromuser, "_idx_", my_mucmessage.Domain))
	var row int64
	row, err = my_mucmessage.Insert()
	if err == nil {
		mid = fmt.Sprint(row)
		mbean.Mid = &mid
	}
	return
}

/**
  离线信息发送成功后 更新 small或large 状态
*/
func UpdateOffMessage(mbean *MBean, status int) {
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
	/***
	message := dao.NewTim_message()
	if toname < fromname {
		message.SetSmall(int64(status))
	} else {
		message.SetLarge(int64(status))
	}
	message.Where(message.Id.EQ(mbean.GetMid()))
	message.Update()
	*/
	message := new(hbase.Tim_message)
	if toname < fromname {
		message.Small = fmt.Sprint(status)
	} else {
		message.Large = fmt.Sprint(status)
	}
	message.Update(utils.Atoi64(mbean.GetMid()))
}

/**
  离线信息发送成功后 更新 small或large 状态  列表
*/
func UpdateOffMessageList(mbeans []*MBean, status int) {
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
	/**
	message := dao.NewTim_message()
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
	message.Update()*/
	my_message := new(hbase.Tim_message)
	if toname < fromname {
		my_message.Small = fmt.Sprint(status)
	} else {
		my_message.Large = fmt.Sprint(status)
	}
	rows := make([]int64, 0)
	for _, mbean := range mbeans {
		rows = append(rows, utils.Atoi64(mbean.GetMid()))
	}
	my_message.Updates(rows)
}

/***/
func LoadMBean(fidname, tidname, domain string, fromstamp, tostamp *string, limitcount int32) (tms []*MBean) {
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
	/**
	timMessage := dao.NewTim_message()
	wheres := make([]*basedao.Where, 0)
	if fromstamp != nil && tostamp != nil {
		wheres = append(wheres, timMessage.Stamp.Between(*fromstamp, *tostamp))
	} else if fromstamp != nil {
		wheres = append(wheres, timMessage.Stamp.GT(*fromstamp))
	} else if tostamp != nil {
		wheres = append(wheres, timMessage.Stamp.LT(*tostamp))
	}
	wheres = append(wheres, timMessage.Chatid.EQ(chatid))
	if isLarge {
		wheres = append(wheres, timMessage.Large.EQ(1))
	} else {
		wheres = append(wheres, timMessage.Small.EQ(1))
	}
	timMessage.Where(wheres...)
	timMessage.OrderBy(timMessage.Id.Desc())
	if limitcount > 0 {
		timMessage.Limit(0, limitcount)
	}
	timMessages, err := timMessage.Selects()
	if err == nil && timMessages != nil {
		tms = make([]*MBean, 0)
		for _, msg := range timMessages {
			tm := new(MBean)
			bb, er := base64Util.Base64Decode(msg.GetStanza())
			if er == nil {
				thrift.NewTDeserializer().Read(tm, bb)
				mid := fmt.Sprint(msg.GetId())
				tm.Mid = &mid
				tms = append(tms, tm)
			} else {
				logger.Error("Base64Decode:", er)
			}
		}
	}**/
	my_Message := new(hbase.Tim_message)
	beans := make([]*hbase.Bean, 0)
	if isLarge {
		b := new(hbase.Bean)
		b.Family = "large"
		b.Value = "1"
		beans = append(beans, b)
	} else {
		b := new(hbase.Bean)
		b.Family = "small"
		b.Value = "1"
		beans = append(beans, b)
	}
	b := new(hbase.Bean)
	b.Family = "index"
	b.Qualifier = chatid
	beans = append(beans, b)
	//	rs, er := hbase.Scans(my_Message.Tablename(), beans, 0, true)
	rs, er := hbase.ScansFromRow(my_Message.Tablename(), beans, 0, true)
	if er != nil {
		return
	}

	tms = make([]*MBean, 0)
	my_Messages := make([]*hbase.Tim_message, 0)
	for _, r := range rs {
		o := new(hbase.Tim_message)
		hbase.Result2object(r, o)
		my_Messages = append(my_Messages, o)
	}
	for _, msg := range my_Messages {
		tm := new(MBean)
		bb, er := utils.Base64Decode(msg.Stanza)
		if er == nil {
			thrift.NewTDeserializer().Read(tm, bb)
			mid := fmt.Sprint(msg.Id)
			tm.Mid = &mid
			tms = append(tms, tm)
		} else {
			logger.Error("Base64Decode:", er)
		}
	}
	return
}

func DelMBean(fidname, tidname, domain, mid string) {
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
	//	timMessage := dao.NewTim_message()
	//	if isLarge {
	//		timMessage.SetLarge(0)
	//	} else {
	//		timMessage.SetSmall(0)
	//	}
	//	timMessage.Where(timMessage.Chatid.EQ(chatid), timMessage.Id.EQ(mid))
	//	timMessage.Update()
	timMessage := new(hbase.Tim_message)
	hbase.Select(timMessage.Tablename(), utils.Atoi64(mid), "", "", timMessage)
	if timMessage.Chatid == chatid {
		tm := new(hbase.Tim_message)
		if isLarge {
			tm.Large = "0"
		} else {
			tm.Small = "0"
		}
		tm.Update(utils.Atoi64(mid))
	}
}

func DelAllMBean(fidname, tidname, domain string) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	chatid := utils.Chatid(fidname, tidname, domain)
	timMessage := new(hbase.Tim_message)
	rs, er := hbase.Selects(timMessage.Tablename(), "index", chatid, 0, false)
	if er == nil {
		rows := make([]int64, 0)
		for _, r := range rs {
			row := hbase.Bytes2hex(r.GetRow())
			rows = append(rows, row)
		}
		isLarge := fidname > tidname
		tm := new(hbase.Tim_message)
		if isLarge {
			tm.Large = "0"
		} else {
			tm.Small = "0"
		}
		tm.Updates(rows)
	}

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

/**
func Auth(tid *Tid, pwd string) (b bool) {
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
}**/

/**
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
**/
/**
func _auth(tid *Tid, pwd string) (b bool) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(string(debug.Stack()))
			}
		}()
		loginname, _ := connect.GetLoginName(tid)
		my_user := dao.NewTim_user()
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
}*/

/**
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
	my_domain := dao.NewTim_domain()
	my_domain.Where(my_domain.Domain.EQ(domain))
	var err error
	my_domain, err = my_domain.Select()
	if err == nil && my_domain != nil && my_domain.GetId() > 0 {
		domainmap.Put(domain, time.Now().UnixNano()/1000000000)
		return true
	}
	return false
}*/

/***
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
}***/

/***
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
	logger.Debug(domain, " ", fromname)
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
				domain := fromtid.GetDomain()
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
***/
/****
func updateVersion() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error(string(debug.Stack()))
		}
	}()
	if CF.Db_Exsit == 0 {
		return
	}
	timDomain := dao.New_config()
	timDomain.Where(timDomain.Keyword.EQ("version"))
	td, err := timDomain.Select()
	if err == nil && td != nil && td.GetId() > 0 {
		timDomain = dao.New_config()
		timDomain.SetValuestr(fmt.Sprint(VersionCode))
		timDomain.SetRemark(fmt.Sprint(VersionName, " | ", VersionCode, " | ", utils.Nowe()))
		timDomain.Where(timDomain.Id.EQ(td.GetId()))
		timDomain.Update()
	} else {
		timDomain = dao.New_config()
		timDomain.SetValuestr(fmt.Sprint(VersionCode))
		timDomain.SetRemark(fmt.Sprint(VersionName, " | ", VersionCode, " | ", utils.Nowe()))
		timDomain.SetCreatetime(utils.NowTime())
		timDomain.SetKeyword("version")
		timDomain.Insert()
	}
}
**/
