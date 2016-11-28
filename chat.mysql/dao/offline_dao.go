package dao

/**
tablename:my_offline
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_offline_Domain struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Domain) Name() string {
	return c.fieldName
}

func (c *my_offline_Domain) Value() interface{} {
	return c.FieldValue
}

type my_offline_Username struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Username) Name() string {
	return c.fieldName
}

func (c *my_offline_Username) Value() interface{} {
	return c.FieldValue
}

type my_offline_Stamp struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Stamp) Name() string {
	return c.fieldName
}

func (c *my_offline_Stamp) Value() interface{} {
	return c.FieldValue
}

type my_offline_Fromuser struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Fromuser) Name() string {
	return c.fieldName
}

func (c *my_offline_Fromuser) Value() interface{} {
	return c.FieldValue
}

type my_offline_Gname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Gname) Name() string {
	return c.fieldName
}

func (c *my_offline_Gname) Value() interface{} {
	return c.FieldValue
}

type my_offline_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Createtime) Name() string {
	return c.fieldName
}

func (c *my_offline_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_offline_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_offline_Id) Name() string {
	return c.fieldName
}

func (c *my_offline_Id) Value() interface{} {
	return c.FieldValue
}

type my_offline_Mid struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_offline_Mid) Name() string {
	return c.fieldName
}

func (c *my_offline_Mid) Value() interface{} {
	return c.FieldValue
}

type my_offline_Msgtype struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_offline_Msgtype) Name() string {
	return c.fieldName
}

func (c *my_offline_Msgtype) Value() interface{} {
	return c.FieldValue
}

type my_offline_Msgmode struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_offline_Msgmode) Name() string {
	return c.fieldName
}

func (c *my_offline_Msgmode) Value() interface{} {
	return c.FieldValue
}

type my_offline_Message_size struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_offline_Message_size) Name() string {
	return c.fieldName
}

func (c *my_offline_Message_size) Value() interface{} {
	return c.FieldValue
}

type my_offline_Stanza struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_offline_Stanza) Name() string {
	return c.fieldName
}

func (c *my_offline_Stanza) Value() interface{} {
	return c.FieldValue
}

type Offline struct {
	basedao.Table
	Domain       *my_offline_Domain
	Username     *my_offline_Username
	Stamp        *my_offline_Stamp
	Fromuser     *my_offline_Fromuser
	Gname        *my_offline_Gname
	Createtime   *my_offline_Createtime
	Id           *my_offline_Id
	Mid          *my_offline_Mid
	Msgtype      *my_offline_Msgtype
	Msgmode      *my_offline_Msgmode
	Message_size *my_offline_Message_size
	Stanza       *my_offline_Stanza
}

func (u *Offline) GetStanza() string {
	return *u.Stanza.FieldValue
}

func (u *Offline) SetStanza(arg string) {
	u.Table.ModifyMap[u.Stanza.fieldName] = arg
	v := string(arg)
	u.Stanza.FieldValue = &v
}

func (u *Offline) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Offline) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Offline) GetMid() int32 {
	return *u.Mid.FieldValue
}

func (u *Offline) SetMid(arg int64) {
	u.Table.ModifyMap[u.Mid.fieldName] = arg
	v := int32(arg)
	u.Mid.FieldValue = &v
}

func (u *Offline) GetMsgtype() int32 {
	return *u.Msgtype.FieldValue
}

func (u *Offline) SetMsgtype(arg int64) {
	u.Table.ModifyMap[u.Msgtype.fieldName] = arg
	v := int32(arg)
	u.Msgtype.FieldValue = &v
}

func (u *Offline) GetMsgmode() int32 {
	return *u.Msgmode.FieldValue
}

func (u *Offline) SetMsgmode(arg int64) {
	u.Table.ModifyMap[u.Msgmode.fieldName] = arg
	v := int32(arg)
	u.Msgmode.FieldValue = &v
}

func (u *Offline) GetMessage_size() int32 {
	return *u.Message_size.FieldValue
}

func (u *Offline) SetMessage_size(arg int64) {
	u.Table.ModifyMap[u.Message_size.fieldName] = arg
	v := int32(arg)
	u.Message_size.FieldValue = &v
}

func (u *Offline) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Offline) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Offline) GetDomain() string {
	return *u.Domain.FieldValue
}

func (u *Offline) SetDomain(arg string) {
	u.Table.ModifyMap[u.Domain.fieldName] = arg
	v := string(arg)
	u.Domain.FieldValue = &v
}

func (u *Offline) GetUsername() string {
	return *u.Username.FieldValue
}

func (u *Offline) SetUsername(arg string) {
	u.Table.ModifyMap[u.Username.fieldName] = arg
	v := string(arg)
	u.Username.FieldValue = &v
}

func (u *Offline) GetStamp() string {
	return *u.Stamp.FieldValue
}

func (u *Offline) SetStamp(arg string) {
	u.Table.ModifyMap[u.Stamp.fieldName] = arg
	v := string(arg)
	u.Stamp.FieldValue = &v
}

func (u *Offline) GetFromuser() string {
	return *u.Fromuser.FieldValue
}

func (u *Offline) SetFromuser(arg string) {
	u.Table.ModifyMap[u.Fromuser.fieldName] = arg
	v := string(arg)
	u.Fromuser.FieldValue = &v
}

func (u *Offline) GetGname() string {
	return *u.Gname.FieldValue
}

func (u *Offline) SetGname(arg string) {
	u.Table.ModifyMap[u.Gname.fieldName] = arg
	v := string(arg)
	u.Gname.FieldValue = &v
}

func (t *Offline) Query(columns ...basedao.Column) ([]Offline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Message_size, t.Stanza, t.Id, t.Mid, t.Msgtype, t.Msgmode, t.Gname, t.Createtime, t.Domain, t.Username, t.Stamp, t.Fromuser}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Offline, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_offline()
		go copyTim_offline(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_offline(channle chan int16, rows []interface{}, t *Offline, columns []basedao.Column) {
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

func (t *Offline) QuerySingle(columns ...basedao.Column) (*Offline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Message_size, t.Stanza, t.Id, t.Mid, t.Msgtype, t.Msgmode, t.Gname, t.Createtime, t.Domain, t.Username, t.Stamp, t.Fromuser}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_offline()
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

func (t *Offline) Select(columns ...basedao.Column) (*Offline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Message_size, t.Stanza, t.Id, t.Mid, t.Msgtype, t.Msgmode, t.Gname, t.Createtime, t.Domain, t.Username, t.Stamp, t.Fromuser}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_offline()
		cpTim_offline(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Offline) Selects(columns ...basedao.Column) ([]*Offline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Message_size, t.Stanza, t.Id, t.Mid, t.Msgtype, t.Msgmode, t.Gname, t.Createtime, t.Domain, t.Username, t.Stamp, t.Fromuser}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Offline, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_offline()
		cpTim_offline(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_offline(buff []interface{}, t *Offline, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "msgtype":
			buff[i] = &t.Msgtype.FieldValue
		case "msgmode":
			buff[i] = &t.Msgmode.FieldValue
		case "message_size":
			buff[i] = &t.Message_size.FieldValue
		case "stanza":
			buff[i] = &t.Stanza.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		case "mid":
			buff[i] = &t.Mid.FieldValue
		case "stamp":
			buff[i] = &t.Stamp.FieldValue
		case "fromuser":
			buff[i] = &t.Fromuser.FieldValue
		case "gname":
			buff[i] = &t.Gname.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "domain":
			buff[i] = &t.Domain.FieldValue
		case "username":
			buff[i] = &t.Username.FieldValue
		}
	}
}

func New_offline(tableName ...string) *Offline {
	message_size := &my_offline_Message_size{fieldName: "message_size"}
	message_size.Field.FieldName = "message_size"
	stanza := &my_offline_Stanza{fieldName: "stanza"}
	stanza.Field.FieldName = "stanza"
	id := &my_offline_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	mid := &my_offline_Mid{fieldName: "mid"}
	mid.Field.FieldName = "mid"
	msgtype_ := &my_offline_Msgtype{fieldName: "msgtype"}
	msgtype_.Field.FieldName = "msgtype"
	msgmode := &my_offline_Msgmode{fieldName: "msgmode"}
	msgmode.Field.FieldName = "msgmode"
	gname := &my_offline_Gname{fieldName: "gname"}
	gname.Field.FieldName = "gname"
	createtime := &my_offline_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	domain := &my_offline_Domain{fieldName: "domain"}
	domain.Field.FieldName = "domain"
	username := &my_offline_Username{fieldName: "username"}
	username.Field.FieldName = "username"
	stamp := &my_offline_Stamp{fieldName: "stamp"}
	stamp.Field.FieldName = "stamp"
	fromuser := &my_offline_Fromuser{fieldName: "fromuser"}
	fromuser.Field.FieldName = "fromuser"
	table := &Offline{Stamp: stamp, Fromuser: fromuser, Gname: gname, Createtime: createtime, Domain: domain, Username: username, Msgtype: msgtype_, Msgmode: msgmode, Message_size: message_size, Stanza: stanza, Id: id, Mid: mid}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_offline"
	}
	return table
}
