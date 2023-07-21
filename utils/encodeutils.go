package models

import (
	"database/sql"
	"log"
)

func InitDB() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", "postgres:atada0x:ethereumsolana@localhost/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		stmt, err := db.Prepare("CREATE TABLE IF NOT EXIST WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL);")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return db, nil
	}
}
