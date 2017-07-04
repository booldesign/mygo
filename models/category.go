package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// Category Struct
type Category struct {
	CateId    int64  `orm:"pk;auto"`
	Title     string `orm:"index"`
	CreatedAt time.Time
	Click     int
	IsDel     int `orm:"default(0)"`
	Price     float32
}

func (m *Category) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Category) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
