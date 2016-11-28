package dao

/**
tablename:my_mucmember
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_mucmember_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucmember_Id) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Id) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Roomtid struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmember_Roomtid) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Roomtid) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Tidname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmember_Tidname) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Tidname) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Type struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucmember_Type) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Type) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Affiliation struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucmember_Affiliation) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Affiliation) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Domain struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmember_Domain) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Domain) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Nickname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmember_Nickname) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Nickname) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Updatetime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmember_Updatetime) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Updatetime) Value() interface{} {
	return c.FieldValue
}

type my_mucmember_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucmember_Createtime) Name() string {
	return c.fieldName
}

func (c *my_mucmember_Createtime) Value() interface{} {
	return c.FieldValue
}

type Mucmember struct {
	basedao.Table
	Createtime  *my_mucmember_Createtime
	Domain      *my_mucmember_Domain
	Nickname    *my_mucmember_Nickname
	Updatetime  *my_mucmember_Updatetime
	Type        *my_mucmember_Type
	Affiliation *my_mucmember_Affiliation
	Id          *my_mucmember_Id
	Roomtid     *my_mucmember_Roomtid
	Tidname     *my_mucmember_Tidname
}

func (u *Mucmember) GetDomain() string {
	return *u.Domain.FieldValue
}

func (u *Mucmember) SetDomain(arg string) {
	u.Table.ModifyMap[u.Domain.fieldName] = arg
	v := string(arg)
	u.Domain.FieldValue = &v
}

func (u *Mucmember) GetNickname() string {
	return *u.Nickname.FieldValue
}

func (u *Mucmember) SetNickname(arg string) {
	u.Table.ModifyMap[u.Nickname.fieldName] = arg
	v := string(arg)
	u.Nickname.FieldValue = &v
}

func (u *Mucmember) GetUpdatetime() string {
	return *u.Updatetime.FieldValue
}

func (u *Mucmember) SetUpdatetime(arg string) {
	u.Table.ModifyMap[u.Updatetime.fieldName] = arg
	v := string(arg)
	u.Updatetime.FieldValue = &v
}

func (u *Mucmember) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Mucmember) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Mucmember) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Mucmember) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Mucmember) GetRoomtid() string {
	return *u.Roomtid.FieldValue
}

func (u *Mucmember) SetRoomtid(arg string) {
	u.Table.ModifyMap[u.Roomtid.fieldName] = arg
	v := string(arg)
	u.Roomtid.FieldValue = &v
}

func (u *Mucmember) GetTidname() string {
	return *u.Tidname.FieldValue
}

func (u *Mucmember) SetTidname(arg string) {
	u.Table.ModifyMap[u.Tidname.fieldName] = arg
	v := string(arg)
	u.Tidname.FieldValue = &v
}

func (u *Mucmember) GetType() int32 {
	return *u.Type.FieldValue
}

func (u *Mucmember) SetType(arg int64) {
	u.Table.ModifyMap[u.Type.fieldName] = arg
	v := int32(arg)
	u.Type.FieldValue = &v
}

func (u *Mucmember) GetAffiliation() int32 {
	return *u.Affiliation.FieldValue
}

func (u *Mucmember) SetAffiliation(arg int64) {
	u.Table.ModifyMap[u.Affiliation.fieldName] = arg
	v := int32(arg)
	u.Affiliation.FieldValue = &v
}

func (t *Mucmember) Query(columns ...basedao.Column) ([]Mucmember, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Nickname, t.Updatetime, t.Createtime, t.Id, t.Roomtid, t.Tidname, t.Type, t.Affiliation}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Mucmember, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_mucmember()
		go copyTim_mucmember(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_mucmember(channle chan int16, rows []interface{}, t *Mucmember, columns []basedao.Column) {
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

func (t *Mucmember) QuerySingle(columns ...basedao.Column) (*Mucmember, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Nickname, t.Updatetime, t.Createtime, t.Id, t.Roomtid, t.Tidname, t.Type, t.Affiliation}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_mucmember()
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

func (t *Mucmember) Select(columns ...basedao.Column) (*Mucmember, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Nickname, t.Updatetime, t.Createtime, t.Id, t.Roomtid, t.Tidname, t.Type, t.Affiliation}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_mucmember()
		cpTim_mucmember(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Mucmember) Selects(columns ...basedao.Column) ([]*Mucmember, error) {
	if columns == nil {
		columns = []basedao.Column{t.Domain, t.Nickname, t.Updatetime, t.Createtime, t.Id, t.Roomtid, t.Tidname, t.Type, t.Affiliation}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Mucmember, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_mucmember()
		cpTim_mucmember(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_mucmember(buff []interface{}, t *Mucmember, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "id":
			buff[i] = &t.Id.FieldValue
		case "roomtid":
			buff[i] = &t.Roomtid.FieldValue
		case "tidname":
			buff[i] = &t.Tidname.FieldValue
		case "type":
			buff[i] = &t.Type.FieldValue
		case "affiliation":
			buff[i] = &t.Affiliation.FieldValue
		case "domain":
			buff[i] = &t.Domain.FieldValue
		case "nickname":
			buff[i] = &t.Nickname.FieldValue
		case "updatetime":
			buff[i] = &t.Updatetime.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		}
	}
}

func New_mucmember(tableName ...string) *Mucmember {
	domain := &my_mucmember_Domain{fieldName: "domain"}
	domain.Field.FieldName = "domain"
	nickname := &my_mucmember_Nickname{fieldName: "nickname"}
	nickname.Field.FieldName = "nickname"
	updatetime := &my_mucmember_Updatetime{fieldName: "updatetime"}
	updatetime.Field.FieldName = "updatetime"
	createtime := &my_mucmember_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	id := &my_mucmember_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	roomtid := &my_mucmember_Roomtid{fieldName: "roomtid"}
	roomtid.Field.FieldName = "roomtid"
	tidname := &my_mucmember_Tidname{fieldName: "tidname"}
	tidname.Field.FieldName = "tidname"
	type_ := &my_mucmember_Type{fieldName: "type"}
	type_.Field.FieldName = "type"
	affiliation := &my_mucmember_Affiliation{fieldName: "affiliation"}
	affiliation.Field.FieldName = "affiliation"
	table := &Mucmember{Id: id, Roomtid: roomtid, Tidname: tidname, Type: type_, Affiliation: affiliation, Domain: domain, Nickname: nickname, Updatetime: updatetime, Createtime: createtime}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_mucmember"
	}
	return table
}
