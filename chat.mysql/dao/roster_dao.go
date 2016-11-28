package dao

/**
tablename:my_roster
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_roster_Loginname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_roster_Loginname) Name() string {
	return c.fieldName
}

func (c *my_roster_Loginname) Value() interface{} {
	return c.FieldValue
}

type my_roster_Username struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_roster_Username) Name() string {
	return c.fieldName
}

func (c *my_roster_Username) Value() interface{} {
	return c.FieldValue
}

type my_roster_Rostername struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_roster_Rostername) Name() string {
	return c.fieldName
}

func (c *my_roster_Rostername) Value() interface{} {
	return c.FieldValue
}

type my_roster_Rostertype struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_roster_Rostertype) Name() string {
	return c.fieldName
}

func (c *my_roster_Rostertype) Value() interface{} {
	return c.FieldValue
}

type my_roster_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_roster_Createtime) Name() string {
	return c.fieldName
}

func (c *my_roster_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_roster_Remarknick struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_roster_Remarknick) Name() string {
	return c.fieldName
}

func (c *my_roster_Remarknick) Value() interface{} {
	return c.FieldValue
}

type my_roster_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_roster_Id) Name() string {
	return c.fieldName
}

func (c *my_roster_Id) Value() interface{} {
	return c.FieldValue
}

type Roster struct {
	basedao.Table
	Username   *my_roster_Username
	Rostername *my_roster_Rostername
	Rostertype *my_roster_Rostertype
	Createtime *my_roster_Createtime
	Remarknick *my_roster_Remarknick
	Id         *my_roster_Id
	Loginname  *my_roster_Loginname
}

func (u *Roster) GetRostername() string {
	return *u.Rostername.FieldValue
}

func (u *Roster) SetRostername(arg string) {
	u.Table.ModifyMap[u.Rostername.fieldName] = arg
	v := string(arg)
	u.Rostername.FieldValue = &v
}

func (u *Roster) GetRostertype() string {
	return *u.Rostertype.FieldValue
}

func (u *Roster) SetRostertype(arg string) {
	u.Table.ModifyMap[u.Rostertype.fieldName] = arg
	v := string(arg)
	u.Rostertype.FieldValue = &v
}

func (u *Roster) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Roster) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Roster) GetRemarknick() string {
	return *u.Remarknick.FieldValue
}

func (u *Roster) SetRemarknick(arg string) {
	u.Table.ModifyMap[u.Remarknick.fieldName] = arg
	v := string(arg)
	u.Remarknick.FieldValue = &v
}

func (u *Roster) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Roster) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Roster) GetLoginname() string {
	return *u.Loginname.FieldValue
}

func (u *Roster) SetLoginname(arg string) {
	u.Table.ModifyMap[u.Loginname.fieldName] = arg
	v := string(arg)
	u.Loginname.FieldValue = &v
}

func (u *Roster) GetUsername() string {
	return *u.Username.FieldValue
}

func (u *Roster) SetUsername(arg string) {
	u.Table.ModifyMap[u.Username.fieldName] = arg
	v := string(arg)
	u.Username.FieldValue = &v
}

func (t *Roster) Query(columns ...basedao.Column) ([]Roster, error) {
	if columns == nil {
		columns = []basedao.Column{t.Loginname, t.Username, t.Rostername, t.Rostertype, t.Createtime, t.Remarknick, t.Id}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Roster, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_roster()
		go copyTim_roster(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_roster(channle chan int16, rows []interface{}, t *Roster, columns []basedao.Column) {
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

func (t *Roster) QuerySingle(columns ...basedao.Column) (*Roster, error) {
	if columns == nil {
		columns = []basedao.Column{t.Loginname, t.Username, t.Rostername, t.Rostertype, t.Createtime, t.Remarknick, t.Id}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_roster()
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

func (t *Roster) Select(columns ...basedao.Column) (*Roster, error) {
	if columns == nil {
		columns = []basedao.Column{t.Loginname, t.Username, t.Rostername, t.Rostertype, t.Createtime, t.Remarknick, t.Id}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_roster()
		cpTim_roster(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Roster) Selects(columns ...basedao.Column) ([]*Roster, error) {
	if columns == nil {
		columns = []basedao.Column{t.Loginname, t.Username, t.Rostername, t.Rostertype, t.Createtime, t.Remarknick, t.Id}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Roster, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_roster()
		cpTim_roster(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_roster(buff []interface{}, t *Roster, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "rostername":
			buff[i] = &t.Rostername.FieldValue
		case "rostertype":
			buff[i] = &t.Rostertype.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "remarknick":
			buff[i] = &t.Remarknick.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		case "loginname":
			buff[i] = &t.Loginname.FieldValue
		case "username":
			buff[i] = &t.Username.FieldValue
		}
	}
}

func New_roster(tableName ...string) *Roster {
	remarknick := &my_roster_Remarknick{fieldName: "remarknick"}
	remarknick.Field.FieldName = "remarknick"
	id := &my_roster_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	loginname := &my_roster_Loginname{fieldName: "loginname"}
	loginname.Field.FieldName = "loginname"
	username := &my_roster_Username{fieldName: "username"}
	username.Field.FieldName = "username"
	rostername := &my_roster_Rostername{fieldName: "rostername"}
	rostername.Field.FieldName = "rostername"
	rostertype_ := &my_roster_Rostertype{fieldName: "rostertype"}
	rostertype_.Field.FieldName = "rostertype"
	createtime := &my_roster_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	table := &Roster{Username: username, Rostername: rostername, Rostertype: rostertype_, Createtime: createtime, Remarknick: remarknick, Id: id, Loginname: loginname}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_roster"
	}
	return table
}
