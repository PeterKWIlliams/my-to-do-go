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

	"github.com/PeterKWIlliams/my-to-do-go/internal/commands"
	"github.com/PeterKWIlliams/my-to-do-go/internal/database"
	"github.com/PeterKWIlliams/my-to-do-go/internal/mytodo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	projectRoot, err := getProjectRoot()
	if err != nil {
		log.Fatalf("Error getting project root: %v", err)
	}

	DBENV := os.Getenv("DBNAME") + ".db"
	DBFile := filepath.Join(projectRoot, DBENV)

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
	dbQueries := database.New(db)

	cfg := &mytodo.Config{
		DB: dbQueries,
	}
	commands := commands.GetCommands()
	arguments := os.Args[1:]
	command, ok := commands[arguments[1]]
	if !ok {
		fmt.Println("command not found type help for a list of commands")
		return
	}
	err = command.Callback(cfg, arguments[2:]...)
	if err != nil {
		fmt.Println(err)
	}
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
