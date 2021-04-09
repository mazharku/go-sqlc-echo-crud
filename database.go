package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	toml "github.com/pelletier/go-toml"
)

func dbconn(dbconfig *toml.Tree) *sql.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v search_path=%v sslmode=disable", dbconfig.Get("dbhost").(string), dbconfig.Get("dbuser").(string), dbconfig.Get("dbpass").(string), dbconfig.Get("dbname").(string), dbconfig.Get("schema").(string))
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected...")
	}

	db.SetMaxOpenConns(SetMaxOpenConns)
	return db
}
