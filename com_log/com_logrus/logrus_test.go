package com_logrus_test

import (
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/zhangxiaofeng05/com/com_log/com_logrus"
)

func TestInit(t *testing.T) {
	logrus.Info("logrus info")
	com_logrus.Init()
	logrus.Info("com_logrus info")

	// add Fields
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
