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
	var judgeArry []*entity.JudgeStatus
	if res := global.DBEngine.Raw("SELECT * FROM judge_status where result = ?", global.PENDING).Scan(&judgeArry); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return judgeArry
}

func GetJudgeStatusById(id int64) *entity.JudgeStatus {
	status := entity.JudgeStatus{}
	if res := global.DBEngine.Raw("SELECT * FROM judge_status where id = ?", id).Scan(&status); res.Error != nil {
		logrus.Error("select judge status error ", res.Error)
		return nil
	}
	return &status
}

func UpdateJudgeStatusResult(id int, result int) bool {
	if res := global.DBEngine.Exec("UPDATE judge_status SET result = ? WHERE id = ?", result, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatusMessage(id int, msg string) bool {
	if res := global.DBEngine.Exec("UPDATE judge_status SET message = ? WHERE id = ?", msg, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatusJudger(id int, judger string) bool {
	if res := global.DBEngine.Exec("UPDATE judge_status SET judger = ? WHERE id = ?", judger, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

func UpdateJudgeStatus(id int, memory, mytime int, result int, testcase string) bool {
	if res := global.DBEngine.Exec("UPDATE judge_status SET memory = ?, time= ?, result = ?, testcase=?  WHERE id = ?", memory, mytime, result, testcase, id); res.Error != nil {
		logrus.Error("update judge status error ", res.Error)
		return false
	}
	return true
}

//---------------------------------Problem-----------------------------------------//

func GetProblemById(pk string) *entity.Problem {
	problem := entity.Problem{}
	if res := global.DBEngine.Raw("SELECT * FROM problem WHERE problem = ?", pk).Scan(&problem); res.Error != nil {
		logrus.Error("select problem error ", res.Error)
		return nil
	}
	return &problem
}

func GetIsHaveDoneProblem(username, problem string) bool {
	selectProblem := fmt.Sprintf("SELECT * FROM judge_status WHERE user = '%s'  AND problem = '%s' AND result = 0", username, problem)
	var problems []entity.Problem
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
	addProSubmitNum := fmt.Sprintf("UPDATE problem_data SET submission = submission+1 WHERE problem = '%s'", problem)
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
	if res := global.DBEngine.Raw("SELECT * FROM problem_data WHERE problem = ?", pk).Scan(&problemData); res.Error != nil {
		logrus.Error("select problem data error ", res.Error)
		return nil
	}
	return &problemData
}

func UpdateProblemData(pk string, result string) bool {
	if res := global.DBEngine.Exec("UPDATE problem_data SET submission = submission + 1, "+
		result+" = "+result+" + 1"+" WHERE problem = ?", pk); res.Error != nil {
		logrus.Error("update problem data submission error ", res.Error)
		return false
	}
	return true
}

func UpdateProblemAuth(pk string, auth int) bool {
	if res := global.DBEngine.Exec("UPDATE  problem SET auth = ? WHERE problem = ?", auth, pk); res.Error != nil {
		logrus.Error("update problem data auth error ", res.Error)
		return false
	}
	return true
}

func UpdateProblemDataAuth(pk string, auth int) bool {
	logrus.Info("pk = ", pk, " auth = ", auth)
	if res := global.DBEngine.Exec("UPDATE  problem_data SET auth = ? WHERE problem = ?", auth, pk); res.Error != nil {
		logrus.Error("update problem data auth error ", res.Error)
		return false
	}
	return true
}

//---------------------------------CaseStatus-----------------------------------------//

func AddCaseStatus(status *entity.CaseStatus) bool {
	if create := global.DBEngine.Exec("INSERT INTO case_status "+
		"(status_id, username, problem, result, time, memory, testcase, case_data, output_data, user_output)"+
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
	updateBoardType := fmt.Sprintf("UPDATE contest_board SET type = %d WHERE submit_id = %d", statue, id)
	if res := global.DBEngine.Exec(updateBoardType); res.Error != nil {
		logrus.Error("update board type error ", res.Error)
		return false
	}
	return true
}

// GetNotExpiredContest 废弃
func GetNotExpiredContest() []*entity.ConstInfo {
	var data []*entity.ConstInfo
	res := global.DBEngine.Raw("SELECT * from contest_info where type <> 'Personal' and TO_SECONDS(NOW()) - TO_SECONDS(begin_time) <= last_time").Scan(&data)
	if res.Error != nil {
		logrus.Error("select not expired contest error ", res.Error)
		return nil
	}
	return data
}

func GetContestProblem(contestId int) []*entity.ContestProblem {
	var problems []*entity.ContestProblem
	if res := global.DBEngine.Raw("SELECT * FROM contest_problem WHERE contest_id = ?", contestId).Scan(&problems); res.Error != nil {
		logrus.Error("select contest problem error ", res.Error)
		return nil
	}
	return problems
}

// GetRunningContest 暂时废弃
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
		if res := global.DBEngine.Raw("SELECT type FROM contest_board  WHERE submit_id = ?", id).Scan(&tp); res.Error == nil && tp != -2 {
			break
		}
	}
	if res := global.DBEngine.Exec("UPDATE contest_board SET type = ?  WHERE submit_id = ?", typ, id); res.Error != nil {
		logrus.Error("update contest board type error ", res.Error)
		return false
	}
	return true
}

//---------------------------------User-----------------------------------------//

func UpdateUserResult(username, result string) bool {
	updateSQL := fmt.Sprintf("UPDATE user_data SET %s = %s + 1 WHERE username = '%s'", result, result, username)
	if res := global.DBEngine.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user result error ", res.Error)
		return false
	}
	return true
}

func UpdateUserScore(username string, score int) bool {
	updateSQL := fmt.Sprintf("UPDATE user_data SET score = score+%d WHERE username = '%s'", score, username)
	if res := global.DBEngine.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user score error ", res.Error)
		return false
	}
	return true
}

func UpdateUserAcPro(problem, username string) bool {
	updateSQL := fmt.Sprintf("INSERT `ac_problem` (`username`, `problem_id`) values ('%s', '%s')", username, problem)
	if res := global.DBEngine.Exec(updateSQL); res.Error != nil {
		logrus.Error("update user ac problem error ", res.Error)
		return false
	}
	return true
}
