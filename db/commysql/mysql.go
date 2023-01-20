package commysql

import (
	"fmt"

	"github.com/zhangxiaofeng05/com/comutil"
)

const (
	DriverName = "mysql"

	DSN = "MYSQL_DSN"
)

func GetEnv() (halfDsn string) {
	// DSN (Data Source Name) : https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dbUser := comutil.GetEnv("MYSQL_USER", "root")
	dbPass := comutil.GetEnv("MYSQL_PASS", "test")
	dbHost := comutil.GetEnv("MYSQL_HOST", "127.0.0.1")
	dbPort := comutil.GetEnv("MYSQL_PORT", "3306")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)", dbUser, dbPass, dbHost, dbPort)
}
