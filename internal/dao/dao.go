package dao

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"xoj_judgehost/global"
	"xoj_judgehost/internal/entity"
)

//---------------------------------JudgeStatus-----------------------------------------//

func GetJudgeAllStatus() []*entity.JudgeStatus {
	judgeArry := []*entity.JudgeStatus{}
	if res := global.DBEngine.Raw("SELECT * FROM judgestatus_judgestatus").Scan(&judgeArry); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return judgeArry
}

func GetJudgeStatus() []*entity.JudgeStatus {
	judgeArry := []*entity.JudgeStatus{}
	if res := global.DBEngine.Raw("SELECT * FROM judgestatus_judgestatus where result = ?", global.PENDING).Scan(&judgeArry); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return judgeArry
}

func GetJudgeStatusById(id int64) *entity.JudgeStatus {
	status := entity.JudgeStatus{}
	if res := global.DBEngine.Raw("SELECT * FROM judgestatus_judgestatus where id = ?", id).Scan(&status); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return &status
}

func UpdateJudgeStatusResult(id int, result int) bool {
	if res := global.DBEngine.Exec("UPDATE judge_backend.judgestatus_judgestatus SET result = ? WHERE id = ?", result, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatusMessage(id int, msg string) bool {
	if res := global.DBEngine.Exec("UPDATE judge_backend.judgestatus_judgestatus SET message = ? WHERE id = ?", msg, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatusJudger(id int, judger string) bool {
	if res := global.DBEngine.Exec("UPDATE judge_backend.judgestatus_judgestatus SET judger = ? WHERE id = ?", judger, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatus(id int, memory, mytime int, result int, testcase string) bool {
	if res := global.DBEngine.Exec("UPDATE judgestatus_judgestatus SET memory = ?, time= ?, result = ?, testcase=?  WHERE id = ?", memory, mytime, result, testcase, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

//---------------------------------Problem-----------------------------------------//

func GetProblemById(pk string) *entity.Problem {
	problem := entity.Problem{}
	if res := global.DBEngine.Raw("SELECT * FROM problem_problem WHERE problem = ?", pk).Scan(&problem); res.Error != nil {
		logrus.Error("select problem error ", res.Error)
		return nil
	}
	return &problem
}

func GetIsHaveDoneProblem(username, problem string) bool {
	selectProblem := fmt.Sprintf("SELECT * FROM judgestatus_judgestatus WHERE user = '%s'  AND problem = '%s' AND result = 0", username, problem)
	problems := []entity.Problem{}
	if res := global.DBEngine.Raw(selectProblem).Scan(&problems); res.Error != nil {
		logrus.Error("select problem error ", res.Error)
		return false
	}
	if len(problems) > 0 {
		return true
	}
	return false
}

func AddProSubmitNum(problem string) bool {
	addProSubmitNum := fmt.Sprintf("UPDATE problem_problemdata SET submission = submission+1 WHERE problem = '%s'", problem)
	if res := global.DBEngine.Exec(addProSubmitNum); res.Error != nil {
		logrus.Error("add problem data ( submission = submission + 1 ) error ", res.Error)
		return false
	}
	return true
}

func GetProblemTimeMemory(pk string) (int, int) {
	problem := GetProblemById(pk)
	return problem.Time, problem.Memory
}

func GetProblemScore(pk string) int {
	problemData := GetProblemDataById(pk)
	return problemData.Score
}

func GetProblemDataById(pk string) *entity.ProblemData {
	problemData := entity.ProblemData{}
	if res := global.DBEngine.Raw("SELECT * FROM problem_problemdata WHERE problem = ?", pk).Scan(&problemData); res.Error != nil {
		logrus.Error("select problem data error ", res.Error)
		return nil
	}
	return &problemData
}

func UpdateProblemData(pk string, result string) bool {
	if res := global.DBEngine.Exec("UPDATE problem_problemdata SET submission = submission + 1, "+
		result+" = "+result+" + 1"+" WHERE problem = ?", pk); res.Error != nil {
		logrus.Error("update problem data submission error ", res.Error)
		return false
	}
	return true
}

func UpdateProblemAuth(pk string, auth int) bool {
	if res := global.DBEngine.Exec("UPDATE  problem_problem SET auth = ? WHERE problem = ?", auth, pk); res.Error != nil {
		logrus.Error("update problem data auth error ", res.Error)
		return false
	}
	return true
}

func UpdateProblemDataAuth(pk string, auth int) bool {
	logrus.Info("pk = ", pk, " auth = ", auth)
	if res := global.DBEngine.Exec("UPDATE  problem_problemdata SET auth = ? WHERE problem = ?", auth, pk); res.Error != nil {
		logrus.Error("update problem data auth error ", res.Error)
		return false
	}
	return true
}

//---------------------------------CaseStatus-----------------------------------------//

func AddCaseStatus(status *entity.CaseStatus) bool {
	if create := global.DBEngine.Exec("INSERT INTO judgestatus_casestatus "+
		"(statusid, username, problem, result, time, memory, testcase, casedata, outputdata, useroutput)"+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		status.StatusId, status.Username, status.Problem, status.Result,
		status.Time, status.Memory, status.TestCase, status.CaseData, status.OutputData, status.UserOutput); create.Error != nil {
		logrus.Error("insert case_status error ", create.Error)
		return false
	}
	return true
}

//---------------------------------Contest-----------------------------------------//

func SetBoard(id int64, statue int) bool {
	updateBoardType := fmt.Sprintf("UPDATE contest_contestboard SET type = %d WHERE submitid = %d", statue, id)
	if res := global.DBEngine.Exec(updateBoardType); res.Error != nil {
		logrus.Error("update board type error ", res.Error)
		return false
	}
	return true
}

func GetNotExpiredContest() []*entity.ConstInfo {
	var data []*entity.ConstInfo
	res := global.DBEngine.Raw("SELECT * from contest_contestinfo where type <> 'Personal' and TO_SECONDS(NOW()) - TO_SECONDS(begintime) <= lasttime").Scan(&data)
	if res.Error != nil {
		logrus.Error("select not expired contest error ", res.Error)
		return nil
	}
	return data
}

func GetContestProblem(contestId int) []*entity.ContestProblem {
	var problems []*entity.ContestProblem
	if res := global.DBEngine.Raw("SELECT * FROM contest_contestproblem WHERE contestid = ?", contestId).Scan(&problems); res.Error != nil {
		logrus.Error("select contest problem error ", res.Error)
		return nil
	}
	return problems
}

func GetRunningContest() []*entity.ConstInfo {
	var data []*entity.ConstInfo
	res := global.DBEngine.Raw(
		"SELECT * from contest_contestinfo where type <> 'Personal' and " +
			"TO_SECONDS(NOW()) - TO_SECONDS(begintime) <= lasttime and TO_SECONDS(NOW()) - TO_SECONDS(begintime) >=-1").Scan(&data)
	if res.Error != nil {
		logrus.Error("select not expired contest error ", res.Error)
		return nil
	}
	return data
}

func UpdateContestBoardTypeBySubmitId(typ, id int) bool {
	// 后端插入 contestBoard 可能会晚于此处更新, 所以要先等待相应 contestBoard 插入后再更新
	for true {
		tp := -2
		if res := global.DBEngine.Raw("SELECT type FROM contest_contestboard  WHERE submitid = ?", id).Scan(&tp); res.Error == nil && tp != -2 {
			break
		}
	}
	if res := global.DBEngine.Exec("UPDATE contest_contestboard SET type = ?  WHERE submitid = ?", typ, id); res.Error != nil {
		logrus.Error("update contest board type error ", res.Error)
		return false
	}
	return true
}

//---------------------------------User-----------------------------------------//

func UpdateUserResult(username, result string) bool {
	updateSQL := fmt.Sprintf("UPDATE user_userdata SET %s = %s + 1 WHERE username = '%s'", result, result, username)
	if res := global.DBEngine.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user result error ", res.Error)
		return false
	}
	return true
}

func UpdateUserScore(username string, score int) bool {
	updateSQL := fmt.Sprintf("UPDATE user_userdata SET score = score+%d WHERE username = '%s'", score, username)
	if res := global.DBEngine.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user score error ", res.Error)
		return false
	}
	return true
}

func UpdateUserAcPro(problem, username string) bool {
	updateSQL := fmt.Sprintf("UPDATE user_userdata SET acpro = concat(acpro,'|%s') WHERE username = '%s'", problem, username)
	if res := global.DBEngine.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user ac problem error ", res.Error)
		return false
	}
	return true
}
