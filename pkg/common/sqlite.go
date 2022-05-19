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

var G_Sql_DB_Obj *sql.DB

func InitSqlite(sqlFile string) {
	if G_Sql_DB_Obj != nil {
		return
	}

	var err error
	G_Sql_DB_Obj, err = sql.Open("sqlite3", sqlFile)
	if err != nil {
		panic(err)
	}

	createSqlTable := `
    CREATE TABLE IF NOT EXISTS PortInfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        host VARCHAR(64) NULL,
		proto varchar(32) null,
        port VARCHAR(64) NULL,
		portInfo varchar(64) null
    );
    `
	G_Sql_DB_Obj.Exec(createSqlTable)
}

func AppendAsset2Sql(host, proto, port, portInfo string) {
	stmt, err := G_Sql_DB_Obj.Prepare("INSERT INTO PortInfo(host, proto, port, portInfo) values(?,?,?,?)")
	if err != nil {
		panic(err)
	}
	stmt.Exec(host, proto, port, portInfo)
}

func CloseDB() {
	G_Sql_DB_Obj.Close()
}
