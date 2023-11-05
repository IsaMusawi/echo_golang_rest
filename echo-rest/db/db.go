package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"restgolang.com/restGolang/echo-rest/config"
)

var db *sql.DB
var err error

func Init() {
	conf :=  config.GetConfig()
	fmt.Print(conf)

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT +  ")/" + conf.DB_NAME
	
	
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("connectionString error ..")
	}

	fmt.Print(connectionString)

	// err = db.Ping()
	// if err != nil {
	// 	panic("DSN invalid")
	// }
}

func CreateCon() *sql.DB {
	return db
}