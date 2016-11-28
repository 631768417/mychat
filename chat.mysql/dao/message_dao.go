package dao

/**
tablename:my_message
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_message_Stanza struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Stanza) Name() string {
	return c.fieldName
}

func (c *my_message_Stanza) Value() interface{} {
	return c.FieldValue
}

type my_message_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Createtime) Name() string {
	return c.fieldName
}

func (c *my_message_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_message_Chatid struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Chatid) Name() string {
	return c.fieldName
}

func (c *my_message_Chatid) Value() interface{} {
	return c.FieldValue
}

type my_message_Fromuser struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Fromuser) Name() string {
	return c.fieldName
}

func (c *my_message_Fromuser) Value() interface{} {
	return c.FieldValue
}

type my_message_Msgtype struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_message_Msgtype) Name() string {
	return c.fieldName
}

func (c *my_message_Msgtype) Value() interface{} {
	return c.FieldValue
}

type my_message_Gname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Gname) Name() string {
	return c.fieldName
}

func (c *my_message_Gname) Value() interface{} {
	return c.FieldValue
}

type my_message_Small struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_message_Small) Name() string {
	return c.fieldName
}

func (c *my_message_Small) Value() interface{} {
	return c.FieldValue
}

type my_message_Large struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_message_Large) Name() string {
	return c.fieldName
}

func (c *my_message_Large) Value() interface{} {
	return c.FieldValue
}

type my_message_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_message_Id) Name() string {
	return c.fieldName
}

func (c *my_message_Id) Value() interface{} {
	return c.FieldValue
}

type my_message_Stamp struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Stamp) Name() string {
	return c.fieldName
}

func (c *my_message_Stamp) Value() interface{} {
	return c.FieldValue
}

type my_message_Touser struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_message_Touser) Name() string {
	return c.fieldName
}

func (c *my_message_Touser) Value() interface{} {
	return c.FieldValue
}

type my_message_Msgmode struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_message_Msgmode) Name() string {
	return c.fieldName
}

func (c *my_message_Msgmode) Value() interface{} {
	return c.FieldValue
}

type Message struct {
	basedao.Table
	Stamp      *my_message_Stamp
	Touser     *my_message_Touser
	Msgmode    *my_message_Msgmode
	Small      *my_message_Small
	Large      *my_message_Large
	Id         *my_message_Id
	Fromuser   *my_message_Fromuser
	Msgtype    *my_message_Msgtype
	Gname      *my_message_Gname
	Stanza     *my_message_Stanza
	Createtime *my_message_Createtime
	Chatid     *my_message_Chatid
}

func (u *Message) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Message) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Message) GetChatid() string {
	return *u.Chatid.FieldValue
}

func (u *Message) SetChatid(arg string) {
	u.Table.ModifyMap[u.Chatid.fieldName] = arg
	v := string(arg)
	u.Chatid.FieldValue = &v
}

func (u *Message) GetFromuser() string {
	return *u.Fromuser.FieldValue
}

func (u *Message) SetFromuser(arg string) {
	u.Table.ModifyMap[u.Fromuser.fieldName] = arg
	v := string(arg)
	u.Fromuser.FieldValue = &v
}

func (u *Message) GetMsgtype() int32 {
	return *u.Msgtype.FieldValue
}

func (u *Message) SetMsgtype(arg int64) {
	u.Table.ModifyMap[u.Msgtype.fieldName] = arg
	v := int32(arg)
	u.Msgtype.FieldValue = &v
}

func (u *Message) GetGname() string {
	return *u.Gname.FieldValue
}

func (u *Message) SetGname(arg string) {
	u.Table.ModifyMap[u.Gname.fieldName] = arg
	v := string(arg)
	u.Gname.FieldValue = &v
}

func (u *Message) GetStanza() string {
	return *u.Stanza.FieldValue
}

func (u *Message) SetStanza(arg string) {
	u.Table.ModifyMap[u.Stanza.fieldName] = arg
	v := string(arg)
	u.Stanza.FieldValue = &v
}

func (u *Message) GetLarge() int32 {
	return *u.Large.FieldValue
}

func (u *Message) SetLarge(arg int64) {
	u.Table.ModifyMap[u.Large.fieldName] = arg
	v := int32(arg)
	u.Large.FieldValue = &v
}

func (u *Message) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Message) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Message) GetStamp() string {
	return *u.Stamp.FieldValue
}

func (u *Message) SetStamp(arg string) {
	u.Table.ModifyMap[u.Stamp.fieldName] = arg
	v := string(arg)
	u.Stamp.FieldValue = &v
}

func (u *Message) GetTouser() string {
	return *u.Touser.FieldValue
}

func (u *Message) SetTouser(arg string) {
	u.Table.ModifyMap[u.Touser.fieldName] = arg
	v := string(arg)
	u.Touser.FieldValue = &v
}

func (u *Message) GetMsgmode() int32 {
	return *u.Msgmode.FieldValue
}

func (u *Message) SetMsgmode(arg int64) {
	u.Table.ModifyMap[u.Msgmode.fieldName] = arg
	v := int32(arg)
	u.Msgmode.FieldValue = &v
}

func (u *Message) GetSmall() int32 {
	return *u.Small.FieldValue
}

func (u *Message) SetSmall(arg int64) {
	u.Table.ModifyMap[u.Small.fieldName] = arg
	v := int32(arg)
	u.Small.FieldValue = &v
}

func (t *Message) Query(columns ...basedao.Column) ([]Message, error) {
	if columns == nil {
		columns = []basedao.Column{t.Stanza, t.Createtime, t.Chatid, t.Fromuser, t.Msgtype, t.Gname, t.Small, t.Large, t.Id, t.Stamp, t.Touser, t.Msgmode}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Message, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_message()
		go copyTim_message(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_message(channle chan int16, rows []interface{}, t *Message, columns []basedao.Column) {
	defer func() { channle <- 1 }()
	for j, core := range rows {
		if core == nil {
			continue
		}
		field := columns[j].Name()
		setfield := "Set" + basedao.ToUpperFirstLetter(field)
		reflect.ValueOf(t).MethodByName(setfield).Call([]reflect.Value{reflect.ValueOf(basedao.GetValue(&core))})
	}
}

func (t *Message) QuerySingle(columns ...basedao.Column) (*Message, error) {
	if columns == nil {
		columns = []basedao.Column{t.Stanza, t.Createtime, t.Chatid, t.Fromuser, t.Msgtype, t.Gname, t.Small, t.Large, t.Id, t.Stamp, t.Touser, t.Msgmode}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_message()
	for j, core := range rs {
		if core == nil {
			continue
		}
		field := columns[j].Name()
		setfield := "Set" + basedao.ToUpperFirstLetter(field)
		reflect.ValueOf(rt).MethodByName(setfield).Call([]reflect.Value{reflect.ValueOf(basedao.GetValue(&core))})
	}
	return rt, nil
}

func (t *Message) Select(columns ...basedao.Column) (*Message, error) {
	if columns == nil {
		columns = []basedao.Column{t.Stanza, t.Createtime, t.Chatid, t.Fromuser, t.Msgtype, t.Gname, t.Small, t.Large, t.Id, t.Stamp, t.Touser, t.Msgmode}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_message()
		cpTim_message(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Message) Selects(columns ...basedao.Column) ([]*Message, error) {
	if columns == nil {
		columns = []basedao.Column{t.Stanza, t.Createtime, t.Chatid, t.Fromuser, t.Msgtype, t.Gname, t.Small, t.Large, t.Id, t.Stamp, t.Touser, t.Msgmode}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Message, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_message()
		cpTim_message(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_message(buff []interface{}, t *Message, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "chatid":
			buff[i] = &t.Chatid.FieldValue
		case "fromuser":
			buff[i] = &t.Fromuser.FieldValue
		case "msgtype":
			buff[i] = &t.Msgtype.FieldValue
		case "gname":
			buff[i] = &t.Gname.FieldValue
		case "stanza":
			buff[i] = &t.Stanza.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		case "stamp":
			buff[i] = &t.Stamp.FieldValue
		case "touser":
			buff[i] = &t.Touser.FieldValue
		case "msgmode":
			buff[i] = &t.Msgmode.FieldValue
		case "small":
			buff[i] = &t.Small.FieldValue
		case "large":
			buff[i] = &t.Large.FieldValue
		}
	}
}

func New_message(tableName ...string) *Message {
	large := &my_message_Large{fieldName: "large"}
	large.Field.FieldName = "large"
	id := &my_message_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	stamp := &my_message_Stamp{fieldName: "stamp"}
	stamp.Field.FieldName = "stamp"
	touser := &my_message_Touser{fieldName: "touser"}
	touser.Field.FieldName = "touser"
	msgmode := &my_message_Msgmode{fieldName: "msgmode"}
	msgmode.Field.FieldName = "msgmode"
	small := &my_message_Small{fieldName: "small"}
	small.Field.FieldName = "small"
	createtime := &my_message_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	chatid := &my_message_Chatid{fieldName: "chatid"}
	chatid.Field.FieldName = "chatid"
	fromuser := &my_message_Fromuser{fieldName: "fromuser"}
	fromuser.Field.FieldName = "fromuser"
	msgtype_ := &my_message_Msgtype{fieldName: "msgtype"}
	msgtype_.Field.FieldName = "msgtype"
	gname := &my_message_Gname{fieldName: "gname"}
	gname.Field.FieldName = "gname"
	stanza := &my_message_Stanza{fieldName: "stanza"}
	stanza.Field.FieldName = "stanza"
	table := &Message{Gname: gname, Stanza: stanza, Createtime: createtime, Chatid: chatid, Fromuser: fromuser, Msgtype: msgtype_, Msgmode: msgmode, Small: small, Large: large, Id: id, Stamp: stamp, Touser: touser}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_message"
	}
	return table
}
