package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
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
	migrationScript, err := ioutil.ReadFile("C:/Users/wotbl/go-figma/migrations/SignUp.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	_, err = DB.Exec(string(migrationScript))
	if err != nil {
		log.Fatal("Error running migration:", err)
	}

	migrationScript, err = ioutil.ReadFile("C:/Users/wotbl/go-figma/migrations/CreateDraft.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	_, err = DB.Exec(string(migrationScript))
	if err != nil {
		log.Fatal("Error running migration:", err)
	}

	migrationScript, err = ioutil.ReadFile("C:/Users/wotbl/go-figma/migrations/CreateFigure.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	_, err = DB.Exec(string(migrationScript))
	if err != nil {
		log.Fatal("Error running migration:", err)
	}




	migrationScript, err = ioutil.ReadFile("C:/Users/wotbl/go-figma/migrations/CreateLine.sql")
	if err != nil {
		log.Fatal("Error reading migration file:", err)
	}

	_, err = DB.Exec(string(migrationScript))
	if err != nil {
		log.Fatal("Error running migration:", err)
	}



	fmt.Println("Migration executed successfully")
}


 