package comlogrus_test

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/zhangxiaofeng05/com/comlog/comlogrus"
)

func TestInit(t *testing.T) {
	logrus.Info("logrus info")
	comlogrus.Init()
	logrus.Info("comlogrus info")

	// add Fields
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
