package config

//MysqlConf export this object
var MysqlConf MysqlConfig
var EtcdConf EtcdConfig

//MysqlConfig configuration for Mysql connection
type MysqlConfig struct {
	Host     string
	Port     uint32
	Username string
	Password string
	Database string
	TimeLoc  string
}

//EtcdConfig configureation for Etcd
type EtcdConfig struct {
	Addr string
}

func init() {
	MysqlConf = MysqlConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "124563",
		Database: "go-video",
		TimeLoc:  "Asia/Shanghai",
	}

	EtcdConf = EtcdConfig{
		Addr: "http://localhost:2379",
	}
}
