package db

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/mattn/go-sqlite3"
	"log"
)

func Db() {
	dbmap := initDb()
	defer dbmap.Db.close()

}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func initDb() *gorp.Dbmap {

}
