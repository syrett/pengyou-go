package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Save(query string, args ...interface{}) error {
	db, err := sql.Open("mysql", "root@/py")
	checkerr(err)
	defer db.Close()

	stmtIns, err := db.Prepare(query)
	checkerr(err)
	defer stmtIns.Close()

	_, err = stmtIns.Exec(args...)

	checkerr(err)
	return nil
}

func List() {
	db, err := sql.Open("mysql", "root@/py")
	checkerr(err)
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO user(name, password, email) VALUES(?, ?, ?)")
	checkerr(err)
	defer stmtIns.Close()

	stmtOut, err := db.Prepare("SELECT password FROM user WHERE name=?")
	checkerr(err)
	defer stmtOut.Close()

	/*	_, err = stmtIns.Exec("liyouyou", "123456", "liyouyou@shopex.cn")
		checkerr(err)
	*/
	var password string

	err = stmtOut.QueryRow("liyouyou").Scan(&password)
	checkerr(err)
	fmt.Printf("liyouyou password:%s\n", password)

	rows, err := db.Query("SELECT * FROM user")
	checkerr(err)
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			checkerr(err)
		}
		fmt.Printf("%s\n", name)
	}
}

func checkerr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
