// Package com_log provide function about log standard library
package com_log

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// Init standard library log init
func Init() {
	// default: LstdFlags = Ldate | Ltime
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// InitLogFile 设置日志 1.输出终端 2.写入日志文件
func InitLogFile(pre string) {
	// 写入日志文件
	logPath := fmt.Sprintf("%s-%s.log", pre, time.Now().Format("2006.01.02"))
	logFile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	// 终端和日志文件都输出
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)
}
