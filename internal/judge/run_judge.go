package judge

import (
	"github.com/sirupsen/logrus"
	"time"
	"xoj_judgehost/global"
	"xoj_judgehost/internal/dao"
	"xoj_judgehost/util/pool"
)

func RunJudge() {
	go changeAuth()
	for true {
		time.Sleep(time.Second * 2)
		pendingStatus := dao.GetJudgeStatus()
		for _, status := range pendingStatus {
			dao.UpdateJudgeStatusResult(int(status.Id), global.WAITING)
			dao.UpdateJudgeStatusJudger(int(status.Id), "XOJ")
		}
		for _, status := range pendingStatus {
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

// 比赛题目设置为auth=2,contest开始时，自动设置题目为auth=3，比赛结束自动设置auth=1
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
