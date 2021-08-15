package config

//MysqlConf export this object
var MysqlConf MysqlConfig
var EtcdConf EtcdConfig
var MicroServiceConf MicroServiceConfig

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

type MicroServiceConfig struct {
	ServiceName string
}

func init() {
	MysqlConf = MysqlConfig{
		Host:     "192.168.31.76",
		Port:     3306,
		Username: "root",
		Password: "124563",
		Database: "go-video",
		TimeLoc:  "Asia/Shanghai",
	}

	EtcdConf = EtcdConfig{
		Addr: "http://localhost:2379",
	}

	MicroServiceConf = MicroServiceConfig{
		ServiceName: "go.video.service.user",
	}
}
