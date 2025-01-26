package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	
	
	"fmt"
	"io/ioutil"
	 


)
var DB *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://test:test@localhost/figma?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	RunMigration()
}



func RunMigration() {
	// Чтение SQL-скрипта из файла
	//migrationScript, err := ioutil.ReadFile("../migrations/SignUp.sql")
	migrationScript, err := ioutil.ReadFile("C:/Users/wotbl/go-figma/migrations/SignUp.sql")

	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	// Выполнение SQL-скрипта
	_, err = DB.Exec(string(migrationScript))
	if err != nil {
		log.Fatal("Error running migration:", err)
	}

	fmt.Println("Migration executed successfully")
}
