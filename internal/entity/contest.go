package entity

import "time"

type ConstInfo struct {
	Id      int `gorm:"primaryKey"`
	Creator string
	Oj      string
	Title   string
	//Level     int
	Des string
	//Note      string
	BeginTime time.Time
	LastTime  int
	Type      string
	Auth      int
	CloneFrom int
	Password  string
	//Classes   string
	IpRange   string
	LockBoard int
	LockTime  int
}

type ContestProblem struct {
	Id           int    `gorm:"primaryKey"`
	ContestId    int    `gorm:"column:contestid"`
	ProblemId    string `gorm:"column:problemid"`
	ProblemTitle string `gorm:"column:problemtitle"`
	Rank         int
}
