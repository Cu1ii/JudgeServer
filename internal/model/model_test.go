package model

import (
	"fmt"
	"testing"
	"xoj_judgehost/util/setting"
)

func TestGetDBEngine(t *testing.T) {
	_, err := NewDBEngine(&setting.MySQLSettingS{
		DSN: "username:password@tcp(localhost:3306)/judge_backend?charset=utf8&parseTime=True&loc=Local",
	})
	fmt.Println("err = ", err)
}
