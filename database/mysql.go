package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"xoj_judgehost/configuration"
)

var MySQLDB = (*gorm.DB)(nil)
var mu = sync.Mutex{}

func NewMySQLDB() *gorm.DB {
	mySQLConfig := configuration.GetMySQLConfig()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mySQLConfig.DSN,                       // DSN data source name
		DisableDatetimePrecision:  mySQLConfig.DisableDatetimePrecision,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    mySQLConfig.DontSupportRenameIndex,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   mySQLConfig.DontSupportRenameColumn,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: mySQLConfig.SkipInitializeWithVersion, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		logrus.Error("open mysql error: ", err)
	}
	logrus.Info("open mysql success")
	return db
}

func GetMySQLDB() *gorm.DB {
	if MySQLDB == nil {
		mu.Lock()
		if MySQLDB == nil {
			MySQLDB = NewMySQLDB()
		}
		mu.Unlock()
	}
	return MySQLDB
}
