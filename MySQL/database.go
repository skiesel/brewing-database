package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
)

var (
	db *sql.DB
)

type dbConfig struct {
	Username string
	Password string
	Server string
	Database string
}

func init() {
	file, err := ioutil.ReadFile("database.config")
	if err != nil {
		panic(err)
	}

	var config dbConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@%s/%s", config.Username, config.Password, config.Server, config.Database)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	fmt.Println("connected!")	
}

func Connected() bool {
	return db != nil
}