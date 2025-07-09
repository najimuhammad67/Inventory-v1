
package config

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./inventory.db")
	if err != nil {
		log.Fatal(err)
	}
	createTables()
	seedAdmin()
}

func createTables() {
	createUser := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		role TEXT NOT NULL DEFAULT 'user');`

	createInventory := `CREATE TABLE IF NOT EXISTS inventory (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		category TEXT,
		quantity INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP);`

	_, err := DB.Exec(createUser)
	if err != nil {
		log.Fatal("Create user table:", err)
	}
	_, err = DB.Exec(createInventory)
	if err != nil {
		log.Fatal("Create inventory table:", err)
	}
}

func seedAdmin() {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'admin'").Scan(&count)
	if err == nil && count == 0 {
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		_, err = DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", "admin", string(hash), "admin")
		if err != nil {
			log.Println("Gagal membuat akun admin:", err)
		} else {
			log.Println("✅ Akun admin berhasil dibuat: admin / admin123")
		}
	}
}
