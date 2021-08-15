package mysqldb

import (
	"fmt"
	userconfig "go-video/src/services/user/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

//GetConnection for Get Mysql connection object
func GetConnection() *gorm.DB {
	if mysqlDB != nil {
		return mysqlDB
	}

	DSN := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		userconfig.MysqlConf.Username,
		userconfig.MysqlConf.Password,
		userconfig.MysqlConf.Host,
		userconfig.MysqlConf.Port,
		userconfig.MysqlConf.Database,
		userconfig.MysqlConf.TimeLoc)

	mysqlDB, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		//can recover by go-micro framework
		panic(err)
	}

	return mysqlDB

}
