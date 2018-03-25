package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Organ struct {
	Id             int              `json:"id" orm:"column(id);auto"`
	Uuid           string           `json:"uuid" orm:"column(uuid);size(45);null" description:"唯一标识"`
	Oid            string           `json:"oid" orm:"column(oid);size(255);null" description:"机构ID"`
	Type           int              `json:"type" orm:"column(type);null" description:"类型"`
	NameEn         string           `json:"nameEn" orm:"column(name_en);size(255);null" description:"英文名"`
	NameCn         string           `json:"nameCn" orm:"column(name_cn);size(255);null" description:"中文名"`
	NameOrigin     string           `json:"nameOrigin" orm:"column(name_origin);size(255);null" description:"原始名称"`
	ShortName      string           `json:"shortName" orm:"column(short_name);size(255);null" description:"简称"`
	FirstLetter    string           `json:"firstLetter" orm:"column(first_letter);size(1);null" description:"首字母"`
	Location       string           `json:"location" orm:"column(location);size(45);null" description:"地域"`
	CountryCode    string           `json:"countryCode" orm:"column(country_code);size(2);null" description:"国别二位码"`
	CountryCn      string           `json:"countryCn" orm:"column(country_cn);size(45);null" description:"国家中文名"`
	CountryEn      string           `json:"countryEn" orm:"column(country_en);size(45);null" description:"国家英文名"`
	Province       string           `json:"province" orm:"column(province);size(45);null" description:"省,州"`
	ProvinceCode   string           `json:"provinceCode" orm:"column(province_code);size(45);null" description:"省州代码"`
	ProvincePinyin string           `json:"provincePinyin" orm:"column(province_pinyin);size(45);null" description:"省州拼音"`
	City           string           `json:"city" orm:"column(city);size(45);null" description:"城市"`
	AbsEn          string           `json:"absEn" orm:"column(abs_en);null" description:"英文简介"`
	AbsCn          string           `json:"absCn" orm:"column(abs_cn);null" description:"中文简介"`
	Logo           string           `json:"logo" orm:"column(logo);size(255);null" description:"学校logo"`
	Banner         string           `json:"banner" orm:"column(banner);size(255);null" description:"横幅图片"`
	BgImg          string           `json:"bgImg" orm:"column(bg_img);size(255);null" description:"背景图片"`
	QsRank         int              `json:"qsRank" orm:"column(qs_rank);null" description:"QS最新排行"`
	Editor         string           `json:"editor" orm:"column(editor);size(45);null" description:"修改人"`
	EditTime       time.Time        `json:"editTime" orm:"column(edit_time);type(datetime);null;auto_now" description:"修改时间"`
	Creator        string           `json:"creator" orm:"column(creator);size(45);null" description:"采集人"`
	CreateTime     time.Time        `json:"createTime" orm:"column(create_time);type(datetime);null;auto_now_add" description:"采集时间"`

	OrganAlias	   []*OrganAlias	`json:"organAlias" orm:"-"`
	Oa			   string			`json:"oa" orm:"-"`
}

func (t *Organ) TableName() string {
	return "organ"
}

func init() {
	orm.RegisterModel(new(Organ))
}

// 通过关键字获取查询总数
func GetTotalByKeyword(keyword string) int64 {
	cond := orm.NewCondition()
	cond1 := cond.Or("oid__icontains", keyword).Or("country_code__icontains", keyword).Or("uuid__icontains", keyword) // 模糊查询功能
	total, err := orm.NewOrm().QueryTable("organ").SetCond(cond1).Count()
	if err != nil {
		return 0
	} else {
		return total
	}
}

// 获取机构数据库总条数
func GetTotal() int64 {
	total, err := orm.NewOrm().QueryTable("organ").Count()
	if err != nil {
		return 0
	} else {
		return total
	}
}

// 通过当前页来输出数据
func GetOrganByPage(page int64, limit int64) (lists []*Organ, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("organ").Limit(limit, (page * limit)).All(&lists)

	if err == nil {
		return lists, nil
	}
	return nil, err
}

// 通过关键字(oid,country_code,uuid)和页码来模糊获取查询出的结果集
func GetOrganList(keyword string, num int64) (lists []orm.Params, err error) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("oid__icontains", keyword).Or("country_code__icontains", keyword).Or("uuid__icontains", keyword) // 模糊查询功能
	_, err = o.QueryTable("organ").SetCond(cond1).Limit(20, (num - 1)).Values(&lists,
		"Id", "Uuid", "Oid", "Type", "NameEn", "NameCn","NameEn","NameOrigin","ShortName","FirstLetter","Location",
		"CountryCode","CountryCn","CountryEn","Province","ProvinceCode","ProvincePinyin","City","AbsEn","AbsCn","Logo",
		"Banner","BgImg","QsRank","Editor","EditTime","Creator","CreateTime") // 此处num在外部要进行限制，num>=1 // 省略字段：,"OrganAlias"

	if err == nil {
		return lists, nil
	}
	return nil, err
}

// 添加机构信息
func AddOrgan(m *Organ) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return id, err
}

// 通过uuid获取机构详情
func GetOrganByUuid(uuid string) (v *Organ, err error) {
	o := orm.NewOrm()
	v = &Organ{Uuid: uuid}
	// v.OrganAlias = &OrganAlias{Uuid:uuid}
	if err = o.Read(v, "Uuid"); err == nil {
		// o.Read(v.OrganAlias, "Uuid") // 读取
		return v, nil
	}
	return nil, err
}

// 通过id更新机构信息
func UpdateOrgan(m *Organ) (err error) {
	o := orm.NewOrm()
	var num int64
	if num, err = o.Update(m); err == nil {
		fmt.Println("Number of records updated in database:", num)
	}

	return
}

// 通过uuid删除机构信息
func DeleteOrgan(uuid string) (err error) {
	o := orm.NewOrm()
	var num int64
	if num, err = o.Delete(&Organ{Uuid: uuid}, "Uuid"); err == nil {
		fmt.Println("Number of records deleted in database:", num)
	}
	return
}