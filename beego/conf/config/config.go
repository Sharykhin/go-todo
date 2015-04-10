package config

import (
	 "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)

    orm.RegisterDataBase("default", "mysql", "root:pass4root@/test?charset=utf8")
    
}