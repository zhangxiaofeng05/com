package commysql_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/zhangxiaofeng05/com/db/commysql"
)

func TestGetEnv(t *testing.T) {
	halfDsn := commysql.GetEnv()
	dsn := fmt.Sprintf("%s/%s?parseTime=true", halfDsn, "dbname")
	t.Logf("mysql dsn: %v", dsn)
}

func ExampleGetEnv() {
	// $ go get github.com/go-sql-driver/mysql

	// import _ "github.com/go-sql-driver/mysql"
	halfDsn := commysql.GetEnv()
	dsn := fmt.Sprintf("%s/%s?parseTime=true", halfDsn, "dev")
	fmt.Println(dsn)
	db, err := sql.Open(commysql.DriverName, dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
