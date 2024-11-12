package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

//init db connection
func InitDB() (*sql.DB, error) {
	godotenv.Load()

	//define db path from env variable
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is not set")
	}

	//connect to sqlite db or create
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	//create user table
	createUsersTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		email TEXT UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createUsersTableQuery)
	if err != nil {
		return nil, err
	}

	//create pref table
	createPreferencesTableQuery := `
	CREATE TABLE IF NOT EXISTS preferences (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		sort_order TEXT,
		hidden_devices TEXT,
		custom_icons TEXT,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);`
	_, err = db.Exec(createPreferencesTableQuery)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to SQLite database")
	return db, nil
}

//close connection
func CloseDB(db *sql.DB) {
	if db != nil {
		db.Close()
		log.Println("Database connection closed")
	}
}
