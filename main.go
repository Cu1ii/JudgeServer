package main

import (
	"github.com/sirupsen/logrus"
	"xoj_judgehost/judge"
	"xoj_judgehost/logs"
)

func main() {
	logs.InitRuntimeLog()
	logrus.Info("start run judge")
	judge.RunJudge()
}
