package dao

/**
tablename:my_property
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_property_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_property_Id) Name() string {
	return c.fieldName
}

func (c *my_property_Id) Value() interface{} {
	return c.FieldValue
}

type my_property_Keyword struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_property_Keyword) Name() string {
	return c.fieldName
}

func (c *my_property_Keyword) Value() interface{} {
	return c.FieldValue
}

type my_property_Valueint struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_property_Valueint) Name() string {
	return c.fieldName
}

func (c *my_property_Valueint) Value() interface{} {
	return c.FieldValue
}

type my_property_Valuestr struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_property_Valuestr) Name() string {
	return c.fieldName
}

func (c *my_property_Valuestr) Value() interface{} {
	return c.FieldValue
}

type my_property_Remark struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_property_Remark) Name() string {
	return c.fieldName
}

func (c *my_property_Remark) Value() interface{} {
	return c.FieldValue
}

type Property struct {
	basedao.Table
	Valueint *my_property_Valueint
	Valuestr *my_property_Valuestr
	Remark   *my_property_Remark
	Id       *my_property_Id
	Keyword  *my_property_Keyword
}

func (u *Property) GetValuestr() string {
	return *u.Valuestr.FieldValue
}

func (u *Property) SetValuestr(arg string) {
	u.Table.ModifyMap[u.Valuestr.fieldName] = arg
	v := string(arg)
	u.Valuestr.FieldValue = &v
}

func (u *Property) GetRemark() string {
	return *u.Remark.FieldValue
}

func (u *Property) SetRemark(arg string) {
	u.Table.ModifyMap[u.Remark.fieldName] = arg
	v := string(arg)
	u.Remark.FieldValue = &v
}

func (u *Property) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Property) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Property) GetKeyword() string {
	return *u.Keyword.FieldValue
}

func (u *Property) SetKeyword(arg string) {
	u.Table.ModifyMap[u.Keyword.fieldName] = arg
	v := string(arg)
	u.Keyword.FieldValue = &v
}

func (u *Property) GetValueint() int32 {
	return *u.Valueint.FieldValue
}

func (u *Property) SetValueint(arg int64) {
	u.Table.ModifyMap[u.Valueint.fieldName] = arg
	v := int32(arg)
	u.Valueint.FieldValue = &v
}

func (t *Property) Query(columns ...basedao.Column) ([]Property, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Property, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_property()
		go copyTim_property(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_property(channle chan int16, rows []interface{}, t *Property, columns []basedao.Column) {
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

func (t *Property) QuerySingle(columns ...basedao.Column) (*Property, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_property()
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

func (t *Property) Select(columns ...basedao.Column) (*Property, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_property()
		cpTim_property(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Property) Selects(columns ...basedao.Column) ([]*Property, error) {
	if columns == nil {
		columns = []basedao.Column{t.Id, t.Keyword, t.Valueint, t.Valuestr, t.Remark}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Property, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_property()
		cpTim_property(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_property(buff []interface{}, t *Property, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "keyword":
			buff[i] = &t.Keyword.FieldValue
		case "valueint":
			buff[i] = &t.Valueint.FieldValue
		case "valuestr":
			buff[i] = &t.Valuestr.FieldValue
		case "remark":
			buff[i] = &t.Remark.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		}
	}
}

func New_property(tableName ...string) *Property {
	valueint := &my_property_Valueint{fieldName: "valueint"}
	valueint.Field.FieldName = "valueint"
	valuestr := &my_property_Valuestr{fieldName: "valuestr"}
	valuestr.Field.FieldName = "valuestr"
	remark := &my_property_Remark{fieldName: "remark"}
	remark.Field.FieldName = "remark"
	id := &my_property_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	keyword := &my_property_Keyword{fieldName: "keyword"}
	keyword.Field.FieldName = "keyword"
	table := &Property{Id: id, Keyword: keyword, Valueint: valueint, Valuestr: valuestr, Remark: remark}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_property"
	}
	return table
}
