package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// User Struct
type User struct {
	UserId             int64     `orm:"pk;auto"`
	Username           string    `orm:"unique;size(30)"`
	AuthKey            string    `orm:"size(32);"`
	PasswordHash       string    `orm:"size(50)"`
	PasswordResetToken string    `orm:"size(50)"`
	Email              string    `orm:"unique;size(50)"`
	Status             int       `orm:"default(0)"`
	CreatedAt          time.Time `orm:"type(datetime)"`
	UpdatedAt          time.Time `orm:"type(datetime)"`
}

// 自定义表名
func (m *User) TableName() string {
	return "user_index"
}

// 多字段索引(复合索引) CREATE INDEX `tbl_user_index_user_id_username` ON `tbl_user_index` (`user_id`, `username`);
func (m *User) TableIndex() [][]string {
	return [][]string{
		[]string{"UserId", "Username"},
	}
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
