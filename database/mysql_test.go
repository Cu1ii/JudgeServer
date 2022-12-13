package database

import (
	"fmt"
	"testing"
	"xoj_judgehost/dao"
)

func TestSelectPendingJudgeStatus(t *testing.T) {
	all := []*dao.JudgeStatus{}
	mySQLDB := GetMySQLDB()
	if find := mySQLDB.Raw("SELECT * FROM judgestatus_judgestatus").Scan(&all); find.Error != nil {
		fmt.Println("find all judge_status fail")
	}
	for _, stauts := range all {
		fmt.Println(stauts)
	}
}

func TestSelectAllProblem(t *testing.T) {
	all := []*dao.Problem{}
	mySQLDB := GetMySQLDB()
	if find := mySQLDB.Raw("SELECT * FROM problem_problem").Scan(&all); find.Error != nil {
		fmt.Println("find all problem fail")
	}
	for _, stauts := range all {
		fmt.Println(stauts)
	}
}

func TestSelectAllProblemData(t *testing.T) {
	all := []*dao.ProblemData{}
	mySQLDB := GetMySQLDB()
	if find := mySQLDB.Raw("SELECT * FROM problem_problemdata").Scan(&all); find.Error != nil {
		fmt.Println("find all problem fail")
	}
	for _, stauts := range all {
		fmt.Println(stauts)
	}
}
