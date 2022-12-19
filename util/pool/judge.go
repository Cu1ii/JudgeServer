package pool

import (
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

var CORE_POOL_SIZE int = runtime.NumCPU()

const (
	EXPIRY_DURATION    = time.Second * 5
	PRE_ALLOC          = false
	MAX_BLOCKING_TASKS = 20
	NONBLOCKING        = true
)

var JudgePool = (*ants.Pool)(nil)

func NewJudgePool() *ants.Pool {
	judgePool, err := ants.NewPool(CORE_POOL_SIZE, ants.WithOptions(ants.Options{
		ExpiryDuration:   EXPIRY_DURATION,
		PreAlloc:         PRE_ALLOC,
		MaxBlockingTasks: MAX_BLOCKING_TASKS,
		Nonblocking:      NONBLOCKING,
	}))
	if err != nil {
		logrus.Error("new judge pool fail: ", err.Error())
	}
	return judgePool
}

func GetJudgePool() *ants.Pool {
	if JudgePool == nil {
		JudgePool = NewJudgePool()
	}
	return JudgePool
}
