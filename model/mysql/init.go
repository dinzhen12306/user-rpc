package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
	"zg5/day1/rpc/userrpc/config"
)

var (
	XDB *xorm.Engine
	err error
)

func init() {
	XDB, err = xorm.NewEngine(config.ServerConfigs.Mysql.DriverName, fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", config.ServerConfigs.Mysql.Username, config.ServerConfigs.Mysql.Password, config.ServerConfigs.Mysql.ConnectionType, config.ServerConfigs.Mysql.Host, config.ServerConfigs.Mysql.Port, config.ServerConfigs.Mysql.DatabaseName, config.ServerConfigs.Mysql.ConnectionParameters))
	if err != nil {
		log.Println("mysql连接失败", err.Error())
	}
	err = XDB.Sync2(new(User))
	if err != nil {
		log.Println("mysql数据表创建失败", err.Error())
	}
	log.Println("mysql连接成功")
}
