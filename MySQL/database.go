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
	Server   string
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Server, config.Database)

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	fmt.Println("connected!")

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
}

func GetBeerXMLRows() *sql.Rows {
	query := "SELECT ElementType.Type, BeerXMLMetaData.Field, BeerXMLMetaData.Description, BeerXMLMetaData.Format, BeerXMLMetaData.Required FROM BeerXMLMetaData INNER JOIN ElementType ON BeerXMLMetaData.ElementType=ElementType.ID ORDER BY Type"
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	return rows
}
