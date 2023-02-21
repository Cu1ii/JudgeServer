package judge

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
	"xoj_judgehost/global"
	"xoj_judgehost/internal/dao"
	"xoj_judgehost/util/pool"
	"xoj_judgehost/util/rabbitmq"
)

func RunJudge() {
	channel, connection, err := rabbitmq.NewRabbitMQConnect(global.RabbitMQSetting)
	if err != nil {
		logrus.Fatalf("RabbitMQ connect error %s", err.Error())
	}
	defer connection.Close()
	defer channel.Close()

	q, err := channel.QueueDeclare(
		global.QUEUENAME, // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)

	// 获取接收消息的Delivery通道
	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logrus.Error("Failed to register a consumer %s", err.Error())
	}
	// go changeAuth()
	for true {
		fmt.Println("get judgeStatusId")
		for d := range msgs {
			msg := string(d.Body)
			atoi, err := strconv.Atoi(msg)
			if err != nil {
				logrus.Error(err)
			}
			status := dao.GetJudgeStatusById(int64(atoi))
			if err := pool.GetJudgePool().Submit(
				func() {
					judge(
						int(status.Id),
						status.Code,
						status.Language,
						status.Problem,
						int(status.Contest),
						status.User,
						status.Oj,
						"XOJ",
						status.SubmitTime,
						status.ContestProblem,
						false,
					)
				}); err != nil {
				logrus.Error("run judge ", status.Problem, " the user is ", status.User, " error: ", err.Error())
			}
		}
	}
}

// 比赛题目设置为auth=2,contest开始时，自动设置题目为auth=3，比赛结束自动设置auth=1 暂时废弃
func changeAuth() {
	curContest := map[int]bool{}
	curPro := map[string]bool{}
	curRunPro := map[string]bool{}
	for true {
		time.Sleep(time.Second * 2)
		allContest := map[int]bool{}
		notExpiredContests := dao.GetNotExpiredContest()
		for _, notExpiredContest := range notExpiredContests {
			allContest[notExpiredContest.Id] = true
			contestProblems := dao.GetContestProblem(notExpiredContest.Id)
			for _, pro := range contestProblems {
				if _, ok := curPro[pro.ProblemId]; !ok {
					logrus.Info("58 pro.Problem", pro.ProblemId)
					curPro[pro.ProblemId] = true
					dao.UpdateProblemDataAuth(pro.ProblemId, 2)
					dao.UpdateProblemAuth(pro.ProblemId, 2)
				}
			}
		}
		runningContests := dao.GetRunningContest()
		for _, runningContest := range runningContests {
			contestProblems := dao.GetContestProblem(runningContest.Id)
			for _, pro := range contestProblems {
				if _, ok := curRunPro[pro.ProblemId]; !ok {
					curRunPro[pro.ProblemId] = true
					dao.UpdateProblemDataAuth(pro.ProblemId, 3)
					dao.UpdateProblemAuth(pro.ProblemId, 3)
				}
			}
		}
		for contestId, _ := range curContest {
			if _, ok := allContest[contestId]; !ok {
				contestProblems := dao.GetContestProblem(contestId)
				for _, pro := range contestProblems {
					if _, ok := curPro[pro.ProblemId]; ok {
						delete(curPro, pro.ProblemId)
					}
					if _, ok := curRunPro[pro.ProblemId]; ok {
						delete(curRunPro, pro.ProblemId)
					}
					dao.UpdateProblemDataAuth(pro.ProblemId, 1)
					dao.UpdateProblemAuth(pro.ProblemId, 1)
				}
			}
		}
		curContest = allContest
	}
}
