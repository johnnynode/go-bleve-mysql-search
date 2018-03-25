package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Country struct {
	Id            int    `json:"id" orm:"column(id);auto"`
	Namecn        string `json:"namecn" orm:"column(namecn);size(255);null"`
	Name          string `json:"name" orm:"column(name);size(255);null"`
	Countrycode   string `json:"countrycode" orm:"column(countrycode);size(255);null"`
	Continentname string `json:"continentname" orm:"column(continentname);size(255);null"`
}

func (t *Country) TableName() string {
	return "country"
}

func init() {
	orm.RegisterModel(new(Country))
}

// AddCountry insert a new Country into database and returns
// last inserted Id on success.
func AddCountry(m *Country) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCountryById retrieves Country by Id. Returns error if
// Id doesn't exist
func GetCountryById(id int) (v *Country, err error) {
	o := orm.NewOrm()
	v = &Country{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllCountry() ([]*Country, error) {
	o := orm.NewOrm()
	var CountryList []*Country
	_,err := o.QueryTable("country").All(&CountryList)

	if err == nil {
		return CountryList, nil
	}
	return nil, err
}

// 通过id更新国家
func UpdateCountryById(mp map[string]interface{}) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("country").Filter("id", mp["id"]).Update(mp)
	return err
}

// 通过id删除国家接口
func DeleteCountry(id int) (err error) {
	o := orm.NewOrm()
	v := Country{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Country{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
