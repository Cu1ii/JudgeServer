package logs

import (
	"bytes"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"time"
)

type mineFomatter struct{}

func (m *mineFomatter) Format(entry *logrus.Entry) ([]byte, error) {
	return nil, nil
}

func init() {
	initRuntimeLog()
}

func initRuntimeLog() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&nested.Formatter{
		TimestampFormat: time.RFC3339,
	})
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("/home/cu1/Project/Go/xoj_judgehost/log/runtime.logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("create file log/runtime.logs failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("runtime log start")
}
