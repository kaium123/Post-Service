package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" 
	"github.com/spf13/viper"
)

var db *sql.DB

func InitDB() *sql.DB {
	dbUrl := viper.GetString("DB_URL")
	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
		CREATE TABLE  IF NOT EXISTS  posts (
			id SERIAL PRIMARY KEY,
			content TEXT,
			user_id INT
		)
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
