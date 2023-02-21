package global

import (
	"gorm.io/gorm"
	"xoj_judgehost/util/setting"
)

const (
	WAITING                  = -6
	PRESENTATION_ERROR       = -5
	COMPILE_ERROR            = -4
	WRONG_ANSWER             = -3
	PENDING                  = -1
	JUDGINNG                 = -2
	CPU_TIME_LIMIT_EXCEEDED  = 1
	REAL_TIME_LIMIT_EXCEEDED = 2
	MEMORY_LIMIT_EXCEEDED    = 3
	RUNTIME_ERROR            = 4
	SYSTEM_ERROR             = 5

	QUEUENAME = "judgeQue"
)

var (
	JudgeSetting    *setting.JudgeSettingS    = (*setting.JudgeSettingS)(nil)
	MySQLSetting    *setting.MySQLSettingS    = (*setting.MySQLSettingS)(nil)
	DBEngine        *gorm.DB                  = (*gorm.DB)(nil)
	RabbitMQSetting *setting.RabbitMQSettingS = (*setting.RabbitMQSettingS)(nil)
)
