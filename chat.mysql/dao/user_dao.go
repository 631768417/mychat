package dao

/**
tablename:my_user
datetime :2016-09-07 11:32:22
*/
import (
	"reflect"

	"chat.mysql/dao/basedao"
)

type my_user_Loginname struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Loginname) Name() string {
	return c.fieldName
}

func (c *my_user_Loginname) Value() interface{} {
	return c.FieldValue
}

type my_user_Username struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Username) Name() string {
	return c.fieldName
}

func (c *my_user_Username) Value() interface{} {
	return c.FieldValue
}

type my_user_Usernick struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Usernick) Name() string {
	return c.fieldName
}

func (c *my_user_Usernick) Value() interface{} {
	return c.FieldValue
}

type my_user_Plainpassword struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Plainpassword) Name() string {
	return c.fieldName
}

func (c *my_user_Plainpassword) Value() interface{} {
	return c.FieldValue
}

type my_user_Encryptedpassword struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Encryptedpassword) Name() string {
	return c.fieldName
}

func (c *my_user_Encryptedpassword) Value() interface{} {
	return c.FieldValue
}

type my_user_Createtime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Createtime) Name() string {
	return c.fieldName
}

func (c *my_user_Createtime) Value() interface{} {
	return c.FieldValue
}

type my_user_Updatetime struct {
	basedao.Field
	fieldName  string
	FieldValue *string
}

func (c *my_user_Updatetime) Name() string {
	return c.fieldName
}

func (c *my_user_Updatetime) Value() interface{} {
	return c.FieldValue
}

type my_user_Id struct {
	basedao.Field
	fieldName  string
	FieldValue *int32
}

func (c *my_user_Id) Name() string {
	return c.fieldName
}

func (c *my_user_Id) Value() interface{} {
	return c.FieldValue
}

type User struct {
	basedao.Table
	Username          *my_user_Username
	Usernick          *my_user_Usernick
	Plainpassword     *my_user_Plainpassword
	Encryptedpassword *my_user_Encryptedpassword
	Createtime        *my_user_Createtime
	Updatetime        *my_user_Updatetime
	Id                *my_user_Id
	Loginname         *my_user_Loginname
}

func (u *User) GetPlainpassword() string {
	return *u.Plainpassword.FieldValue
}

func (u *User) SetPlainpassword(arg string) {
	u.Table.ModifyMap[u.Plainpassword.fieldName] = arg
	v := string(arg)
	u.Plainpassword.FieldValue = &v
}

func (u *User) GetEncryptedpassword() string {
	return *u.Encryptedpassword.FieldValue
}

func (u *User) SetEncryptedpassword(arg string) {
	u.Table.ModifyMap[u.Encryptedpassword.fieldName] = arg
	v := string(arg)
	u.Encryptedpassword.FieldValue = &v
}

func (u *User) GetCreatetime() string {
	return *u.Createtime.FieldValue
}

func (u *User) SetCreatetime(arg string) {
	u.Table.ModifyMap[u.Createtime.fieldName] = arg
	v := string(arg)
	u.Createtime.FieldValue = &v
}

func (u *User) GetUpdatetime() string {
	return *u.Updatetime.FieldValue
}

func (u *User) SetUpdatetime(arg string) {
	u.Table.ModifyMap[u.Updatetime.fieldName] = arg
	v := string(arg)
	u.Updatetime.FieldValue = &v
}

func (u *User) GetId() int32 {
	return *u.Id.FieldValue
}

func (u *User) SetId(arg int64) {
	u.Table.ModifyMap[u.Id.fieldName] = arg
	v := int32(arg)
	u.Id.FieldValue = &v
}

func (u *User) GetLoginname() string {
	return *u.Loginname.FieldValue
}

func (u *User) SetLoginname(arg string) {
	u.Table.ModifyMap[u.Loginname.fieldName] = arg
	v := string(arg)
	u.Loginname.FieldValue = &v
}

func (u *User) GetUsername() string {
	return *u.Username.FieldValue
}

func (u *User) SetUsername(arg string) {
	u.Table.ModifyMap[u.Username.fieldName] = arg
	v := string(arg)
	u.Username.FieldValue = &v
}

func (u *User) GetUsernick() string {
	return *u.Usernick.FieldValue
}

func (u *User) SetUsernick(arg string) {
	u.Table.ModifyMap[u.Usernick.fieldName] = arg
	v := string(arg)
	u.Usernick.FieldValue = &v
}

func (t *User) Query(columns ...basedao.Column) ([]User, error) {
	if columns == nil {
		columns = []basedao.Column{t.Plainpassword, t.Encryptedpassword, t.Createtime, t.Updatetime, t.Id, t.Loginname, t.Username, t.Usernick}
	}
	rs, err := t.Table.Query(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	ts := make([]User, 0, len(rs))
	c := make(chan int16, len(rs))
	for _, rows := range rs {
		t := New_user()
		go copyTim_user(c, rows, t, columns)
		<-c
		ts = append(ts, *t)
	}
	return ts, nil
}

func copyTim_user(channle chan int16, rows []interface{}, t *User, columns []basedao.Column) {
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

func (t *User) QuerySingle(columns ...basedao.Column) (*User, error) {
	if columns == nil {
		columns = []basedao.Column{t.Plainpassword, t.Encryptedpassword, t.Createtime, t.Updatetime, t.Id, t.Loginname, t.Username, t.Usernick}
	}
	rs, err := t.Table.QuerySingle(columns...)
	if rs == nil || err != nil {
		return nil, err
	}
	rt := New_user()
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

func (t *User) Select(columns ...basedao.Column) (*User, error) {
	if columns == nil {
		columns = []basedao.Column{t.Plainpassword, t.Encryptedpassword, t.Createtime, t.Updatetime, t.Id, t.Loginname, t.Username, t.Usernick}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	buff := make([]interface{}, len(columns))
	if rows.Next() {
		n := New_user()
		cpTim_user(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		return n, nil
	}
	return nil, nil
}

func (t *User) Selects(columns ...basedao.Column) ([]*User, error) {
	if columns == nil {
		columns = []basedao.Column{t.Plainpassword, t.Encryptedpassword, t.Createtime, t.Updatetime, t.Id, t.Loginname, t.Username, t.Usernick}
	}
	rows, err := t.Table.Selects(columns...)
	defer rows.Close()
	if err != nil || rows == nil {
		return nil, err
	}
	ns := make([]*User, 0)
	buff := make([]interface{}, len(columns))
	for rows.Next() {
		n := New_user()
		cpTim_user(buff, n, columns)
		row_err := rows.Scan(buff...)
		if row_err != nil {
			return nil, row_err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func cpTim_user(buff []interface{}, t *User, columns []basedao.Column) {
	for i, column := range columns {
		field := column.Name()
		switch field {
		case "createtime":
			buff[i] = &t.Createtime.FieldValue
		case "updatetime":
			buff[i] = &t.Updatetime.FieldValue
		case "id":
			buff[i] = &t.Id.FieldValue
		case "loginname":
			buff[i] = &t.Loginname.FieldValue
		case "username":
			buff[i] = &t.Username.FieldValue
		case "usernick":
			buff[i] = &t.Usernick.FieldValue
		case "plainpassword":
			buff[i] = &t.Plainpassword.FieldValue
		case "encryptedpassword":
			buff[i] = &t.Encryptedpassword.FieldValue
		}
	}
}

func New_user(tableName ...string) *User {
	usernick := &my_user_Usernick{fieldName: "usernick"}
	usernick.Field.FieldName = "usernick"
	plainpassword := &my_user_Plainpassword{fieldName: "plainpassword"}
	plainpassword.Field.FieldName = "plainpassword"
	encryptedpassword := &my_user_Encryptedpassword{fieldName: "encryptedpassword"}
	encryptedpassword.Field.FieldName = "encryptedpassword"
	createtime := &my_user_Createtime{fieldName: "createtime"}
	createtime.Field.FieldName = "createtime"
	updatetime := &my_user_Updatetime{fieldName: "updatetime"}
	updatetime.Field.FieldName = "updatetime"
	id := &my_user_Id{fieldName: "id"}
	id.Field.FieldName = "id"
	loginname := &my_user_Loginname{fieldName: "loginname"}
	loginname.Field.FieldName = "loginname"
	username := &my_user_Username{fieldName: "username"}
	username.Field.FieldName = "username"
	table := &User{Createtime: createtime, Updatetime: updatetime, Id: id, Loginname: loginname, Username: username, Usernick: usernick, Plainpassword: plainpassword, Encryptedpassword: encryptedpassword}
	table.Table.ModifyMap = make(map[string]interface{})
	if len(tableName) == 1 {
		table.Table.TableName = tableName[0]
	} else {
		table.Table.TableName = "my_user"
	}
	return table
}
