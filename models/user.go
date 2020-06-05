package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Avatar  string `json:"avatar"`
	UserId  string `json:"userid"`
	Address string `json:"address"`
}

type UserLogin struct {
	Status           string `json:"status"`
	Type             string `json:"type"`
	CurrentAuthority string `json:"currentAuthority"`
}

type UserModel struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func (m *UserModel) TableName() string {
	return "users"
}

func (m *UserModel) Count() int64 {
	o := orm.NewOrm()

	r,e := o.QueryTable(m).Count()

	if e!=nil {
		logs.Error("count users failed ", e.Error())
		return 0
	}

	return r
}

func (m *UserModel) GetList(offset,limit int) []*UserModel {
	o := orm.NewOrm()

	userList := make([]*UserModel, 0)

	_,e := o.QueryTable(m).Offset(offset).Limit(limit).OrderBy("-id").All(&userList)


	if e!=nil {
		logs.Error("count users failed ", e.Error())
	}

	return userList
}

func (m *UserModel) Save() bool {
	o := orm.NewOrm()

	_,e := o.Insert(m)
	if e!=nil {
		logs.Error("insert users failed ", e.Error())
		return false
	}

	return true
}