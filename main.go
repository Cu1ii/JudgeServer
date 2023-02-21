package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"xoj_judgehost/global"
	"xoj_judgehost/internal/judge"
	"xoj_judgehost/internal/model"
	"xoj_judgehost/logs"
	"xoj_judgehost/util/setting"
)

func init() {
	logs.InitRuntimeLog()
	err := setupSetting()
	if err != nil {
		logrus.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupDBEngine()
	if err != nil {
		logrus.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func main() {
	logrus.Info("start run judge")
	judge.RunJudge()
}

func setupSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = set.ReadSection("judge-environment", &global.JudgeSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("rabbitmq", &global.RabbitMQSetting)
	if err != nil {
		return err
	}

	//// 如果需要通过配置文件读取就取消掉注释
	//err = set.ReadSection("mysql", &global.MySQLSetting)
	//if err != nil {
	//	return err
	//}
	// 不采用 docker 设置环境变量就注释掉
	GetMySQLConfigByEnv()
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.MySQLSetting)
	if err != nil {
		return err
	}

	return nil
}

func GetMySQLConfigByEnv() *setting.MySQLSettingS {
	global.MySQLSetting = &setting.MySQLSettingS{
		DisableDatetimePrecision:  true,
		DefaultStringSize:         256,
		DontSupportRenameColumn:   true,
		DontSupportRenameIndex:    true,
		SkipInitializeWithVersion: false,
	}
	usr := os.Getenv("mysql-user")
	pwd := os.Getenv("mysql-pwd")
	host := os.Getenv("mysql-host")
	port := os.Getenv("mysql-port")
	dbname := os.Getenv("mysql-dbname")
	global.MySQLSetting.DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", usr, pwd, host, port, dbname)
	return global.MySQLSetting
}
