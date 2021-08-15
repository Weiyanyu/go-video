package mysqldb

import (
	"fmt"
	"go-video/src/common/log4video"
	userconfig "go-video/src/services/user/config"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlDB *gorm.DB

//GetConnection for Get Mysql connection object
func GetConnection() *gorm.DB {
	log4video.D("mysql addr(%p)", mysqlDB)
	if mysqlDB != nil {
		return mysqlDB
	}

	DSN := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		userconfig.MysqlConf.Username,
		userconfig.MysqlConf.Password,
		userconfig.MysqlConf.Host,
		userconfig.MysqlConf.Port,
		userconfig.MysqlConf.Database,
		url.QueryEscape(userconfig.MysqlConf.TimeLoc))

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		//can recover by go-micro framework
		panic(err)
	}

	mysqlDB = db

	log4video.D("user service already connect to mysql server(%s:%s), conn(%p)",
		userconfig.MysqlConf.Host,
		userconfig.MysqlConf.Port,
		mysqlDB)

	return mysqlDB

}
