package sql

import (
	"database/sql"
	"fmt"
	"github.com/akimdev15/algolock/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*database.Queries, error) {
	db, err := sql.Open("sqlite3", "sql/mydb.db")
	if err != nil {
		fmt.Println("Error opening database", err)
		return nil, err
	}
	//defer db.Close()

	queries := database.New(db)

	return queries, nil
}
