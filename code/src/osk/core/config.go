package core

import (
	"github.com/miguel-branco/goconfig"
)

const (
	default_host            string = ":18993"
	default_loglevel        string = "INFO"
	default_logfile         string = "/var/log/osk.log"
	default_gin_mode        string = "release"
	default_database_name   string = "osk"
	default_database_driver string = "sqlite"
	default_mysql_user      string = "root"
	default_mysql_passwd    string = "mysql@OMO"
)

type _Config struct {
	Host        string
	LogLevel    string
	LogFile     string
	GinMode     string
	DBName      string
	DBDriver    string
	MysqlUser   string
	MysqlPasswd string
}

var Config _Config

func SetupConfig() {
	conf, err := goconfig.ReadConfigFile(Env.RunPath + "/conf/osk.cfg")
	if nil != err {
		panic(err)
	}

	Config.Host, err = conf.GetString("core", "host")
	if nil != err {
		Config.Host = default_host
	}

	Config.LogLevel, err = conf.GetString("core", "loglevel")
	if nil != err {
		Config.LogLevel = default_loglevel
	}
	Config.LogFile, err = conf.GetString("core", "logfile")
	if nil != err {
		Config.LogFile = default_logfile
	}

	Config.DBName, err = conf.GetString("db", "name")
	if nil != err {
		Config.DBName = default_database_name
	}

	Config.DBDriver, err = conf.GetString("db", "driver")
	if nil != err {
		Config.DBDriver = default_database_driver
	}

	Config.GinMode, err = conf.GetString("gin", "mode")
	if nil != err {
		Config.GinMode = default_gin_mode
	}

	Config.MysqlUser, err = conf.GetString("mysql", "user")
	if nil != err {
		Config.MysqlUser = default_mysql_user
	}

	Config.MysqlPasswd, err = conf.GetString("mysql", "passwd")
	if nil != err {
		Config.MysqlPasswd = default_mysql_passwd
	}
}
