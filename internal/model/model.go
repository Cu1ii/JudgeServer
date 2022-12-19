package model

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xoj_judgehost/util/setting"
)

func NewDBEngine(mysqlSetting *setting.MySQLSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlSetting.DSN,                       // DSN data source name
		DisableDatetimePrecision:  mysqlSetting.DisableDatetimePrecision,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    mysqlSetting.DontSupportRenameIndex,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   mysqlSetting.DontSupportRenameColumn,   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: mysqlSetting.SkipInitializeWithVersion, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		logrus.Error("open mysql error: ", err)
		return nil, err
	}
	logrus.Info("open mysql success")
	return db, nil
}
