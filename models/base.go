package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func RegisterDB() {

	// // 注册 model 并使用表名前缀
	// orm.RegisterModelWithPrefix("tbl_", new(User), new(Category))

	// // 注册 model
	// // orm.RegisterModel(new(User), new(Category)

	// // 参数4(可选)  设置最大空闲连接
	// // 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	// maxIdle := 30
	// maxConn := 30
	// // set default database
	// orm.RegisterDataBase("default", "mysql", "root:123@tcp(192.168.2.62:3306)/gotest?charset=utf8mb4", maxIdle, maxConn)

	dbhost := beego.AppConfig.String("mysqlhost")
	dbport := beego.AppConfig.String("mysqlport")
	dbuser := beego.AppConfig.String("mysqluser")
	dbpassword := beego.AppConfig.String("mysqlpwd")
	dbname := beego.AppConfig.String("mysqldb")
	prefix := beego.AppConfig.String("mysqlprefix")
	charset := beego.AppConfig.String("mysqlcharset")
	timezone := beego.AppConfig.String("timezone")

	// 注册 model 并使用表名前缀
	orm.RegisterModelWithPrefix(prefix, new(User), new(Category))
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=" + charset + "&loc=" + timezone

	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn)
}

func UserAdd(o orm.Ormer, user User) (int64, error) {
	// insert
	return o.Insert(&user)
}
