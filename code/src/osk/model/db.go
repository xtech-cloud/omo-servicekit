package model

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"osk/core"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
)

var base64Coder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

var db *gorm.DB

func Initialize() {
	if core.Config.DBDriver == "mysql" {
		core.Logger.Info("use mysql, database is " + core.Config.DBName)
		conn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", core.Config.MysqlUser, core.Config.MysqlPasswd, core.Config.DBName)
		database, err := gorm.Open("mysql", conn)
		if nil != err {
			panic(err)
		}
		db = database
	} else {
		dbname := fmt.Sprintf("%s/data/%s.db", core.Env.RunPath, core.Config.DBName)
		core.Logger.Infof("use sqlite, database is %s", dbname)
		database, err := gorm.Open("sqlite3", dbname)
		if nil != err {
			panic(err)
		}
		db = database
	}

	var err error

	err = MigrateAccount()
	if nil != err {
		panic(err)
	}

	core.Logger.Infof("migrate tables success")
}

func Release() {
	db.Close()
}

func Save(val interface{}) *gorm.DB {
	return db.Save(val)
}

func NewUUID() string {
	guid, _ := uuid.NewV4()
	h := md5.New()
	h.Write([]byte(guid.String()))
	return hex.EncodeToString(h.Sum(nil))
}

func ToUUID(_content string) string {
	h := md5.New()
	h.Write([]byte(_content))
	return hex.EncodeToString(h.Sum(nil))
}

func ToBase64(_content []byte) string {
	return base64Coder.EncodeToString(_content)
}
