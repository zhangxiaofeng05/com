package com_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/zhangxiaofeng05/com/com_env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jmoiron/sqlx"
)

const (
	DriverName = "mysql"

	DSN = "MYSQL_DSN"
)

func GetEnv() (halfDsn string) {
	// DSN (Data Source Name) : https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dbUser := com_env.GetEnv("MYSQL_USERNAME", "root")
	dbPass := com_env.GetEnv("MYSQL_PASSWORD", "123456")
	dbHost := com_env.GetEnv("MYSQL_HOST", "127.0.0.1")
	dbPort := com_env.GetEnv("MYSQL_HOST_PORT", "3306")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)", dbUser, dbPass, dbHost, dbPort)
}

// Sql Standard library
func Sql(dsn string) (sqlDB *sql.DB, err error) {
	db, err := sql.Open(DriverName, dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func Sqlx(dsn string) (sqlxDB *sqlx.DB, err error) {
	db, err := sqlx.Open(DriverName, dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

func Gorm(dsn string) (gormDB *gorm.DB, err error) {
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return gormDB, err
}
