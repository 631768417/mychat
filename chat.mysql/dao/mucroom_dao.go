package dao

/**
tablename:my_mucroom
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_mucroom_Description struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Description) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Description) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Updatetime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Updatetime) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Updatetime) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Domain struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Domain) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Domain) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Maxusers struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucroom_Maxusers) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Maxusers) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Theme struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Theme) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Theme) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Name struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Name) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Name) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Password struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Password) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Password) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Createtime) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_mucroom_Id) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Id) Value() interface{} {
	return c.FieldValue
}

type my_mucroom_Roomtid struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_mucroom_Roomtid) Name() string {
	return c.fieldName
}

func (c *my_mucroom_Roomtid) Value() interface{} {
	return c.FieldValue
}

type Mucroom struct {
	basedao.Table
	Id          *my_mucroom_Id
	Roomtid     *my_mucroom_Roomtid
	Theme       *my_mucroom_Theme
	Name        *my_mucroom_Name
	Password    *my_mucroom_Password
	Createtime  *my_mucroom_Createtime
	Domain      *my_mucroom_Domain
	Maxusers    *my_mucroom_Maxusers
	Description *my_mucroom_Description
	Updatetime  *my_mucroom_Updatetime
}

func (u *Mucroom) GetPassword() string {
	return *u.Password.FieldValue
}

func (u *Mucroom) SetPassword(arg string) {
	u.Table.ModifyMap[u.Password.fieldName] = arg
	v := string(arg)
	u.Password.FieldValue = &v
}

func (u *Mucroom) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Mucroom) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Mucroom) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Mucroom) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Mucroom) GetRoomtid() string {
	return *u.Roomtid.FieldValue
}

func (u *Mucroom) SetRoomtid(arg string) {
	u.Table.ModifyMap[u.Roomtid.fieldName] = arg
	v := string(arg)
	u.Roomtid.FieldValue = &v
}

func (u *Mucroom) GetTheme() string {
	return *u.Theme.FieldValue
}

func (u *Mucroom) SetTheme(arg string) {
	u.Table.ModifyMap[u.Theme.fieldName] = arg
	v := string(arg)
	u.Theme.FieldValue = &v
}

func (u *Mucroom) GetName() string {
	return *u.Name.FieldValue
}

func (u *Mucroom) SetName(arg string) {
	u.Table.ModifyMap[u.Name.fieldName] = arg
	v := string(arg)
	u.Name.FieldValue = &v
}

func (u *Mucroom) GetDomain() string {
	return *u.Domain.FieldValue
}

func (u *Mucroom) SetDomain(arg string) {
	u.Table.ModifyMap[u.Domain.fieldName] = arg
	v := string(arg)
	u.Domain.FieldValue = &v
}

func (u *Mucroom) GetMaxusers() int32 {
	return *u.Maxusers.FieldValue
}

func (u *Mucroom) SetMaxusers(arg int64) {
	u.Table.ModifyMap[u.Maxusers.fieldName] = arg
	v := int32(arg)
	u.Maxusers.FieldValue = &v
}

func (u *Mucroom) GetDescription() string {
	return *u.Description.FieldValue
}

func (u *Mucroom) SetDescription(arg string) {
	u.Table.ModifyMap[u.Description.fieldName] = arg
	v := string(arg)
	u.Description.FieldValue = &v
}

func (u *Mucroom) GetUpdatetime() string {
	return *u.Updatetime.FieldValue
}

func (u *Mucroom) SetUpdatetime(arg string) {
	u.Table.ModifyMap[u.Updatetime.fieldName] = arg
	v := string(arg)
	u.Updatetime.FieldValue = &v
}

func (t *Mucroom) Query(columns ...basedao.Column) ([]Mucroom, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Roomtid, t.Theme, t.Name, t.Password, t.Createtime, t.Domain, t.Maxusers, t.Description, t.Updatetime}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Mucroom, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_mucroom()
		go copyTim_mucroom(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_mucroom(channle chan int16, rows []interface{}, t *Mucroom, columns []basedao.Column) {
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

func (t *Mucroom) QuerySingle(columns ...basedao.Column) (*Mucroom, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Roomtid, t.Theme, t.Name, t.Password, t.Createtime, t.Domain, t.Maxusers, t.Description, t.Updatetime}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_mucroom()
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

func (t *Mucroom) Select(columns ...basedao.Column) (*Mucroom, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Roomtid, t.Theme, t.Name, t.Password, t.Createtime, t.Domain, t.Maxusers, t.Description, t.Updatetime}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_mucroom()
		cpTim_mucroom(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Mucroom) Selects(columns ...basedao.Column) ([]*Mucroom, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Roomtid, t.Theme, t.Name, t.Password, t.Createtime, t.Domain, t.Maxusers, t.Description, t.Updatetime}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Mucroom, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_mucroom()
		cpTim_mucroom(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_mucroom(buff []interface{}, t *Mucroom, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "maxusers":
			buff[i] = &t.Maxusers.FieldValue
		case "description":
			buff[i] = &t.Description.FieldValue
		case "updatetime":
			buff[i] = &t.Updatetime.FieldValue
		case "domain":
			buff[i] = &t.Domain.FieldValue
		case "roomtid":
			buff[i] = &t.Roomtid.FieldValue
		case "theme":
			buff[i] = &t.Theme.FieldValue
		case "name":
			buff[i] = &t.Name.FieldValue
		case "password":
			buff[i] = &t.Password.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		}
	}
}

func New_mucroom(tableName ...string) *Mucroom {
	name := &my_mucroom_Name{fieldName: "name"}
	name.Field.FieldName = "name"
	password := &my_mucroom_Password{fieldName: "password"}
	password.Field.FieldName = "password"
	createtime := &my_mucroom_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	id := &my_mucroom_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	roomtid := &my_mucroom_Roomtid{fieldName: "roomtid"}
	roomtid.Field.FieldName = "roomtid"
	theme := &my_mucroom_Theme{fieldName: "theme"}
	theme.Field.FieldName = "theme"
	updatetime := &my_mucroom_Updatetime{fieldName: "updatetime"}
	updatetime.Field.FieldName = "updatetime"
	domain := &my_mucroom_Domain{fieldName: "domain"}
	domain.Field.FieldName = "domain"
	maxusers := &my_mucroom_Maxusers{fieldName: "maxusers"}
	maxusers.Field.FieldName = "maxusers"
	description := &my_mucroom_Description{fieldName: "description"}
	description.Field.FieldName = "description"
	table := &Mucroom{Domain: domain, Maxusers: maxusers, Description: description, Updatetime: updatetime, Id: id, Roomtid: roomtid, Theme: theme, Name: name, Password: password, Createtime: createtime}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_mucroom"
	}
	return table
}
