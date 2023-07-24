package com_mysql_test

import (
	"database/sql"
	"fmt"
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
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
