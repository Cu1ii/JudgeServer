package judge

import (
	"github.com/sirupsen/logrus"
	"xoj_judgehost/internal/dao"
	"xoj_judgehost/internal/entity"
)

func compileError(id int, problem, msg string) {
	logrus.Info("Compile error! ", id)
	dao.UpdateJudgeStatusResult(id, -4)
	dao.UpdateJudgeStatusMessage(id, msg)
	dao.UpdateProblemData(problem, "ce")
}

func doneProblem(id int, problem, message string, memory, mytime int, username string, contest, result int, testcase string) {
	// fmt.Println(id, " ", problem, " ", message, " ", mytime, " ", username, " ", contest, " ", result, " ", testcase)
	dao.UpdateJudgeStatus(id, memory, mytime, result, testcase)
	if message != "" {
		dao.UpdateJudgeStatusMessage(id, message)
	}
	if result == 2 || result == 1 {
		dao.UpdateProblemData(problem, "tle")
	}
	if result == 3 {
		dao.UpdateProblemData(problem, "mle")
	}
	if result == 4 {
		dao.UpdateProblemData(problem, "rte")
	}
	if result == 5 {
		dao.UpdateProblemData(problem, "se")
	}
	if result == -5 {
		dao.UpdateProblemData(problem, "pe")
	}
	if result == -3 {
		dao.UpdateProblemData(problem, "wa")
	}
	if contest != 0 {
		dao.UpdateContestBoardTypeBySubmitId(0, id)
	}
	dao.UpdateUserResult(username, "submit")
}

func acProblem(id int, problem, message string, memory, time int, username string, proScore int, isAc bool, contest int) {
	//fmt.Println(time, " ", memory)
	dao.UpdateJudgeStatus(id, memory, time, 0, "")
	if message != "" {
		dao.UpdateJudgeStatusMessage(id, message)
	}
	dao.UpdateProblemData(problem, "ac")
	if !isAc {
		dao.UpdateUserResult(username, "ac")
		dao.UpdateUserScore(username, proScore)
		dao.UpdateUserAcPro(problem, username)
	}
	if contest != 0 {
		dao.UpdateContestBoardTypeBySubmitId(1, id)
	}
	dao.UpdateUserResult(username, "submit")
}

func doneCase(statusId int, username, problem, result string,
	time, memory int, testcase, caseData, outputData, userOutput string) {
	dao.AddCaseStatus(&entity.CaseStatus{
		StatusId:   statusId,
		Username:   username,
		Problem:    problem,
		Result:     result,
		Time:       time,
		Memory:     memory,
		TestCase:   testcase,
		CaseData:   caseData,
		OutputData: outputData,
		UserOutput: userOutput,
	})
}
