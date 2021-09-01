package main

import (
	"99-project/dbconfig"
	"99-project/entity"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	fmt.Printf("Accessing %s ... ", dbconfig.DbName)

	db, err = sql.Open(dbconfig.PostgresDriver, fmt.Sprintf("postgres://root@%s:%s/project?sslmode=disable",dbconfig.Host, dbconfig.Port ))

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected!")
	}

	defer db.Close()

	sqlSelect()
}


func sqlSelect() {

	sqlStatement, err := db.Query("SELECT id, name, documentnumber FROM users")
	checkErr(err)

	for sqlStatement.Next() {

		var entity entity.User

		err = sqlStatement.Scan(&entity.ID, &entity.Name, &entity.DocumentNumber)
		checkErr(err)

		fmt.Printf("%d\t%s\t%s \n", entity.ID, entity.Name, entity.DocumentNumber)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}