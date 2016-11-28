package dao

/**
tablename:my_domain
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_domain_Remark struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_domain_Remark) Name() string {
	return c.fieldName
}

func (c *my_domain_Remark) Value() interface{} {
	return c.FieldValue
}

type my_domain_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_domain_Id) Name() string {
	return c.fieldName
}

func (c *my_domain_Id) Value() interface{} {
	return c.FieldValue
}

type my_domain_Domain struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_domain_Domain) Name() string {
	return c.fieldName
}

func (c *my_domain_Domain) Value() interface{} {
	return c.FieldValue
}

type my_domain_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_domain_Createtime) Name() string {
	return c.fieldName
}

func (c *my_domain_Createtime) Value() interface{} {
	return c.FieldValue
}

type Domain struct {
	basedao.Table
	Domain     *my_domain_Domain
	Createtime *my_domain_Createtime
	Remark     *my_domain_Remark
	Id         *my_domain_Id
}

func (u *Domain) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *Domain) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *Domain) GetRemark() string {
	return *u.Remark.FieldValue
}

func (u *Domain) SetRemark(arg string) {
	u.Table.ModifyMap[u.Remark.fieldName] = arg
	v := string(arg)
	u.Remark.FieldValue = &v
}

func (u *Domain) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *Domain) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *Domain) GetDomain() string {
	return *u.Domain.FieldValue
}

func (u *Domain) SetDomain(arg string) {
	u.Table.ModifyMap[u.Domain.fieldName] = arg
	v := string(arg)
	u.Domain.FieldValue = &v
}

func (t *Domain) Query(columns ...basedao.Column) ([]Domain, error) {
	if columns == nil {
		columns = []basedao.Column{t.Remark, t.Id, t.Domain, t.Createtime}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]Domain, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_domain()
		go copy_domain(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copy_domain(channle chan int16, rows []interface{}, t *Domain, columns []basedao.Column) {
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

func (t *Domain) QuerySingle(columns ...basedao.Column) (*Domain, error) {
	if columns == nil {
		columns = []basedao.Column{t.Remark, t.Id, t.Domain, t.Createtime}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_domain()
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

func (t *Domain) Select(columns ...basedao.Column) (*Domain, error) {
	if columns == nil {
		columns = []basedao.Column{t.Remark, t.Id, t.Domain, t.Createtime}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_domain()
		cp_domain(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *Domain) Selects(columns ...basedao.Column) ([]*Domain, error) {
	if columns == nil {
		columns = []basedao.Column{t.Remark, t.Id, t.Domain, t.Createtime}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*Domain, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_domain()
		cp_domain(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cp_domain(buff []interface{}, t *Domain, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "id":
			buff[i] = &t.Id.FieldValue
		case "domain":
			buff[i] = &t.Domain.FieldValue
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "remark":
			buff[i] = &t.Remark.FieldValue
		}
	}
}

func New_domain(tableName ...string) *Domain {
	id := &my_domain_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	domain := &my_domain_Domain{fieldName: "domain"}
	domain.Field.FieldName = "domain"
	createtime := &my_domain_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	remark := &my_domain_Remark{fieldName: "remark"}
	remark.Field.FieldName = "remark"
	table := &Domain{Createtime: createtime, Remark: remark, Id: id, Domain: domain}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_domain"
	}
	return table
}
