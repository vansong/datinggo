package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goinspect.cn/app/conf"
)

var (
	mysqlDb *sql.DB
)

func NewMysql(conf *conf.Database) (db *sql.DB) {
	if mysqlDb != nil {
		db = mysqlDb
		return
	}
	var dsn = fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?charset=utf8&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(100)
	mysqlDb = db
	return
}