package dao

/**
tablename:my_config
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_config_Valuestr struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_config_Valuestr) Name() string {
	return c.fieldName
}

func (c *my_config_Valuestr) Value() interface{} {
	return c.FieldValue
}

type my_config_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_config_Createtime) Name() string {
	return c.fieldName
}

func (c *my_config_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_config_Remark struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_config_Remark) Name() string {
	return c.fieldName
}

func (c *my_config_Remark) Value() interface{} {
	return c.FieldValue
}

type my_config_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_config_Id) Name() string {
	return c.fieldName
}

func (c *my_config_Id) Value() interface{} {
	return c.FieldValue
}

type my_config_Keyword struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_config_Keyword) Name() string {
	return c.fieldName
}

func (c *my_config_Keyword) Value() interface{} {
	return c.FieldValue
}

type Config struct {
	basedao.Table
	Id         *my_config_Id
	Keyword    *my_config_Keyword
	Valuestr   *my_config_Valuestr
	Createtime *my_config_Createtime
	Remark     *my_config_Remark
}

func (u *Config) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Config) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Config) GetKeyword() string {
	return *u.Keyword.FieldValue
}

func (u *Config) SetKeyword(arg string) {
	u.Table.ModifyMap[u.Keyword.fieldName] = arg
	v := string(arg)
	u.Keyword.FieldValue = &v
}

func (u *Config) GetValuestr() string {
	return *u.Valuestr.FieldValue
}

func (u *Config) SetValuestr(arg string) {
	u.Table.ModifyMap[u.Valuestr.fieldName] = arg
	v := string(arg)
	u.Valuestr.FieldValue = &v
}

func (u *Config) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Config) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Config) GetRemark() string {
	return *u.Remark.FieldValue
}

func (u *Config) SetRemark(arg string) {
	u.Table.ModifyMap[u.Remark.fieldName] = arg
	v := string(arg)
	u.Remark.FieldValue = &v
}

func (t *Config) Query(columns ...basedao.Column) ([]Config, error) {
	if columns == nil {
		columns = []basedao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Config, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_config()
		go copy_config(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copy_config(channle chan int16, rows []interface{}, t *Config, columns []basedao.Column) {
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

func (t *Config) QuerySingle(columns ...basedao.Column) (*Config, error) {
	if columns == nil {
		columns = []basedao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_config()
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

func (t *Config) Select(columns ...basedao.Column) (*Config, error) {
	if columns == nil {
		columns = []basedao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_config()
		cp_config(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Config) Selects(columns ...basedao.Column) ([]*Config, error) {
	if columns == nil {
		columns = []basedao.Column{t.Createtime, t.Remark, t.Id, t.Keyword, t.Valuestr}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Config, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_config()
		cp_config(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cp_config(buff []interface{}, t *Config, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "id":
			buff[i] = &t.Id.FieldValue
		case "keyword":
			buff[i] = &t.Keyword.FieldValue
		case "valuestr":
			buff[i] = &t.Valuestr.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "remark":
			buff[i] = &t.Remark.FieldValue
		}
	}
}

func New_config(tableName ...string) *Config {
	createtime := &my_config_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	remark := &my_config_Remark{fieldName: "remark"}
	remark.Field.FieldName = "remark"
	id := &my_config_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	keyword := &my_config_Keyword{fieldName: "keyword"}
	keyword.Field.FieldName = "keyword"
	valuestr := &my_config_Valuestr{fieldName: "valuestr"}
	valuestr.Field.FieldName = "valuestr"
	table := &Config{Keyword: keyword, Valuestr: valuestr, Createtime: createtime, Remark: remark, Id: id}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_config"
	}
	return table
}
