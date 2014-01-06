package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"pengyou/app/models/db"
	"time"
)

type Active struct {
	Id          int
	Description string
	Starttime   time.Time
	Endtime     time.Time
	Owner       int
	Total       float32
	Type        bool
	Title       string
}

func ListActive() {
	db, err := sql.Open("mysql", "root@/py")
	checkerr(err)
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT * FROM user")
	checkerr(err)
	defer stmtOut.Close()

}

func (active *Active) Save() error {
	query := "INSERT INTO active(description, starttime, endtime, owner, total, type, title) values(?,?,?,?,?,?,?)"
	db.Save(query, active.Description, active.Starttime, active.Endtime, active.Owner, active.Total, active.Type, active.Title)
	return nil
}

func checkerr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
