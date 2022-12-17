package util

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestUnZipInDictionary(t *testing.T) {
	//ZipDictionary("/home/cu1/XOJ/resolutions/1/1.zip", "/home/cu1/XOJ/resolutions/1")
	zipPath := fmt.Sprintf("%s/%s/%s.zip", "/home/cu1/XOJ/resolutions", "1", "1")
	fmt.Println(zipPath)
	if IsFileIn(zipPath) {
		logrus.Info("run this")
		if isSuccess, err := UnZipInDictionary(zipPath, "/home/cu1/XOJ/resolutions/1"); !isSuccess || err != nil {
			logrus.Error("unzip error ", err.Error())
		}
		if isSuccess, err := DeleteFile(zipPath, true); !isSuccess || err != nil {
			logrus.Error("delete error")
		}
	}
}
