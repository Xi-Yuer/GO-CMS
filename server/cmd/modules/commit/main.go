package main

import (
	"database/sql"
	"fmt"
	"github.com/Xi-Yuer/cms/config"
	"log"
	"os/exec"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Commit struct to represent a Git commit
type Commit struct {
	ID      string
	Author  string
	Email   string
	Date    time.Time
	Message string
}

func main() {
	// Database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/cms?charset=utf8&parseTime=True", config.Config.DB.NAME, config.Config.DB.PASSWORD, config.Config.DB.HOST, config.Config.DB.PORT))
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// Ensure the database connection is available
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Execute git log command
	cmd := exec.Command("git", "log", "--pretty=format:%H|%an|%ae|%ad|%s")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Split the output by newline
	lines := strings.Split(string(output), "\n")
	// emptyDataBase
	_, err = db.Exec("DELETE FROM cms.commits")
	// Parse each line and insert into the database
	for _, line := range lines {
		fields := strings.Split(line, "|")
		if len(fields) == 5 {
			commit := Commit{
				ID:      fields[0],
				Author:  fields[1],
				Email:   fields[2],
				Message: fields[4],
			}
			dateStr := strings.TrimSpace(fields[3])
			date, err := time.Parse("Mon Jan 2 15:04:05 2006 -0700", dateStr)
			if err != nil {
				log.Printf("Error parsing date: %s", err)
			} else {
				commit.Date = date
			}

			// Insert commit into the database
			_, err = db.Exec("INSERT INTO cms.commits (commit_id, author, email, commit_date, message) VALUES (?, ?, ?, ?, ?)",
				commit.ID, commit.Author, commit.Email, commit.Date, commit.Message)
			if err != nil {
				log.Printf("Error inserting commit: %s", err)
			}
		}
	}

	fmt.Println("Git commits inserted into the database successfully!")
}
