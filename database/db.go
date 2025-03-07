package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDB() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "./users.db"
	}


	var err error

	DB, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal("Error with connect db: ", err)

	}

	createTable := `Create table if not exists users(
		id integer primary key autoincrement,
		username text unique not null,
		password text not null		
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal("Error with creating users table: ", err)
	}
	createSessionsTable := `Create table if not exists sessions (
		id integer primary key autoincrement,
		username text not null,
		session_token text unique not null,
		created_at datetime default current_timestamp
	);`
	_, err = DB.Exec(createSessionsTable)
	if err != nil {
		log.Fatal("Error with creating sessions table: ", err)
	}
	log.Println("Database successfully created ")
}
