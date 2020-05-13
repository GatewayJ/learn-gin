package connect

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "dev_admin:6&LpBP#^S07VSyv@tcp(106.14.127.93:3306)/qeeyou_client?parseTime=true")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("init db")
	Db.SetMaxIdleConns(20)
	Db.SetMaxOpenConns(20)
}
