package sql

import (
	"database/sql"
	"fmt"
	"github.com/akimdev15/algolock/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

type DbStruct struct {
	DB      *sql.DB
	Queries *database.Queries
}

func InitDB() (DbStruct, error) {
	db, err := sql.Open("sqlite3", "sql/mydb.db")
	if err != nil {
		fmt.Println("Error opening database", err)
		return DbStruct{}, err
	}

	queries := database.New(db)

	return DbStruct{DB: db, Queries: queries}, nil
}
