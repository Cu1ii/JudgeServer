package setting

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

var (
	MySQLSetting = MySQLSettingS{}
	JudgeSetting = JudgeSettingS{}
)

func TestSetting_ReadSection(t *testing.T) {

	err := setupSetting()
	if err != nil {
		logrus.Fatalf("init.setupSetting err: %v", err)
	}

	fmt.Println(MySQLSetting)
	fmt.Println(JudgeSetting)
}

func setupSetting() error {
	set, err := NewSetting()
	if err != nil {
		return err
	}

	err = set.ReadSection("judge-environment", &JudgeSetting)
	if err != nil {
		return err
	}

	err = set.ReadSection("mysql", &MySQLSetting)
	if err != nil {
		return err
	}
	return nil
}
