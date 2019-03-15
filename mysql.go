package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Result struct {
	count int
}

func main() {
	name := "root"
	password := "root"
	endpoint := "localhost"
	dbname := "sample"
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", name, password, endpoint, dbname)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT COUNT(*) FROM users") //
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		result := Result{}
		err := rows.Scan(&result.count)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("count: %v", result.count)
	}
}
