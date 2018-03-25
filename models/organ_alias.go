package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type OrganAlias struct {
	Id       int       `json:"id" orm:"column(id);auto"`
	Uuid     string    `json:"uuid" orm:"column(uuid);size(45);null" description:"标识"`
	Name     string    `json:"name" orm:"column(name);size(255);null" description:"别名"`
	CbUser   string    `json:"cbUser" orm:"column(cb_user);size(45);null"`
	CbTime   time.Time `json:"cbTime" orm:"column(cb_time);type(datetime);null;auto_now_add"`
	EditUser string    `json:"editUser" orm:"column(edit_user);size(45);null"`
	EditTime time.Time `json:"editTime" orm:"column(edit_time);type(datetime);null;auto_now"`
}

func (t *OrganAlias) TableName() string {
	return "organ_alias"
}

func init() {
	orm.RegisterModel(new(OrganAlias))
}

// 添加用户别名
func AddOrganAlias(m *OrganAlias) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// 通过uuid获取机构别名
func GetOrganAliasByUuid(uuid string) (oa []*OrganAlias, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("organ_alias").Filter("Uuid", uuid).All(&oa)
	return oa, err
}

// 通过uuid获取机构别名, 只获取名称
func GetOrganAliasNameByUuid(uuid string) (oaName []orm.Params, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("organ_alias").Filter("Uuid", uuid).Values(&oaName,"Name")
	return
}

// 通过id更新机构别名
func UpdateOrganAliasById(m *OrganAlias) (err error) {
	o := orm.NewOrm()
	v := OrganAlias{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 通过id删除机构别名
func DeleteOrganAlias(id int) (err error) {
	o := orm.NewOrm()
	v := OrganAlias{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OrganAlias{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
