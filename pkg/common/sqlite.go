package common

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 打开或者创建数据库、表
func init() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	sql_table := `
    CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
    );
    `
	db.Exec(sql_table)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)
	stmt.Exec("wangshubo", "国务院", "2017-04-21")
	//checkErr(err)

	db.Close()
}
