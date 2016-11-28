package dao

/**
tablename:my_mucmessage
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_mucmessage_Stanza struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmessage_Stanza) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Stanza) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmessage_Createtime) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucmessage_Id) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Id) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Stamp struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmessage_Stamp) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Stamp) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Fromuser struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmessage_Fromuser) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Fromuser) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Roomtidname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmessage_Roomtidname) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Roomtidname) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Domain struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmessage_Domain) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Domain) Value() interface{} {
	return c.FieldValue
}

type my_mucmessage_Msgtype struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucmessage_Msgtype) Name() string {
	return c.fieldName
}

func (c *my_mucmessage_Msgtype) Value() interface{} {
	return c.FieldValue
}

type Mucmessage struct {
	basedao.Table
	Roomtidname *my_mucmessage_Roomtidname
	Domain      *my_mucmessage_Domain
	Msgtype     *my_mucmessage_Msgtype
	Stanza      *my_mucmessage_Stanza
	Createtime  *my_mucmessage_Createtime
	Id          *my_mucmessage_Id
	Stamp       *my_mucmessage_Stamp
	Fromuser    *my_mucmessage_Fromuser
}

func (u *Mucmessage) GetStanza() string {
	return *u.Stanza.FieldValue
}

func (u *Mucmessage) SetStanza(arg string) {
	u.Table.ModifyMap[u.Stanza.fieldName] = arg
	v := string(arg)
	u.Stanza.FieldValue = &v
}

func (u *Mucmessage) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Mucmessage) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Mucmessage) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Mucmessage) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Mucmessage) GetStamp() string {
	return *u.Stamp.FieldValue
}

func (u *Mucmessage) SetStamp(arg string) {
	u.Table.ModifyMap[u.Stamp.fieldName] = arg
	v := string(arg)
	u.Stamp.FieldValue = &v
}

func (u *Mucmessage) GetFromuser() string {
	return *u.Fromuser.FieldValue
}

func (u *Mucmessage) SetFromuser(arg string) {
	u.Table.ModifyMap[u.Fromuser.fieldName] = arg
	v := string(arg)
	u.Fromuser.FieldValue = &v
}

func (u *Mucmessage) GetRoomtidname() string {
	return *u.Roomtidname.FieldValue
}

func (u *Mucmessage) SetRoomtidname(arg string) {
	u.Table.ModifyMap[u.Roomtidname.fieldName] = arg
	v := string(arg)
	u.Roomtidname.FieldValue = &v
}

func (u *Mucmessage) GetDomain() string {
	return *u.Domain.FieldValue
}

func (u *Mucmessage) SetDomain(arg string) {
	u.Table.ModifyMap[u.Domain.fieldName] = arg
	v := string(arg)
	u.Domain.FieldValue = &v
}

func (u *Mucmessage) GetMsgtype() int32 {
	return *u.Msgtype.FieldValue
}

func (u *Mucmessage) SetMsgtype(arg int64) {
	u.Table.ModifyMap[u.Msgtype.fieldName] = arg
	v := int32(arg)
	u.Msgtype.FieldValue = &v
}

func (t *Mucmessage) Query(columns ...basedao.Column) ([]Mucmessage, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Msgtype, t.Stanza, t.Createtime, t.Id, t.Stamp, t.Fromuser, t.Roomtidname}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Mucmessage, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_mucmessage()
		go copyTim_mucmessage(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_mucmessage(channle chan int16, rows []interface{}, t *Mucmessage, columns []basedao.Column) {
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

func (t *Mucmessage) QuerySingle(columns ...basedao.Column) (*Mucmessage, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Msgtype, t.Stanza, t.Createtime, t.Id, t.Stamp, t.Fromuser, t.Roomtidname}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_mucmessage()
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

func (t *Mucmessage) Select(columns ...basedao.Column) (*Mucmessage, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Msgtype, t.Stanza, t.Createtime, t.Id, t.Stamp, t.Fromuser, t.Roomtidname}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_mucmessage()
		cpTim_mucmessage(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Mucmessage) Selects(columns ...basedao.Column) ([]*Mucmessage, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Msgtype, t.Stanza, t.Createtime, t.Id, t.Stamp, t.Fromuser, t.Roomtidname}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Mucmessage, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_mucmessage()
		cpTim_mucmessage(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_mucmessage(buff []interface{}, t *Mucmessage, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		case "stamp":
			buff[i] = &t.Stamp.FieldValue
		case "fromuser":
			buff[i] = &t.Fromuser.FieldValue
		case "roomtidname":
			buff[i] = &t.Roomtidname.FieldValue
		case "domain":
			buff[i] = &t.Domain.FieldValue
		case "msgtype":
			buff[i] = &t.Msgtype.FieldValue
		case "stanza":
			buff[i] = &t.Stanza.FieldValue
		}
	}
}

func New_mucmessage(tableName ...string) *Mucmessage {
	stamp := &my_mucmessage_Stamp{fieldName: "stamp"}
	stamp.Field.FieldName = "stamp"
	fromuser := &my_mucmessage_Fromuser{fieldName: "fromuser"}
	fromuser.Field.FieldName = "fromuser"
	roomtidname := &my_mucmessage_Roomtidname{fieldName: "roomtidname"}
	roomtidname.Field.FieldName = "roomtidname"
	domain := &my_mucmessage_Domain{fieldName: "domain"}
	domain.Field.FieldName = "domain"
	msgtype_ := &my_mucmessage_Msgtype{fieldName: "msgtype"}
	msgtype_.Field.FieldName = "msgtype"
	stanza := &my_mucmessage_Stanza{fieldName: "stanza"}
	stanza.Field.FieldName = "stanza"
	createtime := &my_mucmessage_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	id := &my_mucmessage_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	table := &Mucmessage{Id: id, Stamp: stamp, Fromuser: fromuser, Roomtidname: roomtidname, Domain: domain, Msgtype: msgtype_, Stanza: stanza, Createtime: createtime}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_mucmessage"
	}
	return table
}
