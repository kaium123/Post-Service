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

	createTableSQL = `
		CREATE TABLE  IF NOT EXISTS  reacts (
			id SERIAL PRIMARY KEY,
			post_id INT,
			reacted_user_id INT,
			post_type TEXT
		)
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL = `
		CREATE TABLE  IF NOT EXISTS  comments (
			id SERIAL PRIMARY KEY,
			post_id INT,
			commented_user_id INT,
			content TEXT

		)
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL = `
		CREATE TABLE  IF NOT EXISTS  shares (
			id SERIAL PRIMARY KEY,
			post_id INT,
			shared_user_id INT
		)
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
