package configuration

import (
	"fmt"
	"testing"
)

func TestGetMySQLConfig(t *testing.T) {
	mySQLConfig := GetMySQLConfig()
	fmt.Println(mySQLConfig)
}

func TestGetJudgeEnvironmentConfig(t *testing.T) {
	judgeEnvironmentConfig := GetJudgeEnvironmentConfig()
	fmt.Println(judgeEnvironmentConfig)
}
