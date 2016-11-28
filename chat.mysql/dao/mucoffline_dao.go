package dao

/**
tablename:my_mucoffline
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_mucoffline_Msgtype struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucoffline_Msgtype) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Msgtype) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Message_size struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucoffline_Message_size) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Message_size) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucoffline_Createtime) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucoffline_Id) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Id) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Mid struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucoffline_Mid) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Mid) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Domain struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucoffline_Domain) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Domain) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Username struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucoffline_Username) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Username) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Stamp struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucoffline_Stamp) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Stamp) Value() interface{} {
	return c.FieldValue
}

type my_mucoffline_Roomid struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucoffline_Roomid) Name() string {
	return c.fieldName
}

func (c *my_mucoffline_Roomid) Value() interface{} {
	return c.FieldValue
}

type Mucoffline struct {
	basedao.Table
	Msgtype      *my_mucoffline_Msgtype
	Message_size *my_mucoffline_Message_size
	Domain       *my_mucoffline_Domain
	Username     *my_mucoffline_Username
	Stamp        *my_mucoffline_Stamp
	Roomid       *my_mucoffline_Roomid
	Createtime   *my_mucoffline_Createtime
	Id           *my_mucoffline_Id
	Mid          *my_mucoffline_Mid
}

func (u *Mucoffline) GetMsgtype() int32 {
	return *u.Msgtype.FieldValue
}

func (u *Mucoffline) SetMsgtype(arg int64) {
	u.Table.ModifyMap[u.Msgtype.fieldName] = arg
	v := int32(arg)
	u.Msgtype.FieldValue = &v
}

func (u *Mucoffline) GetMessage_size() int32 {
	return *u.Message_size.FieldValue
}

func (u *Mucoffline) SetMessage_size(arg int64) {
	u.Table.ModifyMap[u.Message_size.fieldName] = arg
	v := int32(arg)
	u.Message_size.FieldValue = &v
}

func (u *Mucoffline) GetDomain() string {
	return *u.Domain.FieldValue
}

func (u *Mucoffline) SetDomain(arg string) {
	u.Table.ModifyMap[u.Domain.fieldName] = arg
	v := string(arg)
	u.Domain.FieldValue = &v
}

func (u *Mucoffline) GetUsername() string {
	return *u.Username.FieldValue
}

func (u *Mucoffline) SetUsername(arg string) {
	u.Table.ModifyMap[u.Username.fieldName] = arg
	v := string(arg)
	u.Username.FieldValue = &v
}

func (u *Mucoffline) GetStamp() string {
	return *u.Stamp.FieldValue
}

func (u *Mucoffline) SetStamp(arg string) {
	u.Table.ModifyMap[u.Stamp.fieldName] = arg
	v := string(arg)
	u.Stamp.FieldValue = &v
}

func (u *Mucoffline) GetRoomid() string {
	return *u.Roomid.FieldValue
}

func (u *Mucoffline) SetRoomid(arg string) {
	u.Table.ModifyMap[u.Roomid.fieldName] = arg
	v := string(arg)
	u.Roomid.FieldValue = &v
}

func (u *Mucoffline) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Mucoffline) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Mucoffline) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Mucoffline) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Mucoffline) GetMid() int32 {
	return *u.Mid.FieldValue
}

func (u *Mucoffline) SetMid(arg int64) {
	u.Table.ModifyMap[u.Mid.fieldName] = arg
	v := int32(arg)
	u.Mid.FieldValue = &v
}

func (t *Mucoffline) Query(columns ...basedao.Column) ([]Mucoffline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Msgtype, t.Message_size, t.Stamp, t.Roomid, t.Createtime, t.Id, t.Mid, t.Domain, t.Username}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Mucoffline, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_mucoffline()
		go copyTim_mucoffline(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_mucoffline(channle chan int16, rows []interface{}, t *Mucoffline, columns []basedao.Column) {
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

func (t *Mucoffline) QuerySingle(columns ...basedao.Column) (*Mucoffline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Msgtype, t.Message_size, t.Stamp, t.Roomid, t.Createtime, t.Id, t.Mid, t.Domain, t.Username}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_mucoffline()
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

func (t *Mucoffline) Select(columns ...basedao.Column) (*Mucoffline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Msgtype, t.Message_size, t.Stamp, t.Roomid, t.Createtime, t.Id, t.Mid, t.Domain, t.Username}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_mucoffline()
		cpTim_mucoffline(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Mucoffline) Selects(columns ...basedao.Column) ([]*Mucoffline, error) {
	if columns == nil {
		columns = []basedao.Column{t.Msgtype, t.Message_size, t.Stamp, t.Roomid, t.Createtime, t.Id, t.Mid, t.Domain, t.Username}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Mucoffline, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_mucoffline()
		cpTim_mucoffline(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_mucoffline(buff []interface{}, t *Mucoffline, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "mid":
			buff[i] = &t.Mid.FieldValue
		case "domain":
			buff[i] = &t.Domain.FieldValue
		case "username":
			buff[i] = &t.Username.FieldValue
		case "stamp":
			buff[i] = &t.Stamp.FieldValue
		case "roomid":
			buff[i] = &t.Roomid.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		case "message_size":
			buff[i] = &t.Message_size.FieldValue
		case "msgtype":
			buff[i] = &t.Msgtype.FieldValue
		}
	}
}

func New_mucoffline(tableName ...string) *Mucoffline {
	msgtype_ := &my_mucoffline_Msgtype{fieldName: "msgtype"}
	msgtype_.Field.FieldName = "msgtype"
	message_size := &my_mucoffline_Message_size{fieldName: "message_size"}
	message_size.Field.FieldName = "message_size"
	domain := &my_mucoffline_Domain{fieldName: "domain"}
	domain.Field.FieldName = "domain"
	username := &my_mucoffline_Username{fieldName: "username"}
	username.Field.FieldName = "username"
	stamp := &my_mucoffline_Stamp{fieldName: "stamp"}
	stamp.Field.FieldName = "stamp"
	roomid := &my_mucoffline_Roomid{fieldName: "roomid"}
	roomid.Field.FieldName = "roomid"
	createtime := &my_mucoffline_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	id := &my_mucoffline_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	mid := &my_mucoffline_Mid{fieldName: "mid"}
	mid.Field.FieldName = "mid"
	table := &Mucoffline{Msgtype: msgtype_, Message_size: message_size, Roomid: roomid, Createtime: createtime, Id: id, Mid: mid, Domain: domain, Username: username, Stamp: stamp}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_mucoffline"
	}
	return table
}
