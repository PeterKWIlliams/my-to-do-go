package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	projectRoot, err := getProjectRoot()
	if err != nil {
		log.Fatalf("Error getting project root: %v", err)
	}

	envPath := filepath.Join(projectRoot, ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	DBENV := os.Getenv("DBNAME")
	DBFile := filepath.Join(projectRoot, DBENV+".db")
	if _, err := os.Stat(DBFile); os.IsNotExist(err) {
		log.Printf("Database file does not exist. It will be created: %v", DBFile)
	}

	db, err := sql.Open("sqlite3", DBFile)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Database file is located at:", DBFile)
	// data comes in the form of
	// mytodo project create myproject --directory=~/myproject --time=10

	// dbQueries := database.New(db)
	// cfg := &config.Config{
	// 	DB: dbQueries,
	// }
}

func getProjectRoot() (string, error) {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get caller information")
	}
	basepath := filepath.Dir(b)

	for {
		if _, err := os.Stat(filepath.Join(basepath, "go.mod")); err == nil {
			return basepath, nil
		}
		newBasepath := filepath.Dir(basepath)
		if newBasepath == basepath {
			break
		}
		basepath = newBasepath
	}

	return "", fmt.Errorf("go.mod not found in any parent directory")
}
