package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func GetInstance() *sql.DB {
	// TODO 環境変数に変える
	var dataSourceName = fmt.Sprintf("")
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		fmt.Println(err)
	}
	log.Print("db", db)
	return db
}
