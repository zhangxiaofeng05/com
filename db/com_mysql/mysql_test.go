package com_mysql_test

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/zhangxiaofeng05/com/db/com_mysql"
)

func TestGetEnv(t *testing.T) {
	halfDsn := com_mysql.GetEnv()
	dsn := fmt.Sprintf("%s/%s?parseTime=true", halfDsn, "dbname")
	t.Logf("mysql dsn: %v", dsn)
}

func ExampleGetEnv() {
	// $ go get github.com/go-sql-driver/mysql

	// import _ "github.com/go-sql-driver/mysql"
	halfDsn := com_mysql.GetEnv()
	dsn := fmt.Sprintf("%s/%s?parseTime=true", halfDsn, "dev")
	fmt.Println(dsn)
	db, err := sql.Open(com_mysql.DriverName, dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("mysql db.Close() err: %v", err)
		}
	}()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
